package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:PruebaController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:PruebaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:PruebaController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:PruebaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:PruebaController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:PruebaController"],
        beego.ControllerComments{
            Method: "GetA",
            Router: "/GetA",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:UsuarioController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:UsuarioController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:UsuarioController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:UsuarioController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:UsuarioController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:UsuarioController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:UsuarioController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:UsuarioController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:UsuarioController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:UsuarioController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
