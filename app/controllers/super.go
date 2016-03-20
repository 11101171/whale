//****************************************************/
// Author      : ningzhong.zeng
// Revision    : 2016-01-30 22:38:30
//****************************************************/

package controllers

import (
	"net/http"

	"whale/app/models"

	"github.com/revel/revel"
)

type Super struct {
	*revel.Controller
}

func (c *Super) IsGet() bool {
	return c.Request.Method == "Get" || c.Request.Method == "GET"
}

func (c *Super) IsPost() bool {
	return c.Request.Method == "Post" || c.Request.Method == "POST"
}

var (
	LOGIN_USERID   = "002"
	LOGIN_USERNAME = "001"
	LOGIN_USERROLE = "003"
)

func (c *Super) SessionSetUser(user models.User) {
	c.Session[LOGIN_USERID] = user.UserId
	c.Session[LOGIN_USERNAME] = user.Username
	c.Session[LOGIN_USERROLE] = user.GetRole()
}

func (c *Super) SessionGetUser() (user models.User) {
	user = models.User{
		UserId:   c.Session[LOGIN_USERID],
		Username: c.Session[LOGIN_USERNAME],
		Role:     models.Role(c.Session[LOGIN_USERROLE]),
	}
	return user
}

func (c *Super) SessionClear() {
	delete(c.Session, LOGIN_USERID)
	delete(c.Session, LOGIN_USERROLE)
	delete(c.Session, LOGIN_USERNAME)
}

// func init() {
// revel.InterceptFunc(CheckLogin, revel.BEFORE, &App{})
// }
func CheckLogin(c *revel.Controller) revel.Result {
	if c.Session[LOGIN_USERID] == "" {
		return c.Redirect(
			revel.MainRouter.Reverse("Auth.Login", map[string]string{}).Url,
		)
	}
	return nil
}

func CheckLoginAdmin(c *revel.Controller) revel.Result {
	if c.Session[LOGIN_USERID] == "" || models.Role(c.Session[LOGIN_USERROLE]) != models.ROLE_SUPER_ADMIN {
		return c.Redirect(
			revel.MainRouter.Reverse("Auth.Login", map[string]string{}).Url,
		)
	}
	return nil
}

const (
	errInputData = "输入数据有错误"
	errDataBase  = "数据库操作错误"
	suc          = "操作成功"
	err          = "操作失败"

	MsgInsertError = "保存失败"
	MsgSeleteError = "查找失败"
	MsgUpdateError = "更新失败"
	MsgDeleteError = "删除失败"
)

func (c *Super) ErrInputData(description string) revel.Result {
	c.Response.Status = http.StatusNotImplemented
	return c.RenderError(&revel.Error{
		Title:       errInputData,
		Description: description,
	})
}

func (c *Super) ErrDataBase(description string) revel.Result {
	c.Response.Status = http.StatusNotImplemented
	return c.RenderError(&revel.Error{
		Title:       errDataBase,
		Description: description,
	})
}

func (c *Super) RenderJsonSucWithData(data map[string]interface{}) revel.Result {
	return c.RenderJson(map[string]interface{}{
		"r_code": 1,
		"r_msg":  suc,
		"data":   data,
	})
}

func (c *Super) RenderJsonErrWithData(data map[string]interface{}) revel.Result {
	return c.RenderJson(map[string]interface{}{
		"r_code": 0,
		"r_msg":  err,
		"data":   data,
	})
}

func (c *Super) RenderJsonSuc() revel.Result {
	return c.RenderJson(map[string]interface{}{
		"r_code": 1,
		"r_msg":  suc,
	})
}

func (c *Super) RenderJsonErr() revel.Result {
	return c.RenderJson(map[string]interface{}{
		"r_code": 0,
		"r_msg":  err,
	})
}

func (c *Super) HasErr(err error) (r bool) {
	r = false
	if err != nil {
		revel.ERROR.Println(err)
		r = true
	}
	return r
}

func (c *Super) HasErrRows(rows int64, err error) (r bool) {
	r = false
	if err != nil {
		revel.ERROR.Println(rows, ",", err)
		r = true
	}
	return r
}
