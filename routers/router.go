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
