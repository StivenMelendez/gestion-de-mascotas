package models

import (
	"time"
)

type Mascota struct {
	//ID                uint      `bson:"id" json:"id"`
	Foto              string    `bson:"foto" json:"foto"` // Ruta de la imagen
	Nombre            string    `bson:"nombre" json:"nombre"`
	Raza              Raza      `bson:"raza" json:"raza"`
	Peso              float64   `bson:"peso" json:"peso"`
	DuenoID           uint      `bson:"dueno_id" json:"dueno_id"`
	FechaDeNacimiento time.Time `bson:"fecha_de_nacimiento" json:"fecha_de_nacimiento"`
	CreatedAt         time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt         time.Time `bson:"updated_at" json:"updated_at"`
	DeletedAt         time.Time `bson:"deleted_at,omitempty" json:"deleted_at,omitempty"`
}

type Mascotas []Mascota
