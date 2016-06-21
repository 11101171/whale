//****************************************************/
// Addres: https://github.com/11101171
// Author: ningzhong.zeng
//****************************************************/

package controllers

import (
	"whale/app/core/jobs"
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
	var taskLogs []models.TaskLog
	if groupId == "" {
		user := c.SessionGetUser()
		tasks = models.SelectTaskListByUserId(user.UserId)
		taskLogs = models.SelectTaskLogListBy(user.UserId, 0, 10)
	} else {
		tasks = models.SelectTaskListByGroupId(groupId)
	}

	return c.Render(tasks, taskLogs)
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

// 启动任务
func (c Tasks) TaskStart(id string) revel.Result {
	task := models.SelectTaskOneById(id)

	if task.TaskId == "" {
		return c.ErrDataBase(MsgSeleteError)
	}

	job, err := jobs.NewJobFromTask(&task)
	if err != nil {
		return c.ErrInputData("加入任务队列失败")
	}
	if jobs.AddJob(task.CronSpec, job) {
		revel.INFO.Println("加入job成功")
		task.Status = 1
		models.UpdateTaskOne(&task)
	}
	return c.Redirect(routes.Tasks.TaskList(""))
}

// 关闭任务
func (c Tasks) TaskStop(id string) revel.Result {
	task := models.SelectTaskOneById(id)

	if task.TaskId == "" {
		return c.ErrDataBase(MsgSeleteError)
	}

	if jobs.RemoveJob(task.TaskId) {
		revel.INFO.Println("关闭job成功")
		task.Status = 0
		models.UpdateTaskOne(&task)
	}
	return c.Redirect(routes.Tasks.TaskList(""))
}

// 查看log
func (c Tasks) TaskLog(id string) revel.Result {
	taskLog := models.SelectTaskLogOneBy(id)
	if taskLog.TaskLogId == "" {
		return c.ErrDataBase(MsgSeleteError)
	}

	task := models.SelectTaskOneById(taskLog.TaskId)
	if task.TaskId == "" {
		return c.ErrDataBase(MsgSeleteError)
	}

	return c.Render(taskLog, task)
}

// ajax
func (c Tasks) AjaxTaskLogList(page int, pageSize int) revel.Result {
	user := c.SessionGetUser()
	taskLogs := models.SelectTaskLogListBy(user.UserId, page, pageSize)
	return c.RenderJsonSucWithData(map[string]interface{}{
		"taskLogs": taskLogs,
	})
}

func init() {
	revel.InterceptFunc(CheckLogin, revel.BEFORE, &Tasks{})
}
