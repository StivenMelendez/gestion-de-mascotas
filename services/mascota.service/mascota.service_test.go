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

	fechaNacimiento, err := time.Parse("2006-01-02", "2020-01-01")
	if err != nil {
		t.Fatalf("Failed to parse date: %v", err)
	}

	foto := "./images/mascota.jpg"

	mascota := models.Mascota{
		Foto:   foto,
		Nombre: "Pepito",
		Raza: models.Raza{
			Nombre: "Labrador",
			Tipo: models.Tipo{
				Nombre: "Perro",
			},
		},
		Peso:              10.5,
		DuenoID:           1,
		FechaDeNacimiento: fechaNacimiento,
		CreatedAt:         time.Now(),
	}

	// Serializar la mascota a JSON
	mascotaJSON, err := json.Marshal(mascota)
	if err != nil {
		t.Fatalf("Error al serializar la mascota: %v", err)
	}

	// Crear la solicitud con el JSON de la mascota
	req := httptest.NewRequest(http.MethodPost, "/mascotas", bytes.NewBuffer(mascotaJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if err := mascota_service.Set(c); err != nil {
		t.Errorf("Error en la prueba de inserción: %v", err)
	} else {
		t.Logf("Prueba de inserción correcta")
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
