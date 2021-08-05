package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/souliot/pagoda/controllers:HostController"] = append(beego.GlobalControllerRouter["github.com/souliot/pagoda/controllers:HostController"],
        beego.ControllerComments{
            Method: "All",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/souliot/pagoda/controllers:HostController"] = append(beego.GlobalControllerRouter["github.com/souliot/pagoda/controllers:HostController"],
        beego.ControllerComments{
            Method: "Add",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/souliot/pagoda/controllers:HostController"] = append(beego.GlobalControllerRouter["github.com/souliot/pagoda/controllers:HostController"],
        beego.ControllerComments{
            Method: "Update",
            Router: "/",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/souliot/pagoda/controllers:HostController"] = append(beego.GlobalControllerRouter["github.com/souliot/pagoda/controllers:HostController"],
        beego.ControllerComments{
            Method: "DeleteMulti",
            Router: "/",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/souliot/pagoda/controllers:HostController"] = append(beego.GlobalControllerRouter["github.com/souliot/pagoda/controllers:HostController"],
        beego.ControllerComments{
            Method: "One",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/souliot/pagoda/controllers:HostController"] = append(beego.GlobalControllerRouter["github.com/souliot/pagoda/controllers:HostController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/souliot/pagoda/controllers:MetaComponentController"] = append(beego.GlobalControllerRouter["github.com/souliot/pagoda/controllers:MetaComponentController"],
        beego.ControllerComments{
            Method: "All",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/souliot/pagoda/controllers:MetaComponentController"] = append(beego.GlobalControllerRouter["github.com/souliot/pagoda/controllers:MetaComponentController"],
        beego.ControllerComments{
            Method: "Add",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/souliot/pagoda/controllers:MetaComponentController"] = append(beego.GlobalControllerRouter["github.com/souliot/pagoda/controllers:MetaComponentController"],
        beego.ControllerComments{
            Method: "Update",
            Router: "/",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/souliot/pagoda/controllers:MetaComponentController"] = append(beego.GlobalControllerRouter["github.com/souliot/pagoda/controllers:MetaComponentController"],
        beego.ControllerComments{
            Method: "DeleteMulti",
            Router: "/",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/souliot/pagoda/controllers:MetaComponentController"] = append(beego.GlobalControllerRouter["github.com/souliot/pagoda/controllers:MetaComponentController"],
        beego.ControllerComments{
            Method: "One",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/souliot/pagoda/controllers:MetaComponentController"] = append(beego.GlobalControllerRouter["github.com/souliot/pagoda/controllers:MetaComponentController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/souliot/pagoda/controllers:PropertyController"] = append(beego.GlobalControllerRouter["github.com/souliot/pagoda/controllers:PropertyController"],
        beego.ControllerComments{
            Method: "All",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/souliot/pagoda/controllers:PropertyController"] = append(beego.GlobalControllerRouter["github.com/souliot/pagoda/controllers:PropertyController"],
        beego.ControllerComments{
            Method: "Add",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/souliot/pagoda/controllers:PropertyController"] = append(beego.GlobalControllerRouter["github.com/souliot/pagoda/controllers:PropertyController"],
        beego.ControllerComments{
            Method: "Update",
            Router: "/",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/souliot/pagoda/controllers:PropertyController"] = append(beego.GlobalControllerRouter["github.com/souliot/pagoda/controllers:PropertyController"],
        beego.ControllerComments{
            Method: "DeleteMulti",
            Router: "/",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/souliot/pagoda/controllers:PropertyController"] = append(beego.GlobalControllerRouter["github.com/souliot/pagoda/controllers:PropertyController"],
        beego.ControllerComments{
            Method: "One",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/souliot/pagoda/controllers:PropertyController"] = append(beego.GlobalControllerRouter["github.com/souliot/pagoda/controllers:PropertyController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/souliot/pagoda/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/souliot/pagoda/controllers:UserController"],
        beego.ControllerComments{
            Method: "All",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/souliot/pagoda/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/souliot/pagoda/controllers:UserController"],
        beego.ControllerComments{
            Method: "Add",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/souliot/pagoda/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/souliot/pagoda/controllers:UserController"],
        beego.ControllerComments{
            Method: "DeleteMulti",
            Router: "/",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/souliot/pagoda/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/souliot/pagoda/controllers:UserController"],
        beego.ControllerComments{
            Method: "One",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/souliot/pagoda/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/souliot/pagoda/controllers:UserController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/souliot/pagoda/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/souliot/pagoda/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetUserInfo",
            Router: "/getUserInfo",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/souliot/pagoda/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/souliot/pagoda/controllers:UserController"],
        beego.ControllerComments{
            Method: "Login",
            Router: "/login",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/souliot/pagoda/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/souliot/pagoda/controllers:UserController"],
        beego.ControllerComments{
            Method: "Logout",
            Router: "/logout",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
