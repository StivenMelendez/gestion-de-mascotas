package routes

import (
	mascota_service "gestion-de-mascotas/services/mascota.service"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	// Rutas para el servicio de mascotas
	e.POST("/mascotas/crear/:dueno_id", mascota_service.Set)         // Crear una nueva mascota
	e.GET("/mascotas", mascota_service.Get)                          // Obtener todas las mascotas
	e.GET("/mascotas/dueno/:dueno_id", mascota_service.GetByDuenoID) // Obtener mascotas por ID del dueño
	e.PUT("/mascotas/update/:_id", mascota_service.Update)           // Actualizar una mascota
	e.DELETE("/mascotas/delete/:_id", mascota_service.Delete)        // Eliminar una mascota
}
