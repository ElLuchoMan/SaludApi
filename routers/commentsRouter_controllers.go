package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:AccesoHistoriaClinicaController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:AccesoHistoriaClinicaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:AccesoHistoriaClinicaController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:AccesoHistoriaClinicaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:AccesoHistoriaClinicaController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:AccesoHistoriaClinicaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:AccesoHistoriaClinicaController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:AccesoHistoriaClinicaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:AccesoHistoriaClinicaController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:AccesoHistoriaClinicaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:MedicinaAntecedenteController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:MedicinaAntecedenteController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:MedicinaAntecedenteController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:MedicinaAntecedenteController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:MedicinaAntecedenteController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:MedicinaAntecedenteController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:MedicinaAntecedenteController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:MedicinaAntecedenteController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:MedicinaAntecedenteController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:MedicinaAntecedenteController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:MedicinaConsultaFisioterapiaController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:MedicinaConsultaFisioterapiaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:MedicinaConsultaFisioterapiaController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:MedicinaConsultaFisioterapiaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:MedicinaConsultaFisioterapiaController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:MedicinaConsultaFisioterapiaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:MedicinaConsultaFisioterapiaController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:MedicinaConsultaFisioterapiaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:MedicinaConsultaFisioterapiaController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:MedicinaConsultaFisioterapiaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:MedicinaDiagnosticoController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:MedicinaDiagnosticoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:MedicinaDiagnosticoController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:MedicinaDiagnosticoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:MedicinaDiagnosticoController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:MedicinaDiagnosticoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:MedicinaDiagnosticoController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:MedicinaDiagnosticoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:MedicinaDiagnosticoController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:MedicinaDiagnosticoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:MedicinaExamenController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:MedicinaExamenController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:MedicinaExamenController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:MedicinaExamenController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:MedicinaExamenController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:MedicinaExamenController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:MedicinaExamenController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:MedicinaExamenController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:MedicinaExamenController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:MedicinaExamenController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:MedicinaHistoriaClinicaController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:MedicinaHistoriaClinicaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:MedicinaHistoriaClinicaController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:MedicinaHistoriaClinicaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:MedicinaHistoriaClinicaController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:MedicinaHistoriaClinicaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:MedicinaHistoriaClinicaController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:MedicinaHistoriaClinicaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:MedicinaHistoriaClinicaController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:MedicinaHistoriaClinicaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:MedicinaHojaHistoriaController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:MedicinaHojaHistoriaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:MedicinaHojaHistoriaController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:MedicinaHojaHistoriaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:MedicinaHojaHistoriaController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:MedicinaHojaHistoriaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:MedicinaHojaHistoriaController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:MedicinaHojaHistoriaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:MedicinaHojaHistoriaController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:MedicinaHojaHistoriaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:MedicinaSistemasController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:MedicinaSistemasController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:MedicinaSistemasController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:MedicinaSistemasController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:MedicinaSistemasController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:MedicinaSistemasController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:MedicinaSistemasController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:MedicinaSistemasController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:MedicinaSistemasController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:MedicinaSistemasController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:MedicinaTipoAntecedenteController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:MedicinaTipoAntecedenteController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:MedicinaTipoAntecedenteController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:MedicinaTipoAntecedenteController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:MedicinaTipoAntecedenteController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:MedicinaTipoAntecedenteController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:MedicinaTipoAntecedenteController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:MedicinaTipoAntecedenteController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:MedicinaTipoAntecedenteController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:MedicinaTipoAntecedenteController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
