package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:AsignacionEvaluadorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:AsignacionEvaluadorController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:AsignacionEvaluadorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:AsignacionEvaluadorController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:AsignacionEvaluadorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:AsignacionEvaluadorController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:AsignacionEvaluadorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:AsignacionEvaluadorController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:AsignacionEvaluadorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:AsignacionEvaluadorController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:AsignacionEvaluadorItemController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:AsignacionEvaluadorItemController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:AsignacionEvaluadorItemController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:AsignacionEvaluadorItemController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:AsignacionEvaluadorItemController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:AsignacionEvaluadorItemController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:AsignacionEvaluadorItemController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:AsignacionEvaluadorItemController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:AsignacionEvaluadorItemController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:AsignacionEvaluadorItemController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:CambioEstadoCumplidoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:CambioEstadoCumplidoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:CambioEstadoCumplidoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:CambioEstadoCumplidoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:CambioEstadoCumplidoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:CambioEstadoCumplidoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:CambioEstadoCumplidoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:CambioEstadoCumplidoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:CambioEstadoCumplidoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:CambioEstadoCumplidoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:CambioEstadoEvaluacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:CambioEstadoEvaluacionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:CambioEstadoEvaluacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:CambioEstadoEvaluacionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:CambioEstadoEvaluacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:CambioEstadoEvaluacionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:CambioEstadoEvaluacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:CambioEstadoEvaluacionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:CambioEstadoEvaluacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:CambioEstadoEvaluacionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:ClasificacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:ClasificacionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:ClasificacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:ClasificacionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:ClasificacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:ClasificacionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:ClasificacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:ClasificacionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:ClasificacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:ClasificacionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:ComentarioSoporteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:ComentarioSoporteController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:ComentarioSoporteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:ComentarioSoporteController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:ComentarioSoporteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:ComentarioSoporteController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:ComentarioSoporteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:ComentarioSoporteController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:ComentarioSoporteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:ComentarioSoporteController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:CumplidoProveedorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:CumplidoProveedorController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:CumplidoProveedorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:CumplidoProveedorController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:CumplidoProveedorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:CumplidoProveedorController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:CumplidoProveedorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:CumplidoProveedorController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:CumplidoProveedorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:CumplidoProveedorController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:EstadoCumplidoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:EstadoCumplidoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:EstadoCumplidoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:EstadoCumplidoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:EstadoCumplidoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:EstadoCumplidoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:EstadoCumplidoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:EstadoCumplidoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:EstadoCumplidoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:EstadoCumplidoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:EstadoEvaluacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:EstadoEvaluacionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:EstadoEvaluacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:EstadoEvaluacionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:EstadoEvaluacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:EstadoEvaluacionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:EstadoEvaluacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:EstadoEvaluacionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:EstadoEvaluacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:EstadoEvaluacionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:EvaluacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:EvaluacionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:EvaluacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:EvaluacionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:EvaluacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:EvaluacionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:EvaluacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:EvaluacionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:EvaluacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:EvaluacionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:InformeSatisfaccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:InformeSatisfaccionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:InformeSatisfaccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:InformeSatisfaccionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:InformeSatisfaccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:InformeSatisfaccionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:InformeSatisfaccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:InformeSatisfaccionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:InformeSatisfaccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:InformeSatisfaccionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:ItemController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:ItemController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:ItemController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:ItemController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:ItemController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:ItemController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:ItemController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:ItemController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:ItemController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:ItemController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:ResultadoEvaluacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:ResultadoEvaluacionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:ResultadoEvaluacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:ResultadoEvaluacionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:ResultadoEvaluacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:ResultadoEvaluacionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:ResultadoEvaluacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:ResultadoEvaluacionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:ResultadoEvaluacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:ResultadoEvaluacionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:SoporteCumplidoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:SoporteCumplidoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:SoporteCumplidoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:SoporteCumplidoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:SoporteCumplidoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:SoporteCumplidoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:SoporteCumplidoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:SoporteCumplidoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:SoporteCumplidoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:SoporteCumplidoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:TipoCuentaBancariaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:TipoCuentaBancariaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:TipoCuentaBancariaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:TipoCuentaBancariaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:TipoCuentaBancariaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:TipoCuentaBancariaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:TipoCuentaBancariaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:TipoCuentaBancariaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:TipoCuentaBancariaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:TipoCuentaBancariaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:TipoFacturaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:TipoFacturaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:TipoFacturaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:TipoFacturaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:TipoFacturaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:TipoFacturaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:TipoFacturaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:TipoFacturaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:TipoFacturaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:TipoFacturaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:TipoPagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:TipoPagoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:TipoPagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:TipoPagoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:TipoPagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:TipoPagoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:TipoPagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:TipoPagoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:TipoPagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_cumplido_prov_crud/controllers:TipoPagoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
