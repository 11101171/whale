//****************************************************/
// https://github.com/11101171
// Author: ningzhong.zeng
//****************************************************/
package models

import (
	"crypto/md5"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/go-gorp/gorp"
	"github.com/revel/revel"
)

type RoleType int

const (
	ROLE_SUPER_ADMIN RoleType = iota
	ROLE_ADMIN
	ROLE_GUEST
)

func Role(v string) RoleType {
	vi, _ := strconv.Atoi(v)
	switch vi {
	case 0:
		return ROLE_SUPER_ADMIN
	case 1:
		return ROLE_ADMIN
	case 2:
		return ROLE_GUEST
	default:
		return -1
	}
}

type User struct {
	Base
	UserId   string
	Username string
	Password string
	Nickname string
	Role     RoleType
}

func (u *User) GetRole() string {
	switch {
	case u.Role == 0:
		return "超级管理员"
	case u.Role == 1:
		return "管理员"
	default:
		return "游客"
	}
}

func (u *User) Validate(v *revel.Validation) {
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
}

func (u *User) ValidateInsert(v *revel.Validation) {
	v.Required(u.Nickname).Key("user.Nickname").Message("昵称不能为空")
	u.Validate(v)
}

func (u *User) ValidateUpdate(v *revel.Validation) {
	u.ValidateInsert(v)
	v.Check(u.UserId,
		revel.Required{},
	).Key("user.UserId").Message("更新用户需要主键")
}

func (u *User) ValidateLogin(v *revel.Validation) {
	u.Validate(v)
}

func (u *User) PreInsert(_ gorp.SqlExecutor) error {
	u.UserId = RandomString("user", 6)
	return u.PreI()
}

// services
func EnPwd(password string) string {
	hash := md5.New()
	io.WriteString(hash, password+"1d*;]O#")
	return fmt.Sprintf("%x", hash.Sum(nil))
}

func CheckPwd(password string, enPassword string) bool {
	return EnPwd(enPassword) == password
}

func (u *User) Login(v *revel.Validation) bool {
	selectUser := User{}
	err := dbmap.SelectOne(&selectUser,
		"select * from s_user where Username=:Username",
		map[string]string{
			"Username": u.Username,
		},
	)
	CheckErr(err, "Select One User For Login")
	if selectUser.Username == "" {
		ValidationAddErr(v, "登录失败，用户名不存在")
		return false
	}
	if CheckPwd(selectUser.Password, u.Password) {
		u.UserId = selectUser.UserId
		u.Role = selectUser.Role
		return true
	} else {
		ValidationAddErr(v, "登录失败，密码错误")
		return false
	}
}

func InsertUserOne(user *User) bool {
	user.Password = EnPwd(user.Password)
	user.Role = ROLE_ADMIN
	err := dbmap.Insert(user)
	return CheckErr(err, "Insert User One")
}

func SelectUserList() (users []User) {
	CheckErrSQLSelectList(dbmap.Select(
		&users,
		"select * from s_user order by GmtCreate desc",
	))
	return users
}

func SelectUserOneByUserId(userId string) (user User) {
	CheckErr(dbmap.SelectOne(
		&user,
		"select * from s_user where userId=:userId",
		map[string]string{
			"userId": userId,
		},
	), "Select User One")
	return user
}

func UpdateUserOne(user *User) bool {
	return Exec(
		"update s_user set Username=:Username, Password=:Password, Nickname=:Nickname where UserId=:UserId",
		map[string]string{
			"Username": user.Username,
			"Password": EnPwd(user.Password),
			"Nickname": user.Nickname,
			"UserId":   user.UserId,
		},
	)
}

func DeleteUserOneByUserId(userId string) bool {
	return Exec(
		"delete from s_user where userId=:userId",
		map[string]string{
			"userId": userId,
		},
	)
}
