package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/udistrital/evaluacion_cumplido_prov_crud/models"
)

func TestPostEvaluacion(t *testing.T) {
	evaluacion := models.Evaluacion{
		ContratoSuscritoId: 123,
		VigenciaContrato:   2025,
	}

	body, err := json.Marshal(evaluacion)
	if err != nil {
		t.Fatalf("Error al convertir el cuerpo de la solicitud a JSON: %v", err)
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", "http://localhost:8080/v1/evaluacion", bytes.NewBuffer(body))
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
		t.Errorf("Se esperaba 'Success' = true, pero se obtuvo: %v", response["Success"])
	}

	if response["Data"] == nil {
		t.Errorf("Se esperaba que la respuesta contuviera los datos de la evaluación creada")
	}

	t.Log("Respuesta de la solicitud POST:", response)
	t.Log("Test POST Evaluacion finalizado correctamente.")
}

func TestPutEvaluacion(t *testing.T) {
	evaluacion := models.Evaluacion{
		Id:                 16,
		ContratoSuscritoId: 123,
		VigenciaContrato:   2025,
		Activo:             false,
	}
	body, err := json.Marshal(evaluacion)
	if err != nil {
		t.Fatalf("Error al convertir el cuerpo de la solicitud a JSON: %v", err)
	}
	client := &http.Client{}
	req, err := http.NewRequest("PUT", "http://localhost:8080/v1/evaluacion/16", bytes.NewBuffer(body))
	if err != nil {
		t.Fatalf("Error al crear la solicitud PUT: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
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
	if response["Success"] != true {
		t.Errorf("Se esperaba 'Success' = true, pero se obtuvo: %v", response["Success"])
	}
	if response["Data"] == nil {
		t.Errorf("Se esperaba que la respuesta contuviera los datos de la evaluación actualizada")
	}
	t.Log("Respuesta de la solicitud PUT:", response)
	t.Log("Test PUT Evaluacion finalizado correctamente.")
}
func TestGetAllEvaluacion(t *testing.T) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://localhost:8080/v1/evaluacion", nil)
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
	t.Log("Test GET All Evaluacion finalizado correctamente.")
}

func TestGetOneEvaluacion(t *testing.T) {
	req, err := http.NewRequest("GET", "http://localhost:8080/v1/evaluacion/16", nil)
	if err != nil {
		t.Fatalf("Error al crear la solicitud GET: %v", err)
	}
	client := &http.Client{}
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
	if response["Success"] != true {
		t.Errorf("Se esperaba 'Success' = true, pero se obtuvo: %v", response["Success"])
	}
	if response["Data"] == nil {
		t.Errorf("Se esperaba que la respuesta contuviera los datos de la evaluación con ID 16")
	}
	t.Log("Respuesta de la solicitud GET:", response)
	t.Log("Test GET Evaluacion con ID 16 finalizado correctamente.")
}

func TestDeleteEvaluacion(t *testing.T) {
	req, err := http.NewRequest("DELETE", "http://localhost:8080/v1/evaluacion/16", nil)
	if err != nil {
		t.Fatalf("Error al crear la solicitud DELETE: %v", err)
	}
	client := &http.Client{}
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
	if response["Success"] != true {
		t.Errorf("Se esperaba 'Success' = true, pero se obtuvo: %v", response["Success"])
	}
	if response["Message"] != "Delete successful" {
		t.Errorf("Se esperaba 'Message' = 'Delete successful', pero se obtuvo: %v", response["Message"])
	}
	t.Log("Respuesta de la solicitud DELETE:", response)
	t.Log("Test DELETE Evaluacion con ID 16 finalizado correctamente.")
}
