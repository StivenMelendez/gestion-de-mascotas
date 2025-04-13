package models

import (
	"time"

	"gorm.io/gorm"
)

type Mascota struct {
	//gorm.Model
	ID                  uint           `gorm:"primaryKey"`
	Nombre              string         `gorm:"not null"`
	Tipo_id             uint           `gorm:"not null"`
	Raza_id             uint           `gorm:"not null"`
	Peso                float64        `gorm:"type:decimal(10,2)"`
	Dueno_id            uint           `gorm:"not null"`
	Fecha_de_nacimiento time.Time      `gorm:"format:2006-01-02"`
	CreatedAt           time.Time      `gorm:"autoCreateTime"`
	UpdatedAt           time.Time      `gorm:"autoUpdateTime"`
	DeletedAt           gorm.DeletedAt `gorm:"index"`
	//
	Raza Raza `gorm:"foreignKey:Raza_id;references:ID"`
	Tipo Tipo `gorm:"foreignKey:Tipo_id;references:ID"`
}

type Mascotas []*Mascota
