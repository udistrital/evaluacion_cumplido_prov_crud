package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/udistrital/evaluacion_cumplido_prov_crud/models"
)

func TestPostTrCrearSolicitudAsignacionEvaluador(t *testing.T) {
	item := models.SolicitudAsignacionEvaluador{
		EvaluacionId:           43,
		PersonaId:              "123455",
		Cargo:                  "Evaluador",
		PorcentajeEvaluacion:   80.5,
		RolAsignacionEvaluador: "PE",
	}

	body, err := json.Marshal(item)
	if err != nil {
		t.Fatalf("Error al convertir a JSON: %v", err)
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", "http://localhost:8080/v1/crear_solicitud_asignacion_evaluador", bytes.NewReader(body))
	if err != nil {
		t.Fatalf("Error al crear la solicitud POST: %v", err)
	}

	r, err := client.Do(req)
	if err != nil {
		t.Fatalf("Error al ejecutar la solicitud POST: %v", err)
	}
	defer r.Body.Close()

	if r.StatusCode != http.StatusCreated {
		t.Errorf("Se esperaba el código de estado 201, pero se obtuvo: %d", r.StatusCode)
	}

	var response map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&response); err != nil {
		t.Fatalf("Error al decodificar la respuesta: %v", err)
	}

	t.Log("Respuesta de la solicitud POST:", response)
	t.Log("Test POST TrCrearSolicitudAsignacionEvaluador finalizado correctamente.")
}