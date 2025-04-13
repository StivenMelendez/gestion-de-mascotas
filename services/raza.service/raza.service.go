package raza_service

import (
	"encoding/json"
	"fmt"
	"gestion-de-mascotas/models"
	raza_repository "gestion-de-mascotas/repositories/raza.repository"
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

	err = raza_repository.Set(razas)
	if err != nil {
		return fmt.Errorf("error al insertar las razas: %v", err)
	}

	return nil
}

func Get(tipoID uint) (models.Razas, error) {
	razas, err := raza_repository.Get(tipoID)
	if err != nil {
		return nil, fmt.Errorf("error al obtener las razas por tipo: %v", err)
	}
	return razas, nil
}
