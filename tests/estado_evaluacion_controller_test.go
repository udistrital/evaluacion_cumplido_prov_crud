package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

func TestCreateEstadoEvaluacion(t *testing.T) {
	client := &http.Client{}
	estadoEvaluacion := map[string]interface{}{
		"Nombre":            "Evaluación en proceso",
		"CodigoAbreviacion": "EP",
		"Descripcion":       "Descripción de la evaluación en proceso",
		"Activo":            true,
	}
	jsonData, err := json.Marshal(estadoEvaluacion)
	if err != nil {
		t.Fatalf("Error al convertir el objeto a JSON: %v", err)
	}

	req, err := http.NewRequest("POST", "http://localhost:8080/v1/estado_evaluacion", bytes.NewBuffer(jsonData))
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
	t.Log("Test POST EstadoEvaluacion finalizado correctamente.")
}
func TestGetOneEstadoEvaluacion(t *testing.T) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://localhost:8080/v1/estado_evaluacion/16", nil)
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
	t.Log("Test GET EstadoEvaluacion finalizado correctamente.")
}

func TestGetAllEstadoEvaluacion(t *testing.T) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://localhost:8080/v1/estado_evaluacion", nil)
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
	t.Log("Test GET All EstadoEvaluacion finalizado correctamente.")
}
func TestPutEstadoEvaluacion(t *testing.T) {
	client := &http.Client{}
	estadoEvaluacion := map[string]interface{}{
		"Nombre":            "Evaluación Finalizada",
		"CodigoAbreviacion": "EF",
		"Descripcion":       "Descripción de la evaluación finalizada",
		"Activo":            false,
	}

	// Convertir el objeto a JSON
	jsonData, err := json.Marshal(estadoEvaluacion)
	if err != nil {
		t.Fatalf("Error al convertir el objeto a JSON: %v", err)
	}

	req, err := http.NewRequest("PUT", "http://localhost:8080/v1/estado_evaluacion/16", bytes.NewBuffer(jsonData))
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
	t.Log("Test PUT EstadoEvaluacion finalizado correctamente.")
}

func TestDeleteEstadoEvaluacion(t *testing.T) {
	client := &http.Client{}
	req, err := http.NewRequest("DELETE", "http://localhost:8080/v1/estado_evaluacion/16", nil)
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
	t.Log("Test DELETE EstadoEvaluacion finalizado correctamente.")
}
