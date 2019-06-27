package mrp

import (
	"github.com/hexya-addons/web/controllers"
	"github.com/hexya-erp/hexya/src/server"
)

const MODULE_NAME string = "mrp"

func init() {
	server.RegisterModule(&server.Module{
		Name:     MODULE_NAME,
		PreInit:  func() {},
		PostInit: func() {},
	})
	controllers.BackendLess = append(controllers.BackendLess,
		"/static/mrp/src/less/mrp.less",
	)
	controllers.BackendJS = append(controllers.BackendJS,
		"/static/mrp/src/js/mrp.js",
	)

}
