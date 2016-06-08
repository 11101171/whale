package app

import (
	"fmt"
	"strings"
	"whale/app/models"

	"github.com/revel/revel"
)

func init() {
	// Filters is the default set of global filters.
	revel.Filters = []revel.Filter{
		revel.PanicFilter,             // Recover from panics and display an error page instead.
		revel.RouterFilter,            // Use the routing table to select the right Action
		revel.FilterConfiguringFilter, // A hook for adding or removing per-Action filters.
		revel.ParamsFilter,            // Parse parameters into Controller.Params.
		revel.SessionFilter,           // Restore and write the session cookie.
		revel.FlashFilter,             // Restore and write the flash cookie.
		revel.ValidationFilter,        // Restore kept validation errors and save new ones from cookie.
		revel.I18nFilter,              // Resolve the requested language
		HeaderFilter,                  // Add some security based headers
		revel.InterceptorFilter,       // Run interceptors around the action.
		revel.CompressFilter,          // Compress the result.
		ActionInvoker,                 // Invoke the action.
	}
	// revel.TimeFormats = append(revel.TimeFormats, "2006/01/01")
	// ( order dependent )
	// revel.OnAppStart(InitDB)
	// revel.OnAppStart(FillCache)
	revel.OnAppStart(func() {
		models.InitDB()
	})
}

// TODO turn this into revel.HeaderFilter
// should probably also have a filter for CSRF
// not sure if it can go in the same filter or not
var HeaderFilter = func(c *revel.Controller, fc []revel.Filter) {
	// Add some common security headers
	c.Response.Out.Header().Add("X-Frame-Options", "SAMEORIGIN")
	c.Response.Out.Header().Add("X-XSS-Protection", "1; mode=block")
	c.Response.Out.Header().Add("X-Content-Type-Options", "nosniff")

	fc[0](c, fc[1:]) // Execute the next filter stage.
}

// Remember Records
var ActionInvoker = func(c *revel.Controller, fc []revel.Filter) {
	if c.MethodName != "Serve" {
		revel.INFO.Printf("{data} Method[%s] Action[%s] Params[%s] Who[%s]",
			c.Request.Method,
			c.Action,
			c.Params.Form,
			c.Session["001"],
		)
		// ToString("Args=>", c.Args)
		// ToString("Flash=>", c.Flash)
		// ToString("================")
	}
	revel.ActionInvoker(c, fc)
}

func ToString(v ...interface{}) {
	fmt.Printf(strings.Repeat("%+v ", len(v)), v)
	fmt.Println()
}
