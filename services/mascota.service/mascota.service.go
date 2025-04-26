package mascota_service

import (
	mac "gestion-de-mascotas/controllers/mascota.controller"
	"gestion-de-mascotas/models"
	"net/http"
	"strconv"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/labstack/echo/v4"
)

type MascotaService struct {
	DB *mongo.Collection
}

/*func NewMascotaService(db *mongo.Collection) *MascotaService {
	return &MascotaService{DB: db}
}*/

func /*(ms *MascotaService)*/ Set(c echo.Context) error {
	var mascota models.Mascota
	if err := c.Bind(&mascota); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Datos inválidos"})
	}

	if err := mac.Set(mascota); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Error al insertar la mascota"})
	}

	return c.JSON(http.StatusOK,
		map[string]string{"message": "mascota insertada con exito"})
}

func /*(ms *MascotaService)*/ Get(c echo.Context) error {
	mascotas, err := mac.Get()

	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": "error al obtener las mascotas"})
	}

	return c.JSON(http.StatusOK, mascotas)
}

func /*(ms *MascotaService)*/ GetByDuenoID(c echo.Context) error {
	duenoIDStr := c.Param("dueno_id")
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

func /*(ms *MascotaService)*/ Update(c echo.Context) error {
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

func /*(ms *MascotaService)*/ Delete(c echo.Context) error {
	mascota_id_str := c.Param("id")

	mascota_id, err := strconv.ParseInt(mascota_id_str, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "ID de mascota invalido"})
	}

	err = mac.Delete(uint(mascota_id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": "error al eliminar la mascota"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "mascota eliminada con exito"})
}
