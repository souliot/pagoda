// @APIVersion 1.0.0
// @Title padoda
// @Description Pagoda is a rest api for ansible
// @Contact leizhou.lin@watrix.ai
// @TermsOfServiceUrl http://github.com/souliot/
package routers

import (
	"github.com/souliot/pagoda/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/host",
			beego.NSInclude(
				&controllers.HostController{},
			),
		),
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
		beego.NSNamespace("/component",
			beego.NSInclude(
				&controllers.MetaComponentController{},
			),
		),
		beego.NSNamespace("/properties",
			beego.NSInclude(
				&controllers.PropertyController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
