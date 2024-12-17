package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/evaluacion_cumplido_prov_crud/models"
)

// TrCrearSolicitudEvaluacionController operations for TrCrearSolicitudEvaluacion
type TrCrearSolicitudEvaluacionController struct {
	beego.Controller
}

// URLMapping ...
func (c *TrCrearSolicitudEvaluacionController) URLMapping() {
	c.Mapping("Post", c.Post)
}

// Post ...
// @Title Create
// @Description create SolicitudEvaluacion
// @Param	body		body 	models.SolicitudEvaluacion	"Body for SolicitudEvaluacion"
// @Success 201 {object} models.SolicitudEvaluacion
// @Failure 400 The request contains incorrect data
// @router / [post]
func (c *TrCrearSolicitudEvaluacionController) Post() {
	var v models.SolicitudEvaluacion

	// Parsear el cuerpo de la solicitud
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err != nil {
		logs.Error("Error al parsear el cuerpo de la solicitud: ", err)
		c.Data["json"] = map[string]interface{}{"Success": false, "Status": "400", "Message": "Invalid request body", "Error": err.Error()}
		c.Abort("400")
		return
	}

	// Llamar al modelo para crear la solicitud de evaluación
	data, err := models.CrearSolicitudEvaluacion(&v)
	if err != nil {
		logs.Error("Error al crear la solicitud de evaluación: ", err)
		c.Data["json"] = map[string]interface{}{"Success": false, "Status": "500", "Message": "Error creating evaluation request", "Error": err.Error()}
		c.Abort("500")
		return
	}

	// Respuesta exitosa
	c.Ctx.Output.SetStatus(201)
	c.Data["json"] = map[string]interface{}{"Success": true, "Status": "201", "Message": "Evaluation request created successfully", "Data": data}
	c.ServeJSON()
}
