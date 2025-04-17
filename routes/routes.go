package routes

import (
	mascota_service "gestion-de-mascotas/services/mascota.service"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo, mascotaService *mascota_service.MascotaService) {
	e.POST("/mascotas", mascotaService.Set)
	e.GET("/mascotas", mascotaService.Get)
	e.GET("/mascotas/dueno/:dueno_id", mascotaService.GetByDuenoID)
	e.PUT("/mascotas/:id", mascotaService.Update)
	e.DELETE("/mascotas/:id", mascotaService.Delete)

}
