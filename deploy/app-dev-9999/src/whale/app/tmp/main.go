// GENERATED CODE - DO NOT EDIT
package main

import (
	"flag"
	"reflect"
	"github.com/revel/revel"
	controllers0 "github.com/revel/modules/static/app/controllers"
	_ "whale/app"
	controllers "whale/app/controllers"
	models "whale/app/models"
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
	
	revel.RegisterController((*controllers0.Static)(nil),
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
				Name: "Del",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "id", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Operate",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "id", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "user", Type: reflect.TypeOf((*models.User)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					38: []string{ 
						"user",
					},
					40: []string{ 
					},
					45: []string{ 
						"user",
					},
					53: []string{ 
						"user",
					},
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
					21: []string{ 
						"agents",
					},
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
			
		})
	
	revel.RegisterController((*controllers.Servers)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "serverId", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					69: []string{ 
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
					75: []string{ 
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
					84: []string{ 
						"server",
					},
					86: []string{ 
					},
					91: []string{ 
						"server",
					},
					101: []string{ 
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
					139: []string{ 
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
					184: []string{ 
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
					203: []string{ 
					},
				},
			},
			
		})
	
	revel.DefaultValidationKeys = map[string]map[int]string{ 
		"whale/app/models.(*Agent).Validate": { 
			39: "a.Host",
			43: "a.Port",
			47: "a.LoginName",
			51: "a.LoginPass",
			55: "a.Memo",
		},
		"whale/app/models.(*Agent).ValidateUpdate": { 
			66: "a.AgentId",
		},
		"whale/app/models.(*Server).Validate": { 
			32: "server.Theme",
			37: "server.Content",
		},
		"whale/app/models.(*Server).ValidateUpdate": { 
			48: "server.ServerId",
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
	}

	revel.Run(*port)
}
