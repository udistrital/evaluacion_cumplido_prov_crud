package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"
	"testing"

	"github.com/udistrital/evaluacion_cumplido_prov_crud/models"
)

func TestPostItem(t *testing.T) {
	item := models.Item{
		EvaluacionId:  &models.Evaluacion{Id: 1},
		Identificador: "12345",
		Nombre:        "Item Test",
		ValorUnitario: 100.0,
		Iva:           19.0,
		FichaTecnica:  "TestFicha",
		Unidad:        1,
		Cantidad:      10.0,
		TipoNecesidad: 1,
		Activo:        true,
	}

	body, err := json.Marshal(item)
	if err != nil {
		t.Fatalf("Error al convertir el cuerpo de la solicitud a JSON: %v", err)
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", "http://localhost:8080/v1/item", bytes.NewBuffer(body))
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
		t.Errorf("Se esperaba que la respuesta contuviera los datos del item creado")
	}

	t.Log("Respuesta de la solicitud POST:", response)
	t.Log("Test POST Item finalizado correctamente.")
}
func TestPutItem(t *testing.T) {
	item := models.Item{
		Id:            200,
		EvaluacionId:  &models.Evaluacion{Id: 1},
		Identificador: "67890",
		Nombre:        "Item Actualizado",
		ValorUnitario: 150.0,
		Iva:           19.0,
		FichaTecnica:  "Ficha Actualizada",
		Unidad:        2,
		Cantidad:      20.0,
		TipoNecesidad: 2,
		Activo:        true,
	}

	body, err := json.Marshal(item)
	if err != nil {
		t.Fatalf("Error al convertir el cuerpo de la solicitud a JSON: %v", err)
	}

	client := &http.Client{}
	req, err := http.NewRequest("PUT", "http://localhost:8080/v1/item/200", bytes.NewBuffer(body))
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
		t.Errorf("Se esperaba que la respuesta contuviera los datos del item actualizado")
	}

	t.Log("Respuesta de la solicitud PUT:", response)
	t.Log("Test PUT Item finalizado correctamente.")
}

func TestGetOneItem(t *testing.T) {
	itemID := 200
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://localhost:8080/v1/item/"+strconv.Itoa(itemID), nil)
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

	if response["Data"] == nil {
		t.Errorf("Se esperaba que la respuesta contuviera los datos del item con ID %d", itemID)
	}

	t.Log("Respuesta de la solicitud GET:", response)
	t.Log("Test GET One Item finalizado correctamente.")
}

func TestGetAllItems(t *testing.T) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://localhost:8080/v1/item", nil)
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

	if response["Data"] == nil {
		t.Errorf("Se esperaba que la respuesta contuviera los datos de los items")
	}

	t.Log("Respuesta de la solicitud GET All:", response)
	t.Log("Test GET All Items finalizado correctamente.")
}
