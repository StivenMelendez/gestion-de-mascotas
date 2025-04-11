package models

import (
	"time"

	"gorm.io/gorm"
)

type Mascota struct {
	gorm.Model
	Nombre              string
	Raza_id             uint
	Peso                float64
	Dueno_id            uint
	Fecha_de_nacimiento time.Time
	//
	Raza Raza `gorm:"foreignKey:Raza_id;references:ID"`
}
