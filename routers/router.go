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
		beego.NSNamespace("/medicinaConsultaFisioterapia",
			beego.NSInclude(
				&controllers.MedicinaConsultaFisioterapiaController{},
			),
		),
		beego.NSNamespace("/medicinaDiagnostico",
			beego.NSInclude(
				&controllers.MedicinaDiagnosticoController{},
			),
		),
		beego.NSNamespace("/medicinaExamen",
			beego.NSInclude(
				&controllers.MedicinaExamenController{},
			),
		), beego.NSNamespace("/medicinaHistoriaClinica",
			beego.NSInclude(
				&controllers.MedicinaHistoriaClinicaController{},
			),
		), beego.NSNamespace("/medicinaHojaHistoria",
			beego.NSInclude(
				&controllers.MedicinaHojaHistoriaController{},
			),
		), beego.NSNamespace("/medicinaSistemas",
			beego.NSInclude(
				&controllers.MedicinaSistemasController{},
			),
		),
		beego.NSNamespace("/medicinaTipoAntecedente",
			beego.NSInclude(
				&controllers.MedicinaTipoAntecedenteController{},
			),
		),
		beego.NSNamespace("/medicinaTipoExamen",
			beego.NSInclude(
				&controllers.MedicinaTipoExamenController{},
			),
		),
		beego.NSNamespace("/odontologiaAnamnesis",
			beego.NSInclude(
				&controllers.OdontologiaAnamnesisController{},
			),
		),
		beego.NSNamespace("/odontologiaControlPlaca",
			beego.NSInclude(
				&controllers.OdontologiaControlPlacaController{},
			),
		),
		beego.NSNamespace("/odontologiaDiagnostico",
			beego.NSInclude(
				&controllers.OdontologiaDiagnosticoController{},
			),
		),
		beego.NSNamespace("/odontologiaExamenDental",
			beego.NSInclude(
				&controllers.OdontologiaExamenDentalController{},
			),
		),
		beego.NSNamespace("/odontologiaExamenesComplementarios",
			beego.NSInclude(
				&controllers.OdontologiaExamenesComplementariosController{},
			),
		),
		beego.NSNamespace("/odontologiaExamenEstomatologico",
			beego.NSInclude(
				&controllers.OdontologiaExamenEstomatologicoController{},
			),
		),
		beego.NSNamespace("/odontologiaOdontograma",
			beego.NSInclude(
				&controllers.OdontologiaOdontogramaController{},
			),
		),
		beego.NSNamespace("/odontologiaTipoOdontograma",
			beego.NSInclude(
				&controllers.OdontologiaTipoOdontogramaController{},
			),
		),
		beego.NSNamespace("/psicologiaAntecedente",
			beego.NSInclude(
				&controllers.PsicologiaAntecedenteController{},
			),
		),
		beego.NSNamespace("/psicologiaComportamientoConsulta",
			beego.NSInclude(
				&controllers.PsicologiaComportamientoConsultaController{},
			),
		),
		beego.NSNamespace("/psicologiaComposicionFamiliar",
			beego.NSInclude(
				&controllers.PsicologiaComposicionFamiliarController{},
			),
		),
		beego.NSNamespace("/psicologiaDiagnostico",
			beego.NSInclude(
				&controllers.PsicologiaDiagnosticoController{},
			),
		),
		beego.NSNamespace("/psicologiaLimites",
			beego.NSInclude(
				&controllers.PsicologiaLimitesController{},
			),
		),
		beego.NSNamespace("/psicologiaTipoAntecedente",
			beego.NSInclude(
				&controllers.PsicologiaTipoAntecedenteController{},
			),
		),
		beego.NSNamespace("/psicologiaTipoAntecedente",
			beego.NSInclude(
				&controllers.PsicologiaTipoAntecedenteController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
