package tests

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"testing"

	"github.com/udistrital/evaluacion_cumplido_prov_crud/models"
)

func TestPostRolAsignacionEvaluador(t *testing.T) {
	rolAsignacionEvaluador := models.RolAsignacionEvaluador{
		Nombre:            "Evaluador Principal",
		Descripcion:       "Este rol es para el evaluador principal",
		CodigoAbreviacion: "EP",
		Activo:            true,
	}

	body, err := json.Marshal(rolAsignacionEvaluador)
	if err != nil {
		t.Fatalf("Error al convertir el cuerpo de la solicitud a JSON: %v", err)
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", "http://localhost:8080/v1/rol_asignacion_evaluador/", bytes.NewBuffer(body))
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
		t.Errorf("Se esperaba un objeto 'Data', pero no se obtuvo.")
	} else {
		data := response["Data"].(map[string]interface{})
		if data["Nombre"] != "Evaluador Principal" {
			t.Errorf("Se esperaba 'Nombre' = 'Evaluador Principal', pero se obtuvo: %v", data["Nombre"])
		}
		if data["CodigoAbreviacion"] != "EP" {
			t.Errorf("Se esperaba 'CodigoAbreviacion' = 'EP', pero se obtuvo: %v", data["CodigoAbreviacion"])
		}
	}
	t.Log("Respuesta de put:", response)
	t.Log("Test POST RolAsignacionEvaluador finalizado correctamente.")
}

func TestGetAllRolAsignacionEvaluador(t *testing.T) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://localhost:8080/v1/rol_asignacion_evaluador", nil)
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
		t.Errorf("Se esperaba una lista de 'Data', pero no se obtuvo.")
	} else {
		dataList, ok := response["Data"].([]interface{})
		if !ok {
			t.Errorf("Se esperaba que 'Data' fuera una lista de objetos, pero no lo es.")
		} else {
			for i, data := range dataList {
				item, ok := data.(map[string]interface{})
				if !ok {
					t.Errorf("El elemento %d de la lista no es un objeto esperado", i)
				} else {
					if item["Nombre"] == nil {
						t.Errorf("El elemento %d de la lista no tiene el campo 'Nombre'", i)
					}
					if item["CodigoAbreviacion"] == nil {
						t.Errorf("El elemento %d de la lista no tiene el campo 'CodigoAbreviacion'", i)
					}
				}
			}
		}
	}

	t.Log("Respuesta de la solicitud GET All:", response)
	t.Log("Test GET All ResultadoEvaluacion finalizado correctamente.")
}

func TestGetOneRolAsignacionEvaluador(t *testing.T) {
	id := 3

	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://localhost:8080/v1/rol_asignacion_evaluador/"+strconv.Itoa(id), nil)
	if err != nil {
		t.Fatalf("Error al crear la solicitud GET: %v", err)
	}

	r, err := client.Do(req)
	if err != nil {
		t.Fatalf("Error al ejecutar la solicitud GET: %v", err)
	}
	defer r.Body.Close()

	t.Log("Código de estado:", r.StatusCode)
	t.Log("Cuerpo de la respuesta:")
	body, _ := io.ReadAll(r.Body)
	t.Log(string(body))

	if r.StatusCode != http.StatusOK {
		t.Errorf("Se esperaba el código de estado 200, pero se obtuvo: %d", r.StatusCode)
	}

	var response map[string]interface{}
	if err := json.Unmarshal(body, &response); err != nil {
		t.Fatalf("Error al decodificar la respuesta: %v", err)
	}

	if response["Success"] != true {
		t.Errorf("Se esperaba 'Success' = true, pero se obtuvo: %v", response["Success"])
	}

	if response["Data"] == nil {
		t.Errorf("Se esperaba que la respuesta tuviera 'Data', pero no se obtuvo.")
	} else {
		data, ok := response["Data"].(map[string]interface{})
		if !ok {
			t.Errorf("Se esperaba que 'Data' fuera un objeto, pero no lo es.")
		} else {
			if data["Id"] != float64(id) {
				t.Errorf("Se esperaba que 'Id' sea %d, pero se obtuvo: %v", id, data["Id"])
			}
		}
	}
	t.Log("Respuesta de la solicitud GET one:", response)
	t.Log("Test GET One RolAsignacionEvaluador finalizado correctamente.")
}

