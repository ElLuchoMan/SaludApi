// @APIVersion 1.0.0
// @Title Salud API
// @Description Api para el módulo de Salud del proyecto SIBUD
// @Contact baluist@correo.udistrital.edu.co  - ddromeroq@correo.udistrital.edu.co
package routers

import (
	"github.com/ElLuchoMan/SaludApi/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/usuario",
			beego.NSInclude(
				&controllers.UsuarioController{},
			),
		),

		beego.NSNamespace("/prueba",
			beego.NSInclude(
				&controllers.PruebaController{},
			),
		),
		beego.NSNamespace("/accesoHistoria",
			beego.NSInclude(
				&controllers.AccesoHistoriaClinicaController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
