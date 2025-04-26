package raza_service

import (
	"encoding/json"
	"fmt"
	rac "gestion-de-mascotas/controllers/raza.controller"
	"gestion-de-mascotas/models"
	"os"
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

func Get(tipoID uint) (models.Razas, error) {
	razas, err := rac.Get(tipoID)
	if err != nil {
		return nil, fmt.Errorf("error al obtener las razas por tipo: %v", err)
	}
	return razas, nil
}
