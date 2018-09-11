package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["gitlab.com/manuel.diaz/sirel/server/api/controllers/public_controllers:LoginController"] = append(beego.GlobalControllerRouter["gitlab.com/manuel.diaz/sirel/server/api/controllers/public_controllers:LoginController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["gitlab.com/manuel.diaz/sirel/server/api/controllers/public_controllers:PublicAreasController"] = append(beego.GlobalControllerRouter["gitlab.com/manuel.diaz/sirel/server/api/controllers/public_controllers:PublicAreasController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/area`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["gitlab.com/manuel.diaz/sirel/server/api/controllers/public_controllers:PublicAreasController"] = append(beego.GlobalControllerRouter["gitlab.com/manuel.diaz/sirel/server/api/controllers/public_controllers:PublicAreasController"],
		beego.ControllerComments{
			Method: "List",
			Router: `/areas`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["gitlab.com/manuel.diaz/sirel/server/api/controllers/public_controllers:PublicLocalsController"] = append(beego.GlobalControllerRouter["gitlab.com/manuel.diaz/sirel/server/api/controllers/public_controllers:PublicLocalsController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/local`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["gitlab.com/manuel.diaz/sirel/server/api/controllers/public_controllers:PublicLocalsController"] = append(beego.GlobalControllerRouter["gitlab.com/manuel.diaz/sirel/server/api/controllers/public_controllers:PublicLocalsController"],
		beego.ControllerComments{
			Method: "List",
			Router: `/locals`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}