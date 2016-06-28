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

type Cmd struct {
	Base
	CmdId   string
	Shell   string
	AgentId string
}

func (c *Cmd) Validate(v *revel.Validation) {
	c.Shell = strings.TrimSpace(c.Shell)
	c.AgentId = strings.TrimSpace(c.AgentId)

	v.Check(c.AgentId,
		revel.Required{},
	).Key("cmd.AgentId").Message("AgentId不能为空")

	v.Check(c.Shell,
		revel.Required{},
	).Key("cmd.Shell").Message("脚本不能为空")
}

func (c *Cmd) ValidateInsert(v *revel.Validation) {
	c.Validate(v)
}

func (c *Cmd) ValidateUpdate(v *revel.Validation) {
	c.Validate(v)

	c.CmdId = strings.TrimSpace(c.CmdId)
	v.Check(c.CmdId,
		revel.Required{},
	).Key("cmd.CmdId").Message("CmdId不能为空")
}

func (c *Cmd) PreInsert(_ gorp.SqlExecutor) error {
	c.CmdId = RandomString("cmd", 6)
	return c.PreI()
}

func SelectCmdListByAgentId(agentId string) (cmds []Cmd) {
	CheckErrSQLSelectList(dbmap.Select(
		&cmds,
		"select * from o_cmd where AgentId=:AgentId order by GmtCreate desc",
		map[string]string{
			"AgentId": agentId,
		},
	))
	return cmds
}

func SelectCmdOneByCmdId(cmdId string) (cmd Cmd) {
	CheckErrSQLSelectOne(dbmap.SelectOne(
		&cmd,
		"select * from o_cmd where CmdId=:CmdId",
		map[string]string{
			"CmdId": cmdId,
		},
	))
	return cmd
}

func UpdateOneCmd(cmd *Cmd) bool {
	return Exec(
		"update o_cmd set Shell=:Shell where CmdId=:CmdId",
		map[string]string{
			"Shell": cmd.Shell,
			"CmdId": cmd.CmdId,
		},
	)
}

func DeleteOneCmdByCmdId(cmdId string) bool {
	return Exec(
		"delete from o_cmd where CmdId=:CmdId",
		map[string]string{
			"CmdId": cmdId,
		},
	)
}
