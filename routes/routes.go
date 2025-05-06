package routes

import (
	mascota_service "gestion-de-mascotas/services/mascota.service"
	raza_service "gestion-de-mascotas/services/raza.service"
	tipo_service "gestion-de-mascotas/services/tipo.service"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	// Rutas para el servicio de mascotas
	e.POST("/mascotas/crear/:dueno_id", mascota_service.Set)         // Crear una nueva mascota
	e.GET("/mascotas", mascota_service.Get)                          // Obtener todas las mascotas
	e.GET("/mascotas/dueno/:dueno_id", mascota_service.GetByDuenoID) // Obtener mascotas por ID del dueño
	e.PUT("/mascotas/update/:_id", mascota_service.Update)           // Actualizar una mascota
	e.DELETE("/mascotas/delete/:_id", mascota_service.Delete)        // Eliminar una mascota
	//cargado de opciones
	e.GET("/mascotas/tipos", tipo_service.Get)          // Cargar opciones de mascotas desde un archivo JSON
	e.GET("/mascotas/razas/:tipo_id", raza_service.Get) // Obtener todas las razas por tipo de mascota
}
