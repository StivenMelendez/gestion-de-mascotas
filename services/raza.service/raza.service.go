package raza_service

import (
	"encoding/json"
	"fmt"
	rac "gestion-de-mascotas/controllers/raza.controller"
	"gestion-de-mascotas/models"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func Set(filePath string) error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("error al leer el archivo JSON: %v", err)
	}

	var razas []models.Raza
	if err := json.Unmarshal(data, &razas); err != nil {
		return fmt.Errorf("error al decodificar el JSON: %v", err)
	}

	err = rac.Set(razas)
	if err != nil {
		return fmt.Errorf("error al insertar las razas: %v", err)
	}

	return nil
}

func Get(c echo.Context) error {
	tipoID := c.Param("tipo_id") // Obtén el tipo_id como cadena

	razas, err := rac.Get(tipoID) // Pasa el tipo_id directamente
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error al obtener las razas"})
	}
	log.Println("razas en servicio\n", razas)
	return c.JSON(http.StatusOK, razas)
}
