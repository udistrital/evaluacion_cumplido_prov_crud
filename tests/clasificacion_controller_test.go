package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/udistrital/evaluacion_cumplido_prov_crud/models"
)

func TestPostClasificacion(t *testing.T) {
	item := models.Clasificacion{
		Nombre:            "Test1",
		CodigoAbreviacion: "CT",
		Descripcion:       "Test description",
		LimiteInferior:    1,
		LimiteSuperior:    100,
		Activo:            true,
	}

	body, err := json.Marshal(item)
	if err != nil {
		t.Fatalf("Error al convertir a JSON: %v", err)
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", "http://localhost:8080/v1/clasificacion", bytes.NewReader(body))
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
	t.Log("Test POST Clasificacion finalizado correctamente.")
}

func TestPutClasificacion(t *testing.T) {
	item := models.Clasificacion{
		Id:                7,
		Nombre:            "Clas",
		CodigoAbreviacion: "CAU",
		Descripcion:       "DesctAct",
		LimiteInferior:    5,
		LimiteSuperior:    10,
		Activo:            true,
	}

	body, err := json.Marshal(item)
	if err != nil {
		t.Fatalf("Error al convertir a JSON: %v", err)
	}

	client := &http.Client{}
	req, err := http.NewRequest("PUT", "http://localhost:8080/v1/clasificacion/7", bytes.NewReader(body))
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
	t.Log("Test PUT Clasificacion finalizado correctamente.")
}

func TestGetAllClasificacion(t *testing.T) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://localhost:8080/v1/clasificacion", nil)
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

	t.Log("Respuesta de la solicitud GET All:", response)
	t.Log("Test GET All Clasificacion finalizado correctamente.")
}

func TestGetOneClasificacion(t *testing.T) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://localhost:8080/v1/clasificacion/5", nil)
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

	t.Log("Respuesta de la solicitud GET One:", response)
	t.Log("Test GET One Clasificacion finalizado correctamente.")
}

func TestDeleteClasificacion(t *testing.T) {
	client := &http.Client{}
	req, err := http.NewRequest("DELETE", "http://localhost:8080/v1/clasificacion/5", nil)
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
}
