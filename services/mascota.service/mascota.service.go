package mascota_service

import (
	mac "gestion-de-mascotas/controllers/mascota.controller"
	"gestion-de-mascotas/models"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Set(c echo.Context) error {
	var mascota models.Mascota
	if err := c.Bind(&mascota); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Datos inválidos"})
	}

	// Validar campos obligatorios
	if mascota.Nombre == "" || mascota.Raza.Nombre == "" || mascota.Peso <= 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Faltan campos obligatorios"})
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

	mascotas, err := mac.GetByDuenoID(duenoIDStr) // Usar el ID del dueño como string
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Error al obtener las mascotas"})
	}

	return c.JSON(http.StatusOK, mascotas)
}

func Update(c echo.Context) error {
	// Obtener el ObjectId desde la ruta
	mascotaID := c.Param("_id")
	objectID, err := primitive.ObjectIDFromHex(mascotaID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "ID inválido"})
	}

	var mascota models.Mascota
	if err := c.Bind(&mascota); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Datos inválidos"})
	}

	// Pasar el ObjectId al controlador
	err = mac.Update(mascota, objectID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Error al actualizar la mascota"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Mascota actualizada con éxito"})
}

func Delete(c echo.Context) error {
	// Obtener el ObjectId desde la ruta
	mascotaID := c.Param("_id")
	objectID, err := primitive.ObjectIDFromHex(mascotaID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "ID inválido"})
	}

	// Pasar el ObjectId al controlador
	err = mac.Delete(objectID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Error al eliminar la mascota"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Mascota eliminada con éxito"})
}
