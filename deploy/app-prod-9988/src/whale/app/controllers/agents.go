//****************************************************/
// https://github.com/11101171
// Author: ningzhong.zeng
//****************************************************/

package controllers

import (
	"whale/app/models"
	"whale/app/routes"

	"github.com/revel/revel"
)

type Agents struct {
	Super
}

func (c Agents) List() revel.Result {
	user := c.SessionGetUser()
	agents := models.SelectAgentListByUserId(user.UserId)
	return c.Render(agents)
}

func (c Agents) Operate(id string, agent models.Agent) revel.Result {
	if c.IsGet() {
		if id != "" {
			if agent = models.SelectAgentOneByAgentId(id); agent.AgentId == "" {
				return c.ErrDataBase(MsgSeleteError)
			}
			return c.Render(agent)
		}
		return c.Render()
	} else {
		if agent.AgentId != "" {
			agent.ValidateUpdate(c.Validation)
			if c.Validation.HasErrors() {
				return c.Render(agent)
			}
			if !models.UpdateOneAgent(&agent) {
				return c.ErrDataBase(MsgUpdateError)
			}
		} else {
			agent.UserId = c.SessionGetUser().UserId
			agent.ValidateInsert(c.Validation)
			if c.Validation.HasErrors() {
				return c.Render(agent)
			}
			if err := models.DBMap().Insert(&agent); err != nil {
				return c.ErrDataBase(MsgInsertError)
			}
		}
		return c.Redirect(routes.Agents.List())
	}
}

func (c Agents) Del(id string) revel.Result {
	if id != "" && models.DeleteOneAgentByAgentId(id) {
		return c.Redirect(routes.Agents.List())
	}
	return c.RenderJsonErr()
}

func (c Agents) SshList() revel.Result {
	user := c.SessionGetUser()
	agents := models.SelectAgentListByUserId(user.UserId)
	return c.Render(agents)
}

func (c Agents) SshOne(agentId string) revel.Result {
	agent := models.SelectAgentOneByAgentId(agentId)
	if agent.AgentId == "" {
		return c.ErrDataBase(MsgSeleteError)
	}
	return c.Render(agent)
}

func (c Agents) SshJoin(agentId string) revel.Result {
	var sshConfig = models.NewSSHConfig(agentId)
	if sshConfig.Agent.AgentId == "" {
		return c.RenderJsonErr()
	}
	models.Shell(c.Request.Websocket, sshConfig)
	return c.RenderJsonSuc()
}

func init() {
	revel.InterceptFunc(CheckLogin, revel.BEFORE, &Agents{})
}
