// GENERATED CODE - DO NOT EDIT
package routes

import "github.com/revel/revel"


type tSuper struct {}
var Super tSuper


func (_ tSuper) ErrInputData(
		description string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "description", description)
	return revel.MainRouter.Reverse("Super.ErrInputData", args).Url
}

func (_ tSuper) ErrDataBase(
		description string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "description", description)
	return revel.MainRouter.Reverse("Super.ErrDataBase", args).Url
}

func (_ tSuper) RenderJsonSuc(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Super.RenderJsonSuc", args).Url
}

func (_ tSuper) RenderJsonErr(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Super.RenderJsonErr", args).Url
}


type tStatic struct {}
var Static tStatic


func (_ tStatic) Serve(
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.Serve", args).Url
}

func (_ tStatic) ServeModule(
		moduleName string,
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "moduleName", moduleName)
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.ServeModule", args).Url
}


type tUsers struct {}
var Users tUsers


func (_ tUsers) List(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Users.List", args).Url
}

func (_ tUsers) Operate(
		id string,
		user interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	revel.Unbind(args, "user", user)
	return revel.MainRouter.Reverse("Users.Operate", args).Url
}

func (_ tUsers) Del(
		id string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("Users.Del", args).Url
}


type tAgents struct {}
var Agents tAgents


func (_ tAgents) List(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Agents.List", args).Url
}

func (_ tAgents) Operate(
		id string,
		agent interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	revel.Unbind(args, "agent", agent)
	return revel.MainRouter.Reverse("Agents.Operate", args).Url
}

func (_ tAgents) Del(
		id string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("Agents.Del", args).Url
}

func (_ tAgents) SshList(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Agents.SshList", args).Url
}

func (_ tAgents) SshOne(
		agentId string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "agentId", agentId)
	return revel.MainRouter.Reverse("Agents.SshOne", args).Url
}

func (_ tAgents) SshJoin(
		agentId string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "agentId", agentId)
	return revel.MainRouter.Reverse("Agents.SshJoin", args).Url
}


type tApp struct {}
var App tApp


func (_ tApp) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.Index", args).Url
}


type tAuth struct {}
var Auth tAuth


func (_ tAuth) Login(
		user interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "user", user)
	return revel.MainRouter.Reverse("Auth.Login", args).Url
}

func (_ tAuth) Logout(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Auth.Logout", args).Url
}


type tCmds struct {}
var Cmds tCmds


func (_ tCmds) List(
		agentId string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "agentId", agentId)
	return revel.MainRouter.Reverse("Cmds.List", args).Url
}

func (_ tCmds) Operate(
		agentId string,
		id string,
		cmd interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "agentId", agentId)
	revel.Unbind(args, "id", id)
	revel.Unbind(args, "cmd", cmd)
	return revel.MainRouter.Reverse("Cmds.Operate", args).Url
}

func (_ tCmds) Del(
		id string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("Cmds.Del", args).Url
}


type tServers struct {}
var Servers tServers


func (_ tServers) Index(
		serverId string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "serverId", serverId)
	return revel.MainRouter.Reverse("Servers.Index", args).Url
}

func (_ tServers) List(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Servers.List", args).Url
}

func (_ tServers) Operate(
		id string,
		server interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	revel.Unbind(args, "server", server)
	return revel.MainRouter.Reverse("Servers.Operate", args).Url
}

func (_ tServers) Del(
		id string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("Servers.Del", args).Url
}

func (_ tServers) Active(
		serverId string,
		apiParamId string,
		apiParam interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "serverId", serverId)
	revel.Unbind(args, "apiParamId", apiParamId)
	revel.Unbind(args, "apiParam", apiParam)
	return revel.MainRouter.Reverse("Servers.Active", args).Url
}

func (_ tServers) Share(
		serverId string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "serverId", serverId)
	return revel.MainRouter.Reverse("Servers.Share", args).Url
}

func (_ tServers) Encode(
		serverId string,
		apiParamId string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "serverId", serverId)
	revel.Unbind(args, "apiParamId", apiParamId)
	return revel.MainRouter.Reverse("Servers.Encode", args).Url
}

func (_ tServers) Info(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Servers.Info", args).Url
}


