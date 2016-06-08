//****************************************************/
// https://github.com/11101171
// Author: ningzhong.zeng
//****************************************************/

package models

import (
	"strings"

	"github.com/go-gorp/gorp"
	"github.com/revel/revel"
)

type Agent struct {
	Base
	AgentId      string
	Host         string
	Port         string
	LoginName    string
	LoginPass    string
	InitShellCmd string
	Memo         string
	UserId       string
}

func (a *Agent) Validate(v *revel.Validation) {
	a.Host = strings.TrimSpace(a.Host)
	a.Port = strings.TrimSpace(a.Port)
	a.LoginName = strings.TrimSpace(a.LoginName)
	a.LoginPass = strings.TrimSpace(a.LoginPass)
	a.InitShellCmd = strings.TrimSpace(a.InitShellCmd)
	a.Memo = strings.TrimSpace(a.Memo)

	v.Check(a.Host,
		revel.Required{},
		revel.MinSize{Min: 1},
		revel.MaxSize{Max: 100},
	).Key("agent.Host").Message("主机需要大于1字符小于100个字符")
	v.Check(a.Port,
		revel.MinSize{Min: 1},
		revel.MaxSize{Max: 4},
	).Key("agent.Port").Message("端口要大于1个字符小于4个字符")
	v.Check(a.LoginName,
		revel.MinSize{Min: 1},
		revel.MaxSize{Max: 30},
	).Key("agent.LoginName").Message("登录用户名大于1个字符小于30个字符")

	v.Check(a.Memo,
		revel.MinSize{Min: 1},
		revel.MaxSize{Max: 50},
	).Key("agent.Memo").Message("备注大于1个字符小于50个字符")
}

func (a *Agent) ValidateInsert(v *revel.Validation) {
	v.Check(a.LoginPass,
		revel.MinSize{Min: 1},
		revel.MaxSize{Max: 30},
	).Key("agent.LoginPass").Message("登录密码大于1个字符小于30个字符")
	a.Validate(v)
}

func (a *Agent) ValidateUpdate(v *revel.Validation) {
	a.Validate(v)
	v.Check(a.AgentId,
		revel.Required{},
	).Key("agent.AgentId").Message("更新需要主键")
}

func (a *Agent) PreInsert(_ gorp.SqlExecutor) error {
	a.AgentId = RandomString("agent", 6)
	a.LoginPass = desUtil.Encrypt(a.LoginPass)
	return a.PreI()
}

func (a Agent) GetCmds() (cmds []Cmd) {
	CheckErrSQLSelectList(dbmap.Select(
		&cmds,
		"select * from o_cmd where AgentId=:AgentId order by GmtCreate desc",
		map[string]string{
			"AgentId": a.AgentId,
		},
	))
	return cmds
}

func SelectAgentListByUserId(userId string) (agents []Agent) {
	CheckErrSQLSelectList(dbmap.Select(
		&agents,
		"select * from o_agent where UserId=:UserId order by GmtCreate desc",
		map[string]string{
			"UserId": userId,
		},
	))
	return agents
}

func SelectAgentOneByAgentId(agentId string) (agent Agent) {
	CheckErrSQLSelectOne(dbmap.SelectOne(
		&agent,
		"select * from o_agent where AgentId=:AgentId",
		map[string]string{
			"AgentId": agentId,
		},
	))
	return agent
}

func UpdateOneAgent(agent *Agent) bool {
	var sql string
	if agent.LoginPass != "" {
		sql = "update o_agent set Host=:Host,Port=:Port,LoginName=:LoginName,LoginPass=:LoginPass,InitShellCmd=:InitShellCmd,Memo=:Memo where AgentId=:AgentId"
	} else {
		sql = "update o_agent set Host=:Host,Port=:Port,LoginName=:LoginName,InitShellCmd=:InitShellCmd,Memo=:Memo where AgentId=:AgentId"
	}
	return Exec(
		sql,
		map[string]string{
			"Host":         agent.Host,
			"Port":         agent.Port,
			"LoginName":    agent.LoginName,
			"LoginPass":    desUtil.Encrypt(agent.LoginPass),
			"InitShellCmd": agent.InitShellCmd,
			"Memo":         agent.Memo,
			"AgentId":      agent.AgentId,
		},
	)
}

func DeleteOneAgentByAgentId(agentId string) bool {
	return Exec(
		"delete from o_agent where AgentId=:agentId",
		map[string]string{
			"agentId": agentId,
		},
	)
}
