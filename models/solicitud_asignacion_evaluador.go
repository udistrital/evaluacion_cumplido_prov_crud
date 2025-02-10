package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

type SolicitudAsignacionEvaluador struct {
	EvaluacionId           int
	PersonaId              string
	Cargo                  string
	PorcentajeEvaluacion   float64
	RolAsignacionEvaluador string
}

// transaction create a asignacion_evaluador con cambio estado asignacion_evaluador inicial de evaluacion_asignada

func CrearSolicitudAsignacionEvaluador(m *SolicitudAsignacionEvaluador) (asignacion_evaluador *AsignacionEvaluador, err error) {

	o := orm.NewOrm()

	err = o.Begin()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var asignacion AsignacionEvaluador
	var rol_asignacion_evaluador RolAsignacionEvaluador
	var estado_asignacion_evaluador EstadoAsignacionEvaluador
	var evaluacion Evaluacion
	var cambio_estado_asignacion_evaluador CambioEstadoAsignacionEvaluador

	//Se busca el RolAsignacionEvaluador para validar que exista y recuperamos su id
	response, err := GetAllRolAsignacionEvaluador(map[string]string{"codigo_abreviacion": m.RolAsignacionEvaluador}, []string{}, []string{}, []string{}, 0, 1)
	if err != nil {
		o.Rollback()
		return nil, err
	}

	rol_asignacion_evaluador.Id = response[0].(RolAsignacionEvaluador).Id
	evaluacion.Id = m.EvaluacionId

	// Se crea la asignacion_evaluador

	asignacion.EvaluacionId = &evaluacion
	asignacion.PersonaId = m.PersonaId
	asignacion.Cargo = m.Cargo
	asignacion.PorcentajeEvaluacion = m.PorcentajeEvaluacion
	asignacion.RolAsignacionEvaluadorId = &rol_asignacion_evaluador
	asignacion.Activo = true

	created, id_asignacion_evaluador, err := CreateOrUpdateAsignacionEvaluador(&asignacion)
	if err != nil {
		o.Rollback()
		return nil, err
	}

	// Se actualiza el id de la asignacion_evaluador creada o actualizada para que se pueda usar más adelante
	asignacion.Id = id_asignacion_evaluador

	if created {
		//Se busca el id del primer estado de la asignacion_evaluacion que es Evaluación Asignada
		res_estado, err := GetAllEstadoAsignacionEvaluador(map[string]string{"codigo_abreviacion": "EAG"}, []string{}, []string{}, []string{}, 0, 1)
		if err != nil {
			o.Rollback()
			return nil, err
		}

		estado_asignacion_evaluador.Id = res_estado[0].(EstadoAsignacionEvaluador).Id

		//Se crea el primer cambio estado de la asignacion_evaluador
		cambio_estado_asignacion_evaluador.EstadoAsignacionEvaluadorId = &estado_asignacion_evaluador
		cambio_estado_asignacion_evaluador.AsignacionEvaluadorId = &asignacion
		cambio_estado_asignacion_evaluador.Activo = true

		_, err = o.Insert(&cambio_estado_asignacion_evaluador)
		if err != nil {
			o.Rollback()
			return nil, err
		}
	}

	err = o.Commit()

	asignacion_evaluador = &asignacion

	return asignacion_evaluador, nil
}
