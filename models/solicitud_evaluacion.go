package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

type SolicitudEvaluacion struct {
	ContratoSuscritoId int
	VigenciaContrato   int
}

// transaction create a evaluacion con cambio estado evaluacion inicial de Gestión
func CrearSolicitudEvaluacion(m *SolicitudEvaluacion) (evaluacion *Evaluacion, err error) {

	o := orm.NewOrm()

	err = o.Begin()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var evaluacion_crear Evaluacion
	var cambio_estado_evaluacion CambioEstadoEvaluacion
	var estado_evaluacion EstadoEvaluacion

	// Desactivar evaluaciones existentes con el mismo contrato y vigencia
	// _, err = o.QueryTable("evaluacion").
	// 	Filter("contrato_suscrito_id", m.ContratoSuscritoId).
	// 	Filter("vigencia_contrato", m.VigenciaContrato).
	// 	Filter("activo", true).
	// 	Update(orm.Params{
	// 		"activo": false,
	// 	})

	// if err != nil {
	// 	o.Rollback()
	// 	return nil, err
	// }

	// Se crea la evaluacion
	evaluacion_crear.ContratoSuscritoId = m.ContratoSuscritoId
	evaluacion_crear.VigenciaContrato = m.VigenciaContrato
	evaluacion_crear.Activo = true
	id_evaluacion, err := o.Insert(&evaluacion_crear)
	if err != nil {
		o.Rollback()
		return nil, err
	}

	// Se actualiza el id de la evaluacion creada para que se pueda usar más adelante
	evaluacion_crear.Id = int(id_evaluacion)

	// Se busca el id del primer estado de la evaluacion que es Gestión
	res_estado, err := GetAllEstadoEvaluacion(map[string]string{"codigo_abreviacion": "GNT"}, []string{}, []string{}, []string{}, 0, 1)
	if err != nil {
		o.Rollback()
		return nil, err
	}

	// Mapear el estado de la evaluacion
	estado_evaluacion.Id = res_estado[0].(EstadoEvaluacion).Id

	// Se crea el primer cambio estado de la evaluacion
	cambio_estado_evaluacion.EstadoEvaluacionId = &estado_evaluacion
	cambio_estado_evaluacion.EvaluacionId = &evaluacion_crear
	cambio_estado_evaluacion.Activo = true

	_, err = o.Insert(&cambio_estado_evaluacion)
	if err != nil {
		o.Rollback()
		return nil, err
	}

	o.Commit()
	evaluacion = &evaluacion_crear
	return evaluacion, nil
}
