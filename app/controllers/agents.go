//****************************************************/
// https://github.com/11101171
// Author: ningzhong.zeng
//****************************************************/

package controllers

import (
	"whale/app/models"

	"github.com/revel/revel"
)

type Agents struct {
	Super
}

func (c Agents) List() revel.Result {
	user := c.SessionGetUser()
	agents := models.SelectAgentListByUserId(user.UserId)
	return c.Render(agents)
}

func init() {
	revel.InterceptFunc(CheckLogin, revel.BEFORE, &Agents{})
}
