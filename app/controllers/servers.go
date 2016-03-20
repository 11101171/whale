//****************************************************/
// https://github.com/11101171
// Author: ningzhong.zeng
//****************************************************/

package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
	"whale/app/models"
	"whale/app/routes"

	"github.com/revel/revel"
)

type Servers struct {
	Super
}

func (c Servers) Index(serverId string) revel.Result {
	user := c.SessionGetUser()
	servers := models.SelectServerListByUserId(user.UserId)
	var server = models.Server{}
	if serverId != "" {
		server = models.SelectServerOneByServerId(serverId)
	} else {
		if len(servers) > 0 {
			server = servers[0]
		}
	}

	jsonBody := models.Api{}
	// json.Unmarshal([]byte(server.Content), &jsonBody)
	json.Unmarshal([]byte(server.Content), &jsonBody)
	var admin bool = true
	return c.Render(servers, jsonBody, server, admin)
}

func (c Servers) List() revel.Result {
	user := c.SessionGetUser()
	servers := models.SelectServerListByUserId(user.UserId)
	return c.Render(servers)
}

func (c Servers) Operate(id string, server models.Server) revel.Result {
	if c.IsGet() {
		if id != "" {
			if server = models.SelectServerOneByServerId(id); server.ServerId == "" {
				return c.ErrDataBase(MsgSeleteError)
			}
			return c.Render(server)
		}
		return c.Render()
	} else {
		if server.ServerId != "" {
			server.ValidateUpdate(c.Validation)
			if c.Validation.HasErrors() {
				return c.Render(server)
			}
			if !models.UpdateServerOne(&server) {
				return c.ErrDataBase(MsgUpdateError)
			}
		} else {
			server.UserId = c.SessionGetUser().UserId
			server.ValidateInsert(c.Validation)
			if c.Validation.HasErrors() {
				return c.Render(server)
			}
			if err := models.DBMap().Insert(&server); err != nil {
				return c.ErrDataBase(MsgInsertError)
			}
		}
		return c.Redirect(routes.Servers.Index(server.ServerId))
	}
}

func (c Servers) Del(id string) revel.Result {
	if id != "" && models.DeleteServerOneByServerId(id) {
		return c.Redirect(routes.Servers.List())
	}
	return c.RenderJsonErr()
}

func (c Servers) Active(serverId string, apiParamId string, apiParam models.ApiParam) revel.Result {
	if c.IsGet() {
		server := models.SelectServerOneByServerId(serverId)
		if server.ServerId == "" {
			return c.ErrDataBase(MsgSeleteError)
		}

		apiParam := models.ParseApi(server.Content, apiParamId)
		return c.Render(apiParam, serverId, apiParamId)
	} else {
		var values = c.Params.Form
		var location = values["location[]"]
		values.Del("Host")
		values.Del("Method")
		values.Del("location[]")
		req, err := http.NewRequest(apiParam.Method, apiParam.Host, strings.NewReader(values.Encode()))
		if err != nil {
			revel.ERROR.Println("http req err =>", err)
			return c.RenderJsonErrWithData(map[string]interface{}{
				"resp": "错误" + err.Error(),
			})
		}
		for i := 0; i < len(location); i = i + 1 {
			if value := values.Get(location[i]); value != "" {
				req.Header.Add(location[i], value)
			}
		}
		client := &http.Client{
			Timeout: 10 * time.Second,
		}
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded; ")
		resp, err := client.Do(req)
		if err != nil {
			revel.ERROR.Println("http resp err =>", err)
			return c.RenderJsonErrWithData(map[string]interface{}{
				"resp": "错误" + err.Error(),
			})
		}
		// // client := &http.Client{Transport: transport}
		robots, _ := ioutil.ReadAll(resp.Body)
		return c.RenderJsonSucWithData(map[string]interface{}{
			"resp": string(robots),
		})
	}
}

func (c Servers) Share(serverId string) revel.Result {
	server := models.SelectServerOneByServerId(serverId)
	if server.ServerId == "" {
		return c.ErrDataBase(MsgSeleteError)
	}
	jsonBody := models.Api{}
	json.Unmarshal([]byte(server.Content), &jsonBody)
	return c.Render(server, jsonBody)
}

func (c Servers) Encode(serverId string, apiParamId string) revel.Result {
	server := models.SelectServerOneByServerId(serverId)
	if server.ServerId == "" {
		return c.ErrInputData(MsgSeleteError)
	}

	apiParam := models.ParseApi(server.Content, apiParamId)
	// values.Del("Host")
	// values.Del("Method")
	// json.Unmarshal([]byte(api), &apiParam)
	return c.RenderJsonSucWithData(map[string]interface{}{
		"resp": apiParam.Encode(c.Params.Form),
	})
}

func (c Servers) Info() revel.Result {
	return c.Render()
}

func init() {
	// revel.InterceptFunc(CheckLogin, revel.BEFORE, &Servers{})
}
