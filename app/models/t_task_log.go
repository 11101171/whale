//****************************************************/
// Addres: https://github.com/11101171
// Author: ningzhong.zeng
//****************************************************/

package models

import "github.com/go-gorp/gorp"

type TaskLog struct {
	Base
	TaskLogId   string
	TaskId      string
	TaskName    string
	UserId      string
	Output      string
	Error       string
	Status      int
	ProcessTime int
}

func (t *TaskLog) PreInsert(_ gorp.SqlExecutor) error {
	t.TaskLogId = RandomString("TaskLog", 6)
	return t.PreI()
}

func InsertTaskLogOne(taskLog *TaskLog) bool {
	err := dbmap.Insert(taskLog)
	return CheckErr(err, "Insert TaskLog One")
}

func SelectTaskLogListBy(userId string, page int, pageSize int) (taskLogs []TaskLog) {
	CheckErrSQLSelectList(dbmap.Select(
		&taskLogs,
		"select * from t_task_log where UserId=:UserId order by GmtCreate desc limit :page,:pageSize",
		map[string]interface{}{
			"UserId":   userId,
			"page":     page,
			"pageSize": pageSize,
		},
	))
	return taskLogs
}

func SelectTaskLogOneBy(taskLogId string) (taskLog TaskLog) {
	CheckErr(dbmap.SelectOne(
		&taskLog,
		"select * from t_task_log where TaskLogId=:TaskLogId",
		map[string]string{
			"TaskLogId": taskLogId,
		},
	), "Select TaskLog One")
	return taskLog
}

func SelectTaskLogOneByTaskId(taskId string) (taskLog TaskLog) {
	CheckErr(dbmap.SelectOne(
		&taskLog,
		"select * from t_task_log where TaskId=:TaskId",
		map[string]string{
			"TaskId": taskId,
		},
	), "Select TaskLog One")
	return taskLog
}

func UpdateTaskLogOne(taskLog *TaskLog) bool {
	return Exec(
		"update t_task_log set Output=:Output,Error=:Error,Status=:Status,ProcessTime=:ProcessTime where TaskId=:TaskId",
		map[string]interface{}{
			"TaskId":      taskLog.TaskId,
			"Output":      taskLog.Output,
			"Error":       taskLog.Error,
			"Status":      taskLog.Status,
			"ProcessTime": taskLog.ProcessTime,
		},
	)
}

func DeleteTaskLogOneByTaskId(taskId string) bool {
	return Exec(
		"delete from t_task_log where TaskId=:TaskId",
		map[string]string{
			"TaskId": taskId,
		},
	)
}
