package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/udistrital/evaluacion_cumplido_prov_crud/models"
)

// Prueba unitaria funcionlidad Post /Insert de AsignacionEvaluadorItem
func TestPostAsignacionEvaluadorItem(t *testing.T) {
	asingancion_ev_item := models.AsignacionEvaluadorItem{
		AsignacionEvaluadorId: &models.AsignacionEvaluador{Id: 41},
		ItemId:                &models.Item{Id: 181},
		Activo:                true,
	}

	body, err := json.Marshal(asingancion_ev_item)
	if err != nil {
		t.Fatalf("Error al convertir a JSON: %v", err)
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", "http://localhost:8080/v1/asignacion_evaluador_item", bytes.NewReader(body))
	if err != nil {
		t.Errorf("Error al crear la solicitud: %v", err)
		t.Fail()
	}

	r, err := client.Do(req)
	if err != nil {
		t.Errorf("Error al ejecutar la solicitud: %v", err)
		t.Fail()
	}

	if r.StatusCode != http.StatusOK {
		t.Errorf("Se esperaba el código de estado 201, pero se obtuvo: %d", r.StatusCode)
		t.Fail()
	} else {
		t.Log("Test POST AsignacionEvaluadorItem Finalizado Correctamente (OK)")
	}
}
func TestPutAsignacionEvaluadorItem(t *testing.T) {
	item := models.AsignacionEvaluadorItem{
		Id:                    8435,
		AsignacionEvaluadorId: &models.AsignacionEvaluador{Id: 41},
		ItemId:                &models.Item{Id: 181},
		Activo:                true,
	}

	body, err := json.Marshal(item)
	if err != nil {
		t.Fatalf("Error al convertir a JSON: %v", err)
	}

	client := &http.Client{}
	req, err := http.NewRequest("PUT", "http://localhost:8080/v1/asignacion_evaluador_item/8435", bytes.NewReader(body))
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
	t.Log("Test PUT AsignacionEvaluadorItem finalizado correctamente.")
}

func TestGetAsignacionEvaluadorItem(t *testing.T) {
	id := 8435

	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("http://localhost:8080/v1/asignacion_evaluador_item/%d", id), nil)
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

	t.Log("Respuesta de la solicitud GET:", response)
	t.Log("Test GET AsignacionEvaluadorItem finalizado correctamente.")
}

func TestGetAllAsignacionEvaluadorItem(t *testing.T) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://localhost:8080/v1/asignacion_evaluador_item", nil)
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
	t.Log("Test GET All AsignacionEvaluadorItem finalizado correctamente.")
}

func TestDeleteAsignacionEvaluadorItem(t *testing.T) {
	client := &http.Client{}
	req, err := http.NewRequest("DELETE", "http://localhost:8080/v1/asignacion_evaluador_item/8435", nil)
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

	t.Log("Test DELETE AsignacionEvaluadorItem finalizado correctamente.")
}
