package models

import (
	"time"
)

type Tipo struct {
	//gorm.Model
	ID        uint      `gorm:"primaryKey"`
	Nombre    string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	DeletedAt time.Time `gorm:"index"`
	//
	Raza []Raza `gorm:"foreignKey:tipo_id;"`
}

type Tipos []*Tipo
