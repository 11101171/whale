//****************************************************/
// Addres: https://github.com/11101171
// Author: ningzhong.zeng
//****************************************************/

package models

import (
	"strings"

	"github.com/go-gorp/gorp"
	"github.com/revel/revel"
)

const (
	EMAIL_NO_SEND = iota
	EMAIL_SEND_IF_SUC
	EMAIL_SEND_IF_ERR
	EMAIL_SEND_IF_END
)

type Task struct {
	Base
	TaskId        string
	UserId        string
	GroupId       string
	TaskName      string
	TaskType      int
	Description   string
	CronSpec      string
	Concurrent    int
	Command       string
	Status        int
	Notify        int
	NotifyEmail   string
	NotifyContent string
	Timeout       int
	ExecuteTimes  int
}

func (t *Task) Validate(v *revel.Validation) {
	t.TaskName = strings.TrimSpace(t.TaskName)
	// t.TaskType = strings.TrimSpace(t.TaskTypepe)
	t.Description = strings.TrimSpace(t.Description)
	t.CronSpec = strings.TrimSpace(t.CronSpec)
	// t.Concurrent = strings.TrimSpace(t.Concurrent)
	t.Command = strings.TrimSpace(t.Command)
	// t.Status = strings.TrimSpace(t.Status)
	// t.Notify = strings.TrimSpace(t.Notify)
	// t.NotifyEmail = strings.TrimSpace(t.NotifyEmail)
	// t.Timeout = strings.TrimSpace(t.Timeout)
	// t.ExecuteTimes = strings.TrimSpace(t.ExecuteTimes)

	v.Check(t.TaskName,
		revel.Required{},
	).Key("task.TaskName").Message("任务名称不能为空")
	v.Check(t.Command,
		revel.Required{},
	).Key("task.Command").Message("脚本不能为空")
}

func (t *Task) ValidateInsert(v *revel.Validation) {
	t.Validate(v)
	t.ExecuteTimes = 0
}

func (t *Task) ValidateUpdate(v *revel.Validation) {
	t.Validate(v)
}

func (t *Task) PreInsert(_ gorp.SqlExecutor) error {
	t.TaskId = RandomString("task", 6)
	return t.PreI()
}

func InsertTaskOne(task *Task) bool {
	err := dbmap.Insert(task)
	return CheckErr(err, "Insert Task One")
}

func SelectTaskList(status int) (tasks []Task) {
	CheckErrSQLSelectList(dbmap.Select(
		&tasks,
		"select * from t_task where status=:status order by GmtCreate desc",
		map[string]int{
			"status": status,
		},
	))
	return tasks
}

func SelectTaskListByUserId(userId string) (tasks []Task) {
	CheckErrSQLSelectList(dbmap.Select(
		&tasks,
		"select * from t_task where UserId=:UserId order by GmtCreate desc",
		map[string]string{
			"UserId": userId,
		},
	))
	return tasks
}

func SelectTaskListByGroupId(groupId string) (tasks []Task) {
	CheckErrSQLSelectList(dbmap.Select(
		&tasks,
		"select * from t_task where GroupId=:GroupId order by GmtCreate desc",
		map[string]string{
			"GroupId": groupId,
		},
	))
	return tasks
}

func SelectTaskOneById(taskId string) (task Task) {
	CheckErr(dbmap.SelectOne(
		&task,
		"select * from t_task where TaskId=:TaskId",
		map[string]string{
			"TaskId": taskId,
		},
	), "Select Task One")
	return task
}

func UpdateTaskOne(task *Task) bool {
	return Exec(
		`update t_task 
		 set TaskName=:TaskName,TaskType=:TaskType,Description=:Description,
			CronSpec=:CronSpec,Concurrent=:Concurrent,Command=:Command,Status=:Status,
			Notify=:Notify,NotifyEmail=:NotifyEmail,NotifyContent=:NotifyContent,Timeout=:Timeout,ExecuteTimes=:ExecuteTimes 
		 where TaskId=:TaskId`,
		map[string]interface{}{
			"TaskName":      task.TaskName,
			"TaskType":      task.TaskType,
			"Description":   task.Description,
			"CronSpec":      task.CronSpec,
			"Concurrent":    task.Concurrent,
			"Command":       task.Command,
			"Status":        task.Status,
			"Notify":        task.Notify,
			"NotifyEmail":   task.NotifyEmail,
			"NotifyContent": task.NotifyContent,
			"Timeout":       task.Timeout,
			"ExecuteTimes":  task.ExecuteTimes,
			"TaskId":        task.TaskId,
		},
	)
}

func DeleteTaskOneById(taskId string) bool {
	return Exec(
		"delete from t_task where TaskId=:TaskId",
		map[string]string{
			"TaskId": taskId,
		},
	)
}
