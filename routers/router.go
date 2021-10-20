// @APIVersion 1.0.0
// @Title Salud API
// @Description Api para el módulo de Salud del proyecto SIBUD
// @Contact baluist@correo.udistrital.edu.co
// @TermsOfServiceUrl http://www.udistrital.edu.co/politicas-de-privacidad.pdf
// @License BSD-3-Clause
// @LicenseUrl http://opensource.org/licenses/BSD-3-Clause
// @BasePath /SaludApi/v1

package routers

import (
	"github.com/ElLuchoMan/SaludApi/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/accesoHistoria",
			beego.NSInclude(
				&controllers.AccesoHistoriaClinicaController{},
			),
		),
		beego.NSNamespace("/medicinaAntecedente",
			beego.NSInclude(
				&controllers.MedicinaAntecedenteController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
