package models

import "gorm.io/gorm"

type Raza struct {
	gorm.Model
	Nombre string `json:"nombre"`
	//
	Mascotas []Mascota `gorm:"foreignKey:raza_id;"`
}
