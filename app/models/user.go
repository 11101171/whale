//****************************************************/
// Author      : ningzhong.zeng
// Revision    : 2016-01-26 21:40:16
//****************************************************/
package models

import (
	"strings"

	"github.com/go-gorp/gorp"
	"github.com/revel/revel"
)

type User struct {
	Base
	UserId   string
	Username string
	Password string
	Nickname string
}

func (u *User) Validate(v *revel.Validation) {
	// v = &revel.Validation{}
	u.Username = strings.TrimSpace(u.Username)
	u.Password = strings.TrimSpace(u.Password)

	v.Check(u.Username,
		revel.Required{},
		revel.MinSize{Min: 6},
		revel.MaxSize{Max: 12},
	).Key("user.Username").Message("用户名需要大于6位小于12位")
	v.Check(u.Password,
		revel.Required{},
		revel.MinSize{Min: 6},
		revel.MaxSize{Max: 12},
	).Key("user.Password").Message("密码需要大于6位小于12位")
	v.Required(u.Nickname).Key("user.Nickname").Message("昵称不能为空")
}

func (u *User) ValidateInsert(v *revel.Validation) {
	u.Validate(v)
}

func (u *User) ValidateUpdate(v *revel.Validation) {
	u.Validate(v)
	v.Check(u.UserId,
		revel.Required{},
	).Key("user.UserId").Message("更新用户需要主键")
}

func (u *User) PreInsert(_ gorp.SqlExecutor) error {
	u.UserId = RandomString("user", 6)
	return u.PreI()
}

var sql_select_users = `
select * from s_user order by GmtCreate desc
`

func SelectUserList() (users []User) {
	_, err := dbmap.Select(&users, sql_select_users)
	CheckErr(err, "Select Users List")
	return users
}

var sql_select_by_id = `
select * from s_user where userId=:userId
`

func SelectUserOneByUserId(userId string) (user User) {
	err := dbmap.SelectOne(&user, sql_select_by_id, map[string]string{
		"userId": userId,
	})
	CheckErr(err, "Select User One")
	return user
}
