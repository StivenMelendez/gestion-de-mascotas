package tipo_service

import (
	"encoding/json"
	"fmt"
	"gestion-de-mascotas/models"
	tipo_repository "gestion-de-mascotas/repositories/tipo.repository"
	"os"
)

func Set(filePath string) error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("error al leer el archivo JSON: %v", err)
	}

	var tipos []models.Tipo
	if err := json.Unmarshal(data, &tipos); err != nil {
		return fmt.Errorf("error al decodificar el JSON: %v", err)
	}

	err = tipo_repository.Set(tipos)
	if err != nil {
		return fmt.Errorf("error al insertar los tipos: %v", err)
	}

	return nil
}

func Get() (models.Tipos, error) {
	tipos, err := tipo_repository.Get()
	if err != nil {
		return nil, fmt.Errorf("error al obtener los tipos: %v", err)
	}
	return tipos, nil
}
