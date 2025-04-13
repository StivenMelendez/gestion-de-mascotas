package models

import (
	"time"

	"gorm.io/gorm"
)

type Raza struct {
	ID        uint           `gorm:"primaryKey"`
	Nombre    string         `json:"nombre"`
	Tipo_id   uint           `json:"tipo_id"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	//
	Mascotas []Mascota `gorm:"foreignKey:raza_id;"`
	Tipo     Tipo      `gorm:"foreignKey:tipo_id;references:id"`
}

type Razas []Raza //lista
