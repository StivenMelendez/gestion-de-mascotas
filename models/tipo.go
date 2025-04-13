package models

import "gorm.io/gorm"

type Tipo struct {
	gorm.Model
	Nombre string `json:"nombre"`
	//
	Mascotas []Mascota `gorm:"foreignKey:Tipo_id;references:ID"`
}

type Tipos []Tipo
