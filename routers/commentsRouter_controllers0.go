package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["cdve_crud_api/controllers:ContratadoController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:ContratadoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/:tipoDedicacion`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:PrecontratadoController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:PrecontratadoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/:idResolucion`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:PrecontratadoController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:PrecontratadoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:idResolucion/:idPersona`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:ResolucionVinculacionController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:ResolucionVinculacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:PersonaEscalafonController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:PersonaEscalafonController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:ResolucionCompletaController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:ResolucionCompletaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:idResolucion`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:ResolucionCompletaController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:ResolucionCompletaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:idResolucion`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:OrdenadorGastoController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:OrdenadorGastoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:OrdenadorGastoController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:OrdenadorGastoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:OrdenadorGastoController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:OrdenadorGastoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:OrdenadorGastoController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:OrdenadorGastoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:OrdenadorGastoController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:OrdenadorGastoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:DependenciaPadreController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:DependenciaPadreController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:DependenciaPadreController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:DependenciaPadreController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:DependenciaPadreController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:DependenciaPadreController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:DependenciaPadreController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:DependenciaPadreController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:DependenciaPadreController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:DependenciaPadreController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:DependenciaTipoOrdenadorGastoController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:DependenciaTipoOrdenadorGastoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:DependenciaTipoOrdenadorGastoController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:DependenciaTipoOrdenadorGastoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:DependenciaTipoOrdenadorGastoController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:DependenciaTipoOrdenadorGastoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:DependenciaTipoOrdenadorGastoController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:DependenciaTipoOrdenadorGastoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:DependenciaTipoOrdenadorGastoController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:DependenciaTipoOrdenadorGastoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:TipoOrdenadorGastoController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:TipoOrdenadorGastoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:TipoOrdenadorGastoController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:TipoOrdenadorGastoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:TipoOrdenadorGastoController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:TipoOrdenadorGastoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:TipoOrdenadorGastoController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:TipoOrdenadorGastoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:TipoOrdenadorGastoController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:TipoOrdenadorGastoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:FacultadController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:FacultadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:FacultadController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:FacultadController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:idFacultad`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:ProyectoCurricularController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:ProyectoCurricularController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/:nivelAcademico/:idFacultad`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:ProyectoCurricularController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:ProyectoCurricularController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:idProyectoCurricular`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:ContratoGeneralController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:ContratoGeneralController"],
		beego.ControllerComments{
			Method: "InsertarContratos",
			Router: `InsertarContratos`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:VinculacionDocenteController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:VinculacionDocenteController"],
		beego.ControllerComments{
			Method: "InsertarVinculaciones",
			Router: `InsertarVinculaciones`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:OrdenadorGastoController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:OrdenadorGastoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:OrdenadorGastoController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:OrdenadorGastoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:OrdenadorGastoController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:OrdenadorGastoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:OrdenadorGastoController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:OrdenadorGastoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:OrdenadorGastoController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:OrdenadorGastoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:ResolucionController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:ResolucionController"],
		beego.ControllerComments{
			Method: "CancelarResolucion",
			Router: `CancelarResolucion/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

}