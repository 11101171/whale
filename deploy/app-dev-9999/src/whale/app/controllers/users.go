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

func (c Users) Del(id string) revel.Result {
	if id != "" && models.DeleteUserOneByUserId(id) {
		// return c.RenderJsonSuc()
		return c.Redirect(routes.Users.List())
	}
	return c.RenderJsonErr()
}

func (c Users) Operate(id string, user models.User) revel.Result {
	if c.IsGet() {
		if id != "" {
			if user = models.SelectUserOneByUserId(id); user.UserId == "" {
				return c.ErrDataBase(MsgSeleteError)
			}
			return c.Render(user)
		}
		return c.Render()
	} else {
		if user.UserId != "" {
			user.ValidateUpdate(c.Validation)
			if c.Validation.HasErrors() {
				return c.Render(user)
			}
			if !models.UpdateUserOne(&user) {
				return c.ErrDataBase(MsgUpdateError)
			}
		} else {
			user.ValidateInsert(c.Validation)
			if c.Validation.HasErrors() {
				return c.Render(user)
			}
			if !models.InsertUserOne(&user) {
				return c.ErrDataBase(MsgInsertError)
			}
		}
		return c.Redirect(routes.Users.List())
	}
}
