//****************************************************/
// Addres: https://github.com/11101171
// Author: ningzhong.zeng
//****************************************************/

package controllers

import (
	"whale/app/models"
	"whale/app/routes"

	"github.com/revel/revel"
)

type Tasks struct {
	Super
}

func (c Tasks) GroupList() revel.Result {
	user := c.SessionGetUser()
	taskGroups := models.SelectTaskGroupListByUserId(user.UserId)
	return c.Render(taskGroups)
}

func (c Tasks) GroupOperate(id string, taskGroup models.TaskGroup) revel.Result {
	if c.IsGet() {
		if id != "" {
			if taskGroup = models.SelectTaskGroupOneById(id); taskGroup.TaskGroupId == "" {
				return c.ErrDataBase(MsgSeleteError)
			}
			return c.Render(taskGroup)
		}
		return c.Render()
	} else {
		if taskGroup.TaskGroupId != "" {
			taskGroup.ValidateUpdate(c.Validation)
			if c.Validation.HasErrors() {
				return c.Render(taskGroup)
			}
			if !models.UpdateTaskGroupOne(&taskGroup) {
				return c.ErrDataBase(MsgUpdateError)
			}
		} else {
			taskGroup.UserId = c.SessionGetUser().UserId
			taskGroup.ValidateInsert(c.Validation)
			if c.Validation.HasErrors() {
				return c.Render(taskGroup)
			}
			if !models.InsertTaskGroupOne(&taskGroup) {
				return c.ErrDataBase(MsgInsertError)
			}
		}
		return c.Redirect(routes.Tasks.GroupList())
	}
}

func (c Tasks) GroupDel(id string) revel.Result {
	if id != "" && models.DeleteTaskGroupOneById(id) {
		return c.Redirect(routes.Tasks.GroupList())
	}
	return c.RenderJsonErr()
}

func (c Tasks) TaskList(groupId string) revel.Result {
	var tasks []models.Task
	if groupId == "" {
		user := c.SessionGetUser()
		tasks = models.SelectTaskListByUserId(user.UserId)
	} else {
		tasks = models.SelectTaskListByGroupId(groupId)
	}
	return c.Render(tasks)
}

func (c Tasks) TaskOperate(id string, task models.Task) revel.Result {
	group := models.SelectTaskGroupListByUserId(c.SessionGetUser().UserId)
	if c.IsGet() {
		if id != "" {
			if task = models.SelectTaskOneById(id); task.TaskId == "" {
				return c.ErrDataBase(MsgSeleteError)
			}
			return c.Render(task, group)
		}
		return c.Render(group)
	} else {
		if task.TaskId != "" {
			task.ValidateUpdate(c.Validation)
			if c.Validation.HasErrors() {
				return c.Render(task, group)
			}
			if !models.UpdateTaskOne(&task) {
				return c.ErrDataBase(MsgUpdateError)
			}
		} else {
			task.UserId = c.SessionGetUser().UserId
			task.ValidateInsert(c.Validation)
			if c.Validation.HasErrors() {
				return c.Render(task, group)
			}
			if err := models.DBMap().Insert(&task); err != nil {
				return c.ErrDataBase(MsgInsertError)
			}
		}
		return c.Redirect(routes.Tasks.TaskList(task.GroupId))
	}
}

func init() {
	revel.InterceptFunc(CheckLogin, revel.BEFORE, &Tasks{})
}
