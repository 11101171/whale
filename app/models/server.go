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

type Server struct {
	Base
	ServerId string
	UserId   string
	Theme    string
	Content  string
}

func (server *Server) Validate(v *revel.Validation) {
	server.Theme = strings.TrimSpace(server.Theme)
	server.Content = strings.TrimSpace(server.Content)

	revel.INFO.Println(server.Content)
	v.Check(server.Theme,
		revel.Required{},
		revel.MinSize{Min: 1},
		revel.MaxSize{Max: 100},
	).Key("server.Theme").Message("主题需要大于1字符小于100个字符")
	v.Check(server.Content,
		revel.Required{},
		revel.MinSize{Min: 1},
		revel.MaxSize{Max: 30000},
	).Key("server.Content").Message("内容要大于1个字符小于30000个字符")
}

func (server *Server) ValidateInsert(v *revel.Validation) {
	server.Validate(v)
}

func (server *Server) ValidateUpdate(v *revel.Validation) {
	server.ValidateInsert(v)
	v.Check(server.ServerId,
		revel.Required{},
	).Key("server.ServerId").Message("更新需要主键")
}

func (server *Server) PreInsert(_ gorp.SqlExecutor) error {
	server.ServerId = RandomString("server", 6)
	return server.PreI()
}

func SelectServerListByUserId(userId string) (servers []Server) {
	CheckErrSQLSelectList(dbmap.Select(
		&servers,
		"select * from p_server where UserId=:UserId order by GmtCreate desc",
		map[string]string{
			"UserId": userId,
		},
	))
	return servers
}

func SelectServerOneByServerId(serverId string) (server Server) {
	CheckErrSQLSelectOne(dbmap.SelectOne(
		&server,
		"select * from p_server where ServerId=:ServerId",
		map[string]string{
			"ServerId": serverId,
		},
	))
	return server
}

func UpdateOneServer(server *Server) bool {
	return Exec(
		"update p_server set Theme=:Theme, Content=:Content where ServerId=:ServerId",
		map[string]string{
			"Theme":    server.Theme,
			"Content":  server.Content,
			"ServerId": server.ServerId,
		},
	)
}

func DeleteOneServerByServerId(serverId string) bool {
	return Exec(
		"delete from p_server where ServerId=:ServerId",
		map[string]string{
			"ServerId": serverId,
		},
	)
}
