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

type Cmds struct {
	Super
}

func (c Cmds) List(agentId string) revel.Result {
	if agentId == "" {
		c.ErrInputData(MsgSeleteError)
	}
	cmds := models.SelectCmdListByAgentId(agentId)
	return c.Render(cmds)
}

func (c Cmds) Operate(agentId string, id string, cmd models.Cmd) revel.Result {
	if c.IsGet() {
		if id != "" {
			if cmd = models.SelectCmdOneByCmdId(id); cmd.CmdId == "" {
				return c.ErrDataBase(MsgSeleteError)
			}
			return c.Render(cmd, agentId)
		}
		return c.Render(agentId)
	} else {
		if agentId == "" {
			return c.ErrInputData(errInputData)
		}
		cmd.AgentId = agentId
		if cmd.CmdId != "" {
			cmd.ValidateUpdate(c.Validation)
			if c.Validation.HasErrors() {
				return c.Render(cmd, agentId)
			}
			if !models.UpdateOneCmd(&cmd) {
				return c.ErrDataBase(MsgUpdateError)
			}
		} else {
			cmd.ValidateInsert(c.Validation)
			if c.Validation.HasErrors() {
				return c.Render(cmd, agentId)
			}
			if err := models.DBMap().Insert(&cmd); err != nil {
				return c.ErrDataBase(MsgInsertError)
			}
		}
		return c.Redirect(routes.Agents.List())
	}
}

func (c Cmds) Del(id string) revel.Result {
	if id != "" && models.DeleteOneCmdByCmdId(id) {
		return c.Redirect(routes.Agents.List())
	}
	return c.RenderJsonErr()
}

func init() {
	revel.InterceptFunc(CheckLogin, revel.BEFORE, &Cmds{})
}
