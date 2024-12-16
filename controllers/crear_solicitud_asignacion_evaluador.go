package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/evaluacion_cumplido_prov_crud/models"
)

// TrCrearSolicitudAsignacionEvaluadorController operations for TrCrearSolicitudAsignacionEvaluador
type TrCrearSolicitudAsignacionEvaluadorController struct {
	beego.Controller
}

// URLMapping ...
func (c *TrCrearSolicitudAsignacionEvaluadorController) URLMapping() {
	c.Mapping("Post", c.Post)
}

// Post ...
// @Title Create
// @Description create SolicitudAsignacionEvaluador
// @Param	body		body 	models.SolicitudAsignacionEvaluador	"Body for SolicitudAsignacionEvaluador"
// @Success 201 {object} models.SolicitudAsignacionEvaluador
// @Failure 400 The request contains incorrect data
// @router / [post]
func (c *TrCrearSolicitudAsignacionEvaluadorController) Post() {
	var v models.SolicitudAsignacionEvaluador

	// Parsear el cuerpo de la solicitud
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err != nil {
		logs.Error("Error al parsear el cuerpo de la solicitud: ", err)
		c.Data["json"] = map[string]interface{}{"Success": false, "Status": "400", "Message": "Invalid request body", "Error": err.Error()}
		c.Abort("400")
	}

	// Llamar al modelo para crear la solicitud
	data, err := models.CrearSolicitudAsignacionEvaluador(&v)
	if err != nil {
		logs.Error("Error al crear la solicitud de asignación: ", err)
		c.Data["json"] = map[string]interface{}{"Success": false, "Status": "500", "Message": "Error creating request", "Error": err.Error()}
		c.Abort("500")
	}

	// Respuesta exitosa
	c.Ctx.Output.SetStatus(201)
	c.Data["json"] = map[string]interface{}{"Success": true, "Status": "201", "Message": "Request created successfully", "Data": data}
	c.ServeJSON()
}
