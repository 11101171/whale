// GENERATED CODE - DO NOT EDIT
package main

import (
	"flag"
	"reflect"
	"github.com/revel/revel"
	controllers1 "github.com/revel/modules/static/app/controllers"
	_ "github.com/revel/modules/testrunner/app"
	controllers0 "github.com/revel/modules/testrunner/app/controllers"
	_ "whale/app"
	controllers "whale/app/controllers"
	_ "whale/app/core"
	models "whale/app/models"
	tests "whale/tests"
	"github.com/revel/revel/testing"
)

var (
	runMode    *string = flag.String("runMode", "", "Run mode.")
	port       *int    = flag.Int("port", 0, "By default, read from app.conf")
	importPath *string = flag.String("importPath", "", "Go Import Path for the app.")
	srcPath    *string = flag.String("srcPath", "", "Path to the source root.")

	// So compiler won't complain if the generated code doesn't reference reflect package...
	_ = reflect.Invalid
)

func main() {
	flag.Parse()
	revel.Init(*runMode, *importPath, *srcPath)
	revel.INFO.Println("Running revel server")
	
	revel.RegisterController((*controllers.Super)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "ErrInputData",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "description", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "ErrDataBase",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "description", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "RenderJsonSuc",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "RenderJsonErr",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers0.TestRunner)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					70: []string{ 
						"testSuites",
					},
				},
			},
			&revel.MethodType{
				Name: "Run",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "suite", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "test", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					107: []string{ 
					},
				},
			},
			&revel.MethodType{
				Name: "List",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers1.Static)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Serve",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "prefix", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "filepath", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "ServeModule",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "moduleName", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "prefix", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "filepath", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers.Users)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "List",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					21: []string{ 
						"users",
					},
				},
			},
			&revel.MethodType{
				Name: "Operate",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "id", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "user", Type: reflect.TypeOf((*models.User)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					30: []string{ 
						"user",
					},
					32: []string{ 
					},
					37: []string{ 
						"user",
					},
					45: []string{ 
						"user",
					},
				},
			},
			&revel.MethodType{
				Name: "Del",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "id", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers.Agents)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "List",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					22: []string{ 
						"agents",
					},
				},
			},
			&revel.MethodType{
				Name: "Operate",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "id", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "agent", Type: reflect.TypeOf((*models.Agent)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					31: []string{ 
						"agent",
					},
					33: []string{ 
					},
					38: []string{ 
						"agent",
					},
					47: []string{ 
						"agent",
					},
				},
			},
			&revel.MethodType{
				Name: "Del",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "id", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "SshList",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					67: []string{ 
						"agents",
					},
				},
			},
			&revel.MethodType{
				Name: "SshOne",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "agentId", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					75: []string{ 
						"agent",
					},
				},
			},
			&revel.MethodType{
				Name: "SshJoin",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "agentId", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers.App)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					10: []string{ 
					},
				},
			},
			
		})
	
	revel.RegisterController((*controllers.Auth)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Login",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "user", Type: reflect.TypeOf((*models.User)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					22: []string{ 
					},
					25: []string{ 
						"user",
					},
					35: []string{ 
						"user",
						"loginMsg",
					},
				},
			},
			&revel.MethodType{
				Name: "Logout",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers.Cmds)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "List",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "agentId", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					24: []string{ 
						"cmds",
					},
				},
			},
			&revel.MethodType{
				Name: "Operate",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "agentId", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "id", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "cmd", Type: reflect.TypeOf((*models.Cmd)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					33: []string{ 
						"cmd",
						"agentId",
					},
					35: []string{ 
						"agentId",
					},
					44: []string{ 
						"cmd",
						"agentId",
					},
					52: []string{ 
						"cmd",
						"agentId",
					},
				},
			},
			&revel.MethodType{
				Name: "Del",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "id", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers.Servers)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "serverId", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					40: []string{ 
						"servers",
						"jsonBody",
						"server",
						"admin",
					},
				},
			},
			&revel.MethodType{
				Name: "List",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					46: []string{ 
						"servers",
					},
				},
			},
			&revel.MethodType{
				Name: "Operate",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "id", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "server", Type: reflect.TypeOf((*models.Server)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					55: []string{ 
						"server",
					},
					57: []string{ 
					},
					62: []string{ 
						"server",
					},
					71: []string{ 
						"server",
					},
				},
			},
			&revel.MethodType{
				Name: "Del",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "id", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Active",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "serverId", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "apiParamId", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "apiParam", Type: reflect.TypeOf((*models.ApiParam)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					96: []string{ 
						"apiParam",
						"serverId",
						"apiParamId",
					},
				},
			},
			&revel.MethodType{
				Name: "Share",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "serverId", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					141: []string{ 
						"server",
						"jsonBody",
					},
				},
			},
			&revel.MethodType{
				Name: "Encode",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "serverId", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "apiParamId", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Info",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					160: []string{ 
					},
				},
			},
			
		})
	
	revel.RegisterController((*controllers.Tasks)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "GroupList",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					24: []string{ 
						"taskGroups",
					},
				},
			},
			&revel.MethodType{
				Name: "GroupOperate",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "id", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "taskGroup", Type: reflect.TypeOf((*models.TaskGroup)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					33: []string{ 
						"taskGroup",
					},
					35: []string{ 
					},
					40: []string{ 
						"taskGroup",
					},
					49: []string{ 
						"taskGroup",
					},
				},
			},
			&revel.MethodType{
				Name: "GroupDel",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "id", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "TaskList",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "groupId", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					77: []string{ 
						"tasks",
						"taskLogs",
					},
				},
			},
			&revel.MethodType{
				Name: "TaskOperate",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "id", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "task", Type: reflect.TypeOf((*models.Task)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					87: []string{ 
						"task",
						"group",
					},
					89: []string{ 
						"group",
					},
					94: []string{ 
						"task",
						"group",
					},
					103: []string{ 
						"task",
						"group",
					},
				},
			},
			&revel.MethodType{
				Name: "TaskDel",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "id", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "TaskStart",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "id", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "TaskStop",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "id", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "TaskLog",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "id", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					168: []string{ 
						"taskLog",
						"task",
					},
				},
			},
			&revel.MethodType{
				Name: "AjaxTaskLogList",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "page", Type: reflect.TypeOf((*int)(nil)) },
					&revel.MethodArg{Name: "pageSize", Type: reflect.TypeOf((*int)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Send",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.DefaultValidationKeys = map[string]map[int]string{ 
		"whale/app/models.(*Agent).Validate": { 
			39: "a.Host",
			43: "a.Port",
			47: "a.LoginName",
			52: "a.Memo",
		},
		"whale/app/models.(*Agent).ValidateInsert": { 
			59: "a.LoginPass",
		},
		"whale/app/models.(*Agent).ValidateUpdate": { 
			67: "a.AgentId",
		},
		"whale/app/models.(*Cmd).Validate": { 
			28: "c.AgentId",
			32: "c.Shell",
		},
		"whale/app/models.(*Cmd).ValidateUpdate": { 
			45: "c.CmdId",
		},
		"whale/app/models.(*Server).Validate": { 
			32: "server.Theme",
			37: "server.Content",
		},
		"whale/app/models.(*Server).ValidateUpdate": { 
			48: "server.ServerId",
		},
		"whale/app/models.(*Task).Validate": { 
			56: "t.TaskName",
			59: "t.Command",
		},
		"whale/app/models.(*TaskGroup).Validate": { 
			29: "tg.GroupName",
			32: "tg.Description",
		},
		"whale/app/models.(*User).Validate": { 
			68: "u.Username",
			73: "u.Password",
		},
		"whale/app/models.(*User).ValidateInsert": { 
			77: "u.Nickname",
		},
		"whale/app/models.(*User).ValidateUpdate": { 
			85: "u.UserId",
		},
	}
	testing.TestSuites = []interface{}{ 
		(*tests.AppTest)(nil),
	}

	revel.Run(*port)
}
