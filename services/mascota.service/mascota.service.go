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
		log.Printf("Error al enlazar los datos de la mascota: %v", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Datos inválidos"})
	}

	log.Printf("Datos de la mascota recibidos: %+v", mascota)

	// Validar campos obligatorios
	if mascota.Nombre == "" || mascota.Raza.Nombre == "" || mascota.Raza.Tipo.Nombre == "" || mascota.Peso <= 0 {
		log.Printf("Faltan campos obligatorios: %+v", mascota)
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Faltan campos obligatorios"})
	}

	// Convertir raza._id a ObjectID
	razaID, err := primitive.ObjectIDFromHex(mascota.Raza.ID.Hex())
	if err != nil {
		log.Printf("El ID de la raza no es válido: %v", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "El ID de la raza no es válido"})
	}

	// Convertir raza.tipo._id a ObjectID
	tipoID, err := primitive.ObjectIDFromHex(mascota.Raza.Tipo.ID.Hex())
	if err != nil {
		log.Printf("El ID del tipo no es válido: %v", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "El ID del tipo no es válido"})
	}

	mascota.Raza.ID = razaID
	mascota.Raza.Tipo.ID = tipoID

	// Llamar al controlador para guardar la mascota
	if err := mac.Set(mascota); err != nil {
		log.Printf("Error al insertar la mascota: %v", err)
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
	var mascota models.Mascota
	if err := c.Bind(&mascota); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Datos inválidos"})
	}
	// Pasar el ObjectId al controlador
	err = mac.Delete(mascota, objectID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Error al eliminar la mascota"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Mascota eliminada con éxito"})
}
