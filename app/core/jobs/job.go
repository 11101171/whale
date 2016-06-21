//****************************************************/
// Addres: https://github.com/11101171
// Author: ningzhong.zeng
//****************************************************/

package jobs

import (
	"bytes"
	"fmt"
	"os/exec"
	"runtime/debug"
	"sync"
	"time"
	"whale/app/models"

	"github.com/revel/revel"
)

type Job struct {
	id         string
	userId     string
	name       string
	task       *models.Task
	runFunc    func() ([]byte, []byte, error)
	running    sync.Mutex
	status     int
	Concurrent bool
}

func NewJobFromTask(task *models.Task) (*Job, error) {
	if task.TaskId == "" {
		return nil, fmt.Errorf("ToJob: 缺少id")
	}
	job := NewCommandJob(task.TaskId, task.TaskName, task.Command)
	job.task = task
	job.Concurrent = task.Concurrent == 0
	return job, nil
}

func NewCommandJob(id string, name string, command string) *Job {
	job := &Job{
		id:   id,
		name: name,
	}
	job.runFunc = func() ([]byte, []byte, error) {
		bufOut := new(bytes.Buffer)
		bufErr := new(bytes.Buffer)
		cmd := exec.Command("/bin/bash", "-c", command)
		cmd.Stdout = bufOut
		cmd.Stderr = bufErr
		err := cmd.Run()

		return bufOut.Bytes(), bufErr.Bytes(), err
	}
	return job
}

func (j *Job) Status() int {
	return j.status
}

func (j *Job) GetName() string {
	return j.name
}

func (j *Job) GetId() string {
	return j.id
}

func (j *Job) GetUserId() string {
	return j.userId
}

func (j Job) Run() {
	//revel.INFO.Panicln("执行---------")
	revel.INFO.Println("---------run----------")
	defer func() {
		if err := recover(); err != nil {
			revel.ERROR.Println(err, "\n", string(debug.Stack()))
		}
	}()

	if j.Concurrent {
		j.running.Lock()
		defer j.running.Unlock()
	}

	if workPool != nil {
		workPool <- true
		defer func() {
			<-workPool
		}()
	}

	j.status = 1
	defer func() {
		j.status = 0
	}()

	//_, _, err := j.runFunc()
	bout, berr, err := j.runFunc()
	if err != nil {
		revel.ERROR.Println("job.runFunc", err)
	}

	t := time.Now()
	ut := time.Now().Sub(t) / time.Millisecond

	// 插入日志
	log := new(models.TaskLog)
	log.UserId = j.task.UserId
	log.TaskId = j.id
	log.TaskName = j.task.TaskName
	log.Output = string(bout)
	log.Error = string(berr)
	log.ProcessTime = int(ut)
	if err != nil {
		log.Status = 0
		log.Error = err.Error() + ":" + string(berr)
	} else {
		log.Status = 1
	}
	models.InsertTaskLogOne(log)

	// 更新上次执行时间
	// j.task.PrevTime = t.Unix()
	j.task.ExecuteTimes++
	models.UpdateTaskOne(j.task)
}
