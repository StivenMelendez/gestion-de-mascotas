package mascota_service

import (
	mac "gestion-de-mascotas/controllers/mascota.controller"
	"gestion-de-mascotas/models"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func Set(c echo.Context) error {
	var mascota models.Mascota
	if err := c.Bind(&mascota); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Datos inválidos"})
	}

	if err := mac.Set(mascota); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Error al insertar la mascota"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Mascota insertada con éxito"})
}

func Get(c echo.Context) error {
	mascotas, err := mac.Get()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Error al obtener las mascotas"})
	}

	return c.JSON(http.StatusOK, mascotas)
}

func GetByDuenoID(c echo.Context) error {
	duenoIDStr := c.Param("dueno_id")
	log.Println("Solicitud recibida para dueno_id:", duenoIDStr)
	duenoID, err := strconv.Atoi(duenoIDStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "ID de dueño inválido"})
	}

	mascotas, err := mac.GetByDuenoID(uint(duenoID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Error al obtener las mascotas"})
	}

	return c.JSON(http.StatusOK, mascotas)
}

func Update(c echo.Context) error {
	mascotaID := c.Param("id")
	id, err := strconv.Atoi(mascotaID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "ID inválido"})
	}

	var mascota models.Mascota
	if err := c.Bind(&mascota); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Datos inválidos"})
	}

	err = mac.Update(mascota, uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Error al actualizar la mascota"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Mascota actualizada con éxito"})
}

func Delete(c echo.Context) error {
	mascotaIDStr := c.Param("id")
	mascotaID, err := strconv.Atoi(mascotaIDStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "ID inválido"})
	}

	err = mac.Delete(uint(mascotaID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Error al eliminar la mascota"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Mascota eliminada con éxito"})
}
