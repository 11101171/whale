//****************************************************/
// Author      : ningzhong.zeng
// Revision    : 2016-01-28 16:02:15
//****************************************************/

package controllers

import (
	"whale/app/models"
	"whale/app/routes"

	"github.com/revel/revel"
)

type Users struct {
	Super
}

func (c Users) List() revel.Result {
	users := models.SelectUserList()
	return c.Render(users)
}

// func (c Users) Operate(userId string) revel.Result {
// revel.INFO.Println("userId=>", userId)
// return c.Render()
// }

func (c Users) Del(id string) revel.Result {
	revel.INFO.Println("delete", id)
	return c.ErrDataBase("This a test method")
}

func (c Users) Edit(id string) revel.Result {
	if c.IsGet() {
		revel.INFO.Println("id", id)
		if id != "" {
			user := models.SelectUserOneByUserId(id)
			c.RenderArgs["user"] = user
		}
		return c.Redirect(routes.Users.List())
	} else {
		// user.ValidateInsert(c.Validation)
		// if c.Validation.HasErrors() {
		// c.Validation.Keep()
		// c.FlashParams()
		// return c.Redirect(routes.Users.Operate())
		// }
		// if c.HasErr(models.DBMap().Insert(&user)) {
		// return c.ErrDataBase("保存用户失败")
		// }
		return c.Redirect(routes.Users.List())
	}
}

func (c Users) Add(user models.User) revel.Result {
	if c.IsGet() {
		return c.Redirect(routes.Users.List())
	} else {
		user.ValidateInsert(c.Validation)
		if c.Validation.HasErrors() {
			c.Validation.Keep()
			c.FlashParams()
			return c.Redirect(routes.Users.List())
		}
		if c.HasErr(models.DBMap().Insert(&user)) {
			return c.ErrDataBase("保存用户失败")
		}
		return c.Redirect(routes.Users.List())
	}
}

func (c Users) Operate(id string, user models.User) revel.Result {
	if c.IsGet() {
		revel.INFO.Println("id", id)
		if id != "" {
			user := models.SelectUserOneByUserId(id)
			return c.Render(user)
		}
		return c.Render()
	} else {

		user.ValidateInsert(c.Validation)
		if c.Validation.HasErrors() {
			return c.Render()
		}
		if c.HasErr(models.DBMap().Insert(&user)) {
			return c.ErrDataBase("保存用户失败")
		}
		return c.Redirect(routes.Users.List())
	}
}
