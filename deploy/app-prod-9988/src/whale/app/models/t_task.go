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

type Task struct {
	Base
	TaskId       string
	UserId       string
	GroupId      string
	TaskName     string
	TaskType     int
	Description  string
	CronSpec     string
	Concurrent   int
	Command      string
	Status       int
	Notify       int
	NotifyEmail  string
	Timeout      int
	ExecuteTimes int
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
	t.NotifyEmail = strings.TrimSpace(t.NotifyEmail)
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
		"update t_task set TaskName=:TaskName where TaskId=:TaskId",
		map[string]string{
			"TaskName": task.TaskName,
			"TaskId":   task.TaskId,
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
