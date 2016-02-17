//****************************************************/
// Author      : ningzhong.zeng
// Revision    : 2016-01-30 22:38:30
//****************************************************/

package controllers

import (
	"net/http"

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

const (
	errInputData = "输入数据有错误"
	errDataBase  = "数据库操作错误"
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

func (c *Super) HasErr(err error) (r bool) {
	r = false
	if err != nil {
		revel.ERROR.Println(err)
		r = true
	}
	return r
}
