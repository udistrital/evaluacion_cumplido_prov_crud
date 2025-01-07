package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/udistrital/evaluacion_cumplido_prov_crud/models"
)

func TestPostResultadoEvaluacion(t *testing.T) {
	resultadoEvaluacion := models.ResultadoEvaluacion{
		AsignacionEvaluadorId: &models.AsignacionEvaluador{Id: 44},
		ClasificacionId:       &models.Clasificacion{Id: 2},
		ResultadoEvaluacion:   "{\"score\": 95}",
		Observaciones:         "ResultadoOk",
		Activo:                true,
	}

	body, err := json.Marshal(resultadoEvaluacion)
	if err != nil {
		t.Fatalf("Error al convertir el cuerpo de la solicitud a JSON: %v", err)
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", "http://localhost:8080/v1/resultado_evaluacion", bytes.NewBuffer(body))
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

	t.Log("Respuesta de la solicitud POST:", response)
	t.Log("Test POST ResultadoEvaluacion finalizado correctamente.")
}

func TestGetAllResultadoEvaluacion(t *testing.T) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://localhost:8080/v1/resultado_evaluacion", nil)
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

	if response["Success"] != true {
		t.Errorf("Se esperaba 'Success' = true, pero se obtuvo: %v", response["Success"])
	}

	t.Log("Respuesta de la solicitud GET All:", response)
	t.Log("Test GET All ResultadoEvaluacion finalizado correctamente.")
}

func TestGetOneResultadoEvaluacion(t *testing.T) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://localhost:8080/v1/resultado_evaluacion/3", nil)
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

	if response["Success"] != true {
		t.Errorf("Se esperaba 'Success' = true, pero se obtuvo: %v", response["Success"])
	}

	t.Log("Respuesta de la solicitud GET One:", response)
	t.Log("Test GET One ResultadoEvaluacion finalizado correctamente.")
}

func TestDeleteResultadoEvaluacion(t *testing.T) {
	client := &http.Client{}
	req, err := http.NewRequest("DELETE", "http://localhost:8080/v1/resultado_evaluacion/30", nil)
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

	if response["Success"] != true {
		t.Errorf("Se esperaba 'Success' = true, pero se obtuvo: %v", response["Success"])
	}

	t.Log("Respuesta de la solicitud DELETE:", response)
	t.Log("Test DELETE ResultadoEvaluacion finalizado correctamente.")
}
