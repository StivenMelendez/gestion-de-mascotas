package mascota_service_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"gestion-de-mascotas/models"
	mascota_service "gestion-de-mascotas/services/mascota.service"

	"github.com/labstack/echo/v4"
)

func TestSet(t *testing.T) {
	e := echo.New()

	mascota := models.Mascota{
		Foto:              "./images/mascota.jpg",
		Nombre:            "Pepito",
		Raza:              models.Raza{Nombre: "Labrador", Tipo: models.Tipo{Nombre: "Perro"}},
		Peso:              10.5,
		DuenoID:           1,
		FechaDeNacimiento: time.Now(),
	}

	body, _ := json.Marshal(mascota)
	req := httptest.NewRequest(http.MethodPost, "/mascotas", bytes.NewBuffer(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if err := mascota_service.Set(c); err != nil {
		t.Errorf("Error en la prueba de inserción: %v", err)
	}
}

func TestGet(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/mascotas", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if err := mascota_service.Get(c); err != nil {
		t.Errorf("Error en la prueba de obtención: %v", err)
	} else {
		t.Logf("Prueba de obtención correcta")
	}
}

func TestGetByDuenoID(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/mascotas/dueno/1", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.SetParamNames("dueno_id")
	c.SetParamValues("1")

	err := mascota_service.GetByDuenoID(c)
	if err != nil {
		t.Errorf("Error en la prueba de obtención por dueño ID: %v", err)
	}

	if rec.Code != http.StatusOK {
		t.Errorf("Código de estado esperado %d, obtenido %d", http.StatusOK, rec.Code)
	}
}

func TestUpdate(t *testing.T) {
	e := echo.New()

	body := `{
        "foto": "./images/mascota_actualizada.jpg",
        "nombre": "Pedrito",
        "peso": 9.5,
        "raza": {"id": 1, "nombre": "Labrador"},
        "fecha_de_nacimiento": "2020-01-01T00:00:00Z"
    }`

	req := httptest.NewRequest(http.MethodPut, "/mascotas/1", bytes.NewBuffer([]byte(body)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.SetParamNames("id")
	c.SetParamValues("1")

	err := mascota_service.Update(c)
	if err != nil {
		t.Errorf("Error en la prueba de actualización: %v", err)
	}

	if rec.Code != http.StatusOK {
		t.Errorf("Código de estado esperado %d, obtenido %d", http.StatusOK, rec.Code)
	}
}

func TestDelete(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodDelete, "/mascotas/1", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.SetParamNames("id")
	c.SetParamValues("1")

	err := mascota_service.Delete(c)
	if err != nil {
		t.Errorf("Error en la prueba de eliminación: %v", err)
	}

	if rec.Code != http.StatusOK {
		t.Errorf("Código de estado esperado %d, obtenido %d", http.StatusOK, rec.Code)
	}
}
