//****************************************************/
// https://github.com/11101171
// Author: ningzhong.zeng
//****************************************************/

package controllers

import (
	"fmt"
	"whale/app/models"
	"whale/app/routes"

	"github.com/revel/revel"
)

type Auth struct {
	Super
}

func (c Auth) Login(user models.User) revel.Result {
	if c.IsGet() {
		return c.Render()
	} else {
		if user.ValidateLogin(c.Validation); c.Validation.HasErrors() {
			return c.Render(user)
		}

		if user.Login(c.Validation) {
			c.SessionSetUser(user)
			return c.Redirect(routes.App.Index())
		} else {
			loginMsg := fmt.Sprintf("%s",
				c.Validation.ErrorMap()["Err"],
			)
			return c.Render(user, loginMsg)
		}
	}
}

func (c Auth) Logout() revel.Result {
	c.SessionClear()
	return c.Redirect(
		revel.MainRouter.Reverse("Auth.Login", map[string]string{}).Url,
	)
}
