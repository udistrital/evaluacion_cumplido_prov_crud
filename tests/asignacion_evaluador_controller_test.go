package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/udistrital/evaluacion_cumplido_prov_crud/models"
)

func TestPostAsignacionEvaluador(t *testing.T) {
	item := models.AsignacionEvaluador{
		EvaluacionId:             &models.Evaluacion{Id: 1},
		PersonaId:                "12345752",
		Cargo:                    "Cargo Test",
		PorcentajeEvaluacion:     85.5,
		RolAsignacionEvaluadorId: &models.RolAsignacionEvaluador{Id: 2},
		Activo:                   true,
	}

	body, err := json.Marshal(item)
	if err != nil {
		t.Fatalf("Error al convertir a JSON: %v", err)
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", "http://localhost:8080/v1/asignacion_evaluador", bytes.NewReader(body))
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

	if response["Data"] == nil {
		t.Fatal("No se recibió el campo 'Data' en la respuesta")
	}

	t.Log("Respuesta de la solicitud POST:", response)
	t.Log("Test POST AsignacionEvaluador finalizado correctamente.")
}

func TestPutAsignacionEvaluador(t *testing.T) {
	item := models.AsignacionEvaluador{
		Id:                       56,
		EvaluacionId:             &models.Evaluacion{Id: 1},
		PersonaId:                "12345",
		Cargo:                    "Cargo Actualizado",
		PorcentajeEvaluacion:     90.0,
		RolAsignacionEvaluadorId: &models.RolAsignacionEvaluador{Id: 2},
		Activo:                   true,
	}

	body, err := json.Marshal(item)
	if err != nil {
		t.Fatalf("Error al convertir a JSON: %v", err)
	}

	client := &http.Client{}
	req, err := http.NewRequest("PUT", "http://localhost:8080/v1/asignacion_evaluador/56", bytes.NewReader(body))
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

	if response["Data"] == nil {
		t.Fatal("No se recibió el campo 'Data' en la respuesta")
	}

	t.Log("Respuesta de la solicitud PUT:", response)
	t.Log("Test PUT AsignacionEvaluador finalizado correctamente.")
}

func TestGetAllAsignacionEvaluador(t *testing.T) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://localhost:8080/v1/asignacion_evaluador", nil)
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

	if response["Data"] == nil {
		t.Fatal("No se recibió el campo 'Data' en la respuesta")
	}

	t.Log("Respuesta de la solicitud GET All:", response)
	t.Log("Test GET All AsignacionEvaluador finalizado correctamente.")
}

func TestGetOneAsignacionEvaluador(t *testing.T) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://localhost:8080/v1/asignacion_evaluador/41", nil)
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

	if response["Data"] == nil {
		t.Fatal("No se recibió el campo 'Data' en la respuesta")
	}

	t.Log("Respuesta de la solicitud GET One:", response)
	t.Log("Test GET One AsignacionEvaluador finalizado correctamente.")
}
func TestDeleteAsignacionEvaluador(t *testing.T) {
	client := &http.Client{}
	req, err := http.NewRequest("DELETE", "http://localhost:8080/v1/asignacion_evaluador/55", nil)
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

	if response["Data"] == nil {
		t.Fatal("No se recibió el campo 'Data' en la respuesta")
	}

	t.Log("Respuesta de la solicitud DELETE:", response)
	t.Log("Test DELETE AsignacionEvaluador finalizado correctamente.")
}
