// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/udistrital/evaluacion_cumplido_prov_crud/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/estado_cumplido",
			beego.NSInclude(
				&controllers.EstadoCumplidoController{},
			),
		),

		beego.NSNamespace("/soporte_cumplido",
			beego.NSInclude(
				&controllers.SoporteCumplidoController{},
			),
		),

		beego.NSNamespace("/informe_satisfaccion",
			beego.NSInclude(
				&controllers.InformeSatisfaccionController{},
			),
		),

		beego.NSNamespace("/cumplido_proveedor",
			beego.NSInclude(
				&controllers.CumplidoProveedorController{},
			),
		),

		beego.NSNamespace("/cambio_estado_cumplido",
			beego.NSInclude(
				&controllers.CambioEstadoCumplidoController{},
			),
		),

		beego.NSNamespace("/comentario_soporte",
			beego.NSInclude(
				&controllers.ComentarioSoporteController{},
			),
		),

		beego.NSNamespace("/tipo_pago",
			beego.NSInclude(
				&controllers.TipoPagoController{},
			),
		),

		beego.NSNamespace("/tipo_cuenta_bancaria",
			beego.NSInclude(
				&controllers.TipoCuentaBancariaController{},
			),
		),

		beego.NSNamespace("/tipo_factura",
			beego.NSInclude(
				&controllers.TipoFacturaController{},
			),
		),

		beego.NSNamespace("/clasificacion",
			beego.NSInclude(
				&controllers.ClasificacionController{},
			),
		),

		beego.NSNamespace("/resultado_evaluacion",
			beego.NSInclude(
				&controllers.ResultadoEvaluacionController{},
			),
		),

		beego.NSNamespace("/asignacion_evaluador",
			beego.NSInclude(
				&controllers.AsignacionEvaluadorController{},
			),
		),

		beego.NSNamespace("/evaluacion",
			beego.NSInclude(
				&controllers.EvaluacionController{},
			),
		),

		beego.NSNamespace("/cambio_estado_evaluacion",
			beego.NSInclude(
				&controllers.CambioEstadoEvaluacionController{},
			),
		),

		beego.NSNamespace("/estado_evaluacion",
			beego.NSInclude(
				&controllers.EstadoEvaluacionController{},
			),
		),

		beego.NSNamespace("/item",
			beego.NSInclude(
				&controllers.ItemController{},
			),
		),

		beego.NSNamespace("/asignacion_evaluador_item",
			beego.NSInclude(
				&controllers.AsignacionEvaluadorItemController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