func TestPutRolAsignacionEvaluador(t *testing.T) {
	id := 3

	updatedRol := models.RolAsignacionEvaluador{
		Id:                id,
		Nombre:            "Nuevo Rol Actualizado",
		Descripcion:       "Descripción actualizada del rol",
		CodigoAbreviacion: "NRA",
		Activo:            true,
	}

	body, err := json.Marshal(updatedRol)
	if err != nil {
		t.Fatalf("Error al convertir el cuerpo de la solicitud a JSON: %v", err)
	}

	client := &http.Client{}
	req, err := http.NewRequest("PUT", "http://localhost:8080/v1/rol_asignacion_evaluador/"+strconv.Itoa(id), bytes.NewBuffer(body))
	if err != nil {
		t.Fatalf("Error al crear la solicitud PUT: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	r, err := client.Do(req)
	if err != nil {
		t.Fatalf("Error al ejecutar la solicitud PUT: %v", err)
	}
	defer r.Body.Close()

	t.Log("Código de estado:", r.StatusCode)
	t.Log("Cuerpo de la respuesta:")
	bodyResponse, _ := io.ReadAll(r.Body)
	t.Log(string(bodyResponse))

	if r.StatusCode != http.StatusOK {
		t.Errorf("Se esperaba el código de estado 200, pero se obtuvo: %d", r.StatusCode)
	}

	var response map[string]interface{}
	if err := json.Unmarshal(bodyResponse, &response); err != nil {
		t.Fatalf("Error al decodificar la respuesta: %v", err)
	}

	if response["Success"] != true {
		t.Errorf("Se esperaba 'Success' = true, pero se obtuvo: %v", response["Success"])
	}

	if response["Data"] == nil {
		t.Errorf("Se esperaba que la respuesta tuviera 'Data', pero no se obtuvo.")
	} else {
		data, ok := response["Data"].(map[string]interface{})
		if !ok {
			t.Errorf("Se esperaba que 'Data' fuera un objeto, pero no lo es.")
		} else {
			if data["Id"] != float64(id) {
				t.Errorf("Se esperaba que 'Id' sea %d, pero se obtuvo: %v", id, data["Id"])
			}
			if data["Nombre"] != updatedRol.Nombre {
				t.Errorf("Se esperaba que 'Nombre' sea '%s', pero se obtuvo: %v", updatedRol.Nombre, data["Nombre"])
			}
			if data["Descripcion"] != updatedRol.Descripcion {
				t.Errorf("Se esperaba que 'Descripcion' sea '%s', pero se obtuvo: %v", updatedRol.Descripcion, data["Descripcion"])
			}
			if data["CodigoAbreviacion"] != updatedRol.CodigoAbreviacion {
				t.Errorf("Se esperaba que 'CodigoAbreviacion' sea '%s', pero se obtuvo: %v", updatedRol.CodigoAbreviacion, data["CodigoAbreviacion"])
			}
			if data["Activo"] != updatedRol.Activo {
				t.Errorf("Se esperaba que 'Activo' sea %v, pero se obtuvo: %v", updatedRol.Activo, data["Activo"])
			}
		}
	}
	t.Log("Respuesta de la solicitud put :", response)
	t.Log("Test PUT RolAsignacionEvaluador finalizado correctamente.")
}

func TestDeleteRolAsignacionEvaluador(t *testing.T) {
	client := &http.Client{}
	req, err := http.NewRequest("DELETE", "http://localhost:8080/v1/rol_asignacion_evaluador/8", nil)
	if err != nil {
		t.Fatalf("Error al crear la solicitud DELETE: %v", err)
	}

	r, err := client.Do(req)
	if err != nil {
		t.Fatalf("Error al ejecutar la solicitud DELETE: %v", err)
	}
	defer r.Body.Close()

	t.Log("Código de estado:", r.StatusCode)
	t.Log("Cuerpo de la respuesta:")
	bodyResponse, _ := io.ReadAll(r.Body)
	t.Log(string(bodyResponse))

	if r.StatusCode != http.StatusOK {
		t.Errorf("Se esperaba el código de estado 200, pero se obtuvo: %d", r.StatusCode)
	}

	var response map[string]interface{}
	if err := json.Unmarshal(bodyResponse, &response); err != nil {
		t.Fatalf("Error al decodificar la respuesta: %v", err)
	}

	if response["Success"] != true {
		t.Errorf("Se esperaba 'Success' = true, pero se obtuvo: %v", response["Success"])
	}

	t.Log("Test DELETE RolAsignacionEvaluador finalizado correctamente.")
}
