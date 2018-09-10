// @APIVersion 1.0.0
// @Title api-sirel
// @Description Documentation for api-sirel
// @Contact di@upr.edu.cu
// @License UPR
package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	"gitlab.com/manuel.diaz/sirel/server/api/controllers"
)

func init() {
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "HEAD", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "content-type", "Access-Control-Allow-Origin", "authHd"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
	}))

	ns := beego.NewNamespace("/api",
		beego.NSNamespace("/public",
			beego.NSInclude(
				&controllers.LoginController{},
				&controllers.PublicAreasController{},
			),
		),
		beego.NSNamespace("/private",
			beego.NSBefore(controllers.AuthFilter),
			beego.NSInclude(
				&controllers.LogoutController{},
				&controllers.ProfileController{},
			),
		),
		beego.NSNamespace("/admin",
			beego.NSBefore(controllers.AuthFilter),
			beego.NSBefore(controllers.AdminFilter),
			beego.NSInclude(
				&controllers.AdminUsersController{},
				&controllers.AdminAreasController{},
				&controllers.AdminLocalsController{},
			),
		),
	)
	beego.AddNamespace(ns)

	beego.ErrorController(&controllers.ErrorController{})
}
