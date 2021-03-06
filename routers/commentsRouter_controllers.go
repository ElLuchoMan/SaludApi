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

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:MedicinaTipoExamenController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:MedicinaTipoExamenController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:MedicinaTipoExamenController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:MedicinaTipoExamenController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:MedicinaTipoExamenController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:MedicinaTipoExamenController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:MedicinaTipoExamenController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:MedicinaTipoExamenController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:MedicinaTipoExamenController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:MedicinaTipoExamenController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:OdontologiaAnamnesisController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:OdontologiaAnamnesisController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:OdontologiaAnamnesisController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:OdontologiaAnamnesisController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:OdontologiaAnamnesisController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:OdontologiaAnamnesisController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:OdontologiaAnamnesisController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:OdontologiaAnamnesisController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:OdontologiaAnamnesisController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:OdontologiaAnamnesisController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:OdontologiaControlPlacaController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:OdontologiaControlPlacaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:OdontologiaControlPlacaController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:OdontologiaControlPlacaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:OdontologiaControlPlacaController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:OdontologiaControlPlacaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:OdontologiaControlPlacaController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:OdontologiaControlPlacaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:OdontologiaControlPlacaController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:OdontologiaControlPlacaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:OdontologiaDiagnosticoController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:OdontologiaDiagnosticoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:OdontologiaDiagnosticoController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:OdontologiaDiagnosticoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:OdontologiaDiagnosticoController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:OdontologiaDiagnosticoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:OdontologiaDiagnosticoController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:OdontologiaDiagnosticoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:OdontologiaDiagnosticoController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:OdontologiaDiagnosticoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:OdontologiaExamenDentalController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:OdontologiaExamenDentalController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:OdontologiaExamenDentalController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:OdontologiaExamenDentalController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:OdontologiaExamenDentalController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:OdontologiaExamenDentalController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:OdontologiaExamenDentalController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:OdontologiaExamenDentalController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:OdontologiaExamenDentalController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:OdontologiaExamenDentalController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:OdontologiaExamenEstomatologicoController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:OdontologiaExamenEstomatologicoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:OdontologiaExamenEstomatologicoController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:OdontologiaExamenEstomatologicoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:OdontologiaExamenEstomatologicoController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:OdontologiaExamenEstomatologicoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:OdontologiaExamenEstomatologicoController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:OdontologiaExamenEstomatologicoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:OdontologiaExamenEstomatologicoController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:OdontologiaExamenEstomatologicoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:OdontologiaExamenesComplementariosController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:OdontologiaExamenesComplementariosController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:OdontologiaExamenesComplementariosController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:OdontologiaExamenesComplementariosController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:OdontologiaExamenesComplementariosController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:OdontologiaExamenesComplementariosController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:OdontologiaExamenesComplementariosController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:OdontologiaExamenesComplementariosController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:OdontologiaExamenesComplementariosController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:OdontologiaExamenesComplementariosController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:OdontologiaOdontogramaController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:OdontologiaOdontogramaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:OdontologiaOdontogramaController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:OdontologiaOdontogramaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:OdontologiaOdontogramaController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:OdontologiaOdontogramaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:OdontologiaOdontogramaController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:OdontologiaOdontogramaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:OdontologiaOdontogramaController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:OdontologiaOdontogramaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:OdontologiaTipoOdontogramaController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:OdontologiaTipoOdontogramaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:OdontologiaTipoOdontogramaController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:OdontologiaTipoOdontogramaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:OdontologiaTipoOdontogramaController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:OdontologiaTipoOdontogramaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:OdontologiaTipoOdontogramaController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:OdontologiaTipoOdontogramaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:OdontologiaTipoOdontogramaController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:OdontologiaTipoOdontogramaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:PsicologiaAntecedenteController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:PsicologiaAntecedenteController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:PsicologiaAntecedenteController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:PsicologiaAntecedenteController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:PsicologiaAntecedenteController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:PsicologiaAntecedenteController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:PsicologiaAntecedenteController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:PsicologiaAntecedenteController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:PsicologiaAntecedenteController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:PsicologiaAntecedenteController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:PsicologiaComportamientoConsultaController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:PsicologiaComportamientoConsultaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:PsicologiaComportamientoConsultaController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:PsicologiaComportamientoConsultaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:PsicologiaComportamientoConsultaController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:PsicologiaComportamientoConsultaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:PsicologiaComportamientoConsultaController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:PsicologiaComportamientoConsultaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:PsicologiaComportamientoConsultaController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:PsicologiaComportamientoConsultaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:PsicologiaComposicionFamiliarController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:PsicologiaComposicionFamiliarController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:PsicologiaComposicionFamiliarController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:PsicologiaComposicionFamiliarController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:PsicologiaComposicionFamiliarController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:PsicologiaComposicionFamiliarController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:PsicologiaComposicionFamiliarController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:PsicologiaComposicionFamiliarController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:PsicologiaComposicionFamiliarController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:PsicologiaComposicionFamiliarController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:PsicologiaDiagnosticoController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:PsicologiaDiagnosticoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:PsicologiaDiagnosticoController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:PsicologiaDiagnosticoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:PsicologiaDiagnosticoController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:PsicologiaDiagnosticoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:PsicologiaDiagnosticoController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:PsicologiaDiagnosticoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:PsicologiaDiagnosticoController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:PsicologiaDiagnosticoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:PsicologiaLimitesController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:PsicologiaLimitesController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:PsicologiaLimitesController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:PsicologiaLimitesController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:PsicologiaLimitesController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:PsicologiaLimitesController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:PsicologiaLimitesController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:PsicologiaLimitesController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:PsicologiaLimitesController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:PsicologiaLimitesController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:PsicologiaTipoAntecedenteController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:PsicologiaTipoAntecedenteController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:PsicologiaTipoAntecedenteController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:PsicologiaTipoAntecedenteController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:PsicologiaTipoAntecedenteController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:PsicologiaTipoAntecedenteController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:PsicologiaTipoAntecedenteController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:PsicologiaTipoAntecedenteController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:PsicologiaTipoAntecedenteController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:PsicologiaTipoAntecedenteController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:PsicologiaValoracionInterpersonalController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:PsicologiaValoracionInterpersonalController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:PsicologiaValoracionInterpersonalController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:PsicologiaValoracionInterpersonalController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:PsicologiaValoracionInterpersonalController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:PsicologiaValoracionInterpersonalController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:PsicologiaValoracionInterpersonalController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:PsicologiaValoracionInterpersonalController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:PsicologiaValoracionInterpersonalController"] = append(beego.GlobalControllerRouter["github.com/ElLuchoMan/SaludApi/controllers:PsicologiaValoracionInterpersonalController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
