package mascota_service_test

import (
	"bytes"
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

	req := httptest.NewRequest(http.MethodPost, "/mascotas", bytes.NewBuffer([]byte(`{}`)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	fechaNacimiento, err := time.Parse("2006-01-02", "2020-01-01")
	if err != nil {
		t.Fatalf("Failed to parse date: %v", err)
	}

	mascota := models.Mascota{
		ID:                1,
		Nombre:            "Pepito",
		RazaID:            1,
		Peso:              10.5,
		DuenoID:           1,
		FechaDeNacimiento: fechaNacimiento,
		CreatedAt:         time.Now(),
	}

	service := mascota_service.NewMascotaService(nil)

	err2 := service.Set(c, mascota)

	if err2 != nil {
		t.Errorf("Error en la prueba de insercion: %v", err2)
	} else {
		t.Logf("Prueba de insercion correcta")
	}
}

func TestGet(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/mascotas", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	service := mascota_service.NewMascotaService(nil)

	err := service.Get(c)

	if err != nil {
		t.Errorf("Error en la prueba de obtencion: %v", err)
	} else {
		t.Logf("Prueba de obtencion correcta")
	}
}
func TestUpdate(t *testing.T) {
	e := echo.New()

	body := `{
        "nombre": "Pedrito",
        "peso": 9.5,
        "raza_id": 1,
        "fecha_de_nacimiento": "2020-01-01T00:00:00Z"
    }`

	req := httptest.NewRequest(http.MethodPut, "/mascotas/1", bytes.NewBuffer([]byte(body)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.SetParamNames("id")
	c.SetParamValues("1")

	service := mascota_service.NewMascotaService(nil)

	err := service.Update(c)
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

	service := mascota_service.NewMascotaService(nil)

	err := service.Delete(c)
	if err != nil {
		t.Errorf("Error en la prueba de eliminación: %v", err)
	}

	if rec.Code != http.StatusOK {
		t.Errorf("Código de estado esperado %d, obtenido %d", http.StatusOK, rec.Code)
	}
}
