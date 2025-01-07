package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/udistrital/evaluacion_cumplido_prov_crud/models"
)

func TestPostCambioEstadoAsignacionEvaluador(t *testing.T) {
	client := &http.Client{}
	cambioEstado := models.CambioEstadoAsignacionEvaluador{
		EstadoAsignacionEvaluadorId: &models.EstadoAsignacionEvaluador{Id: 1},
		AsignacionEvaluadorId:       &models.AsignacionEvaluador{Id: 1},
		Activo:                      true,
	}

	bodyData, err := json.Marshal(cambioEstado)
	if err != nil {
		t.Fatalf("Error al convertir el cuerpo a JSON: %v", err)
	}

	req, err := http.NewRequest("POST", "http://localhost:8080/v1/cambio_estado_asignacion_evaluador", bytes.NewBuffer(bodyData))
	if err != nil {
		t.Fatalf("Error al crear la solicitud POST: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
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

	if response["Success"] != true {
		t.Fatal("La respuesta no contiene 'Success': true")
	}

	t.Log("Respuesta de la solicitud POST:", response)
	t.Log("Test POST CambioEstadoAsignacionEvaluador finalizado correctamente.")
}

func TestPutCambioEstadoAsignacionEvaluador(t *testing.T) {
	cambioEstado := models.CambioEstadoAsignacionEvaluador{
		Id:                          121,
		EstadoAsignacionEvaluadorId: &models.EstadoAsignacionEvaluador{Id: 1},
		AsignacionEvaluadorId:       &models.AsignacionEvaluador{Id: 1},
		Activo:                      true,
	}

	body, err := json.Marshal(cambioEstado)
	if err != nil {
		t.Fatalf("Error al convertir a JSON: %v", err)
	}

	client := &http.Client{}
	req, err := http.NewRequest("PUT", "http://localhost:8080/v1/cambio_estado_asignacion_evaluador/122", bytes.NewReader(body))
	if err != nil {
		t.Fatalf("Error al crear la solicitud PUT: %v", err)
	}

	r, err := client.Do(req)
	if err != nil {
		t.Fatalf("Error al ejecutar la solicitud PUT: %v", err)
	}
	defer r.Body.Close()

	if r.StatusCode != http.StatusOK {
		t.Errorf("Se esperaba el código de estado 200, pero se obtuvo: %d", r.StatusCode)
	}

	var response map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&response); err != nil {
		t.Fatalf("Error al decodificar la respuesta: %v", err)
	}

	t.Log("Respuesta de la solicitud PUT:", response)
	t.Log("Test PUT CambioEstadoAsignacionEvaluador finalizado correctamente.")
}

func TestDeleteCambioEstadoAsignacionEvaluador(t *testing.T) {
	client := &http.Client{}
	req, err := http.NewRequest("DELETE", "http://localhost:8080/v1/cambio_estado_asignacion_evaluador/123", nil)
	if err != nil {
		t.Fatalf("Error al crear la solicitud DELETE: %v", err)
	}

	r, err := client.Do(req)
	if err != nil {
		t.Fatalf("Error al ejecutar la solicitud DELETE: %v", err)
	}
	defer r.Body.Close()

	if r.StatusCode != http.StatusOK {
		t.Errorf("Se esperaba el código de estado 200, pero se obtuvo: %d", r.StatusCode)
	}

	var response map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&response); err != nil {
		t.Fatalf("Error al decodificar la respuesta: %v", err)
	}

	t.Log("Respuesta de la solicitud DELETE:", response)
	t.Log("Test DELETE CambioEstadoAsignacionEvaluador finalizado correctamente.")
}

func TestGetOneCambioEstadoAsignacionEvaluador(t *testing.T) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://localhost:8080/v1/cambio_estado_asignacion_evaluador/121", nil)
	if err != nil {
		t.Fatalf("Error al crear la solicitud GET: %v", err)
	}

	r, err := client.Do(req)
	if err != nil {
		t.Fatalf("Error al ejecutar la solicitud GET: %v", err)
	}
	defer r.Body.Close()

	if r.StatusCode != http.StatusOK {
		t.Errorf("Se esperaba el código de estado 200, pero se obtuvo: %d", r.StatusCode)
	}

	var response map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&response); err != nil {
		t.Fatalf("Error al decodificar la respuesta: %v", err)
	}

	t.Log("Respuesta de la solicitud GET:", response)
	t.Log("Test GET One CambioEstadoAsignacionEvaluador finalizado correctamente.")
}

func TestGetAllCambioEstadoAsignacionEvaluador(t *testing.T) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://localhost:8080/v1/cambio_estado_asignacion_evaluador", nil)
	if err != nil {
		t.Fatalf("Error al crear la solicitud GET: %v", err)
	}

	r, err := client.Do(req)
	if err != nil {
		t.Fatalf("Error al ejecutar la solicitud GET: %v", err)
	}
	defer r.Body.Close()

	if r.StatusCode != http.StatusOK {
		t.Errorf("Se esperaba el código de estado 200, pero se obtuvo: %d", r.StatusCode)
	}

	var response map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&response); err != nil {
		t.Fatalf("Error al decodificar la respuesta: %v", err)
	}

	t.Log("Respuesta de la solicitud GET:", response)
	t.Log("Test GET All CambioEstadoAsignacionEvaluador finalizado correctamente.")
}
