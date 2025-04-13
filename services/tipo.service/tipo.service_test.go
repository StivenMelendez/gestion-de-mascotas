package tipo_service_test

import (
	tipo_service "gestion-de-mascotas/services/tipo.service"
	"testing"
)

func TestSetTipos(t *testing.T) {
	filePath := "../../default-info/tipos-mascotas.json"

	err := tipo_service.Set(filePath)
	if err != nil {
		t.Errorf("Error al insertar los tipos desde el archivo JSON: %v", err)
	} else {
		t.Logf("Tipos insertados correctamente desde el archivo JSON.")
	}
}

func TestGetTipos(t *testing.T) {
	tipos, err := tipo_service.Get()
	if err != nil {
		t.Errorf("Error al obtener los tipos desde la base de datos: %v", err)
	} else {
		t.Logf("Tipos obtenidos correctamente: %+v", tipos)
	}

	if len(tipos) == 0 {
		t.Errorf("No se obtuvieron tipos desde la base de datos.")
	}
}
