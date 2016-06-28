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

type TaskGroup struct {
	Base
	TaskGroupId string
	UserId      string
	GroupName   string
	Description string
}

func (tg *TaskGroup) Validate(v *revel.Validation) {
	tg.GroupName = strings.TrimSpace(tg.GroupName)
	tg.Description = strings.TrimSpace(tg.Description)

	v.Check(tg.GroupName,
		revel.Required{},
	).Key("taskGroup.GroupName").Message("组名不能为空")
	v.Check(tg.Description,
		revel.Required{},
	).Key("taskGroup.Description").Message("描述不能为空")
}

func (tg *TaskGroup) ValidateInsert(v *revel.Validation) {
	tg.Validate(v)
}

func (tg *TaskGroup) ValidateUpdate(v *revel.Validation) {
	tg.ValidateInsert(v)
}

func (tg *TaskGroup) PreInsert(_ gorp.SqlExecutor) error {
	tg.TaskGroupId = RandomString("g", 6)
	return tg.PreI()
}

func InsertTaskGroupOne(taskGroup *TaskGroup) bool {
	err := dbmap.Insert(taskGroup)
	return CheckErr(err, "Insert taskGroup One =>")
}

func SelectTaskGroupList() (taskGroups []TaskGroup) {
	CheckErrSQLSelectList(dbmap.Select(
		&taskGroups,
		"select * from t_task_group order by GmtCreate desc",
	))
	return taskGroups
}

func SelectTaskGroupListByUserId(userId string) (taskGroups []TaskGroup) {
	CheckErrSQLSelectList(dbmap.Select(
		&taskGroups,
		"select * from t_task_group where UserId=:UserId order by GmtCreate desc",
		map[string]string{
			"UserId": userId,
		},
	))
	return taskGroups
}

func SelectTaskGroupOneById(taskGroupId string) (taskGroup TaskGroup) {
	CheckErr(dbmap.SelectOne(
		&taskGroup,
		"select * from t_task_group where TaskGroupId=:TaskGroupId",
		map[string]string{
			"TaskGroupId": taskGroupId,
		},
	), "Select TaskGroup One =>")
	return taskGroup
}

func UpdateTaskGroupOne(taskGroup *TaskGroup) bool {
	return Exec(
		"update t_task_group set GroupName=:GroupName,Description=:Description where TaskGroupId=:TaskGroupId",
		map[string]string{
			"GroupName":   taskGroup.GroupName,
			"Description": taskGroup.Description,
			"TaskGroupId": taskGroup.TaskGroupId,
		},
	)
}

func DeleteTaskGroupOneById(taskGroupId string) bool {
	return Exec(
		"delete from t_task_group where TaskGroupId=:TaskGroupId",
		map[string]string{
			"TaskGroupId": taskGroupId,
		},
	)
}
