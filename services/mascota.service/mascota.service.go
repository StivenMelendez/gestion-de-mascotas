package mascota_service

import (
	"gestion-de-mascotas/models"
	mar "gestion-de-mascotas/repositories/mascota.repository"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"gorm.io/gorm"
)

type MascotaService struct {
	DB *gorm.DB
}

func NewMascotaService(db *gorm.DB) *MascotaService {
	return &MascotaService{DB: db}
}

func (ms *MascotaService) Set(c echo.Context, mascota models.Mascota) error {
	err := mar.Set(mascota)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK,
		map[string]string{"message": "mascota insertada con exito"})
}

func (ms *MascotaService) Get(c echo.Context) error {
	mascotas, err := mar.Get()

	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": "error al obtener las mascotas"})
	}

	return c.JSON(http.StatusOK, mascotas)
}

func (ms *MascotaService) Update(c echo.Context) error {
	mascotaID := c.Param("id")
	id, err := strconv.Atoi(mascotaID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "ID inválido"})
	}

	var mascota models.Mascota
	if err := c.Bind(&mascota); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Datos inválidos"})
	}

	err = mar.Update(mascota, uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Error al actualizar la mascota"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Mascota actualizada con éxito"})
}

func (ms *MascotaService) Delete(c echo.Context) error {
	mascota_id_str := c.Param("id")

	mascota_id, err := strconv.ParseInt(mascota_id_str, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "ID de mascota invalido"})
	}

	err = mar.Delete(uint(mascota_id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": "error al eliminar la mascota"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "mascota eliminada con exito"})
}
