package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Mascota struct {
	ID     primitive.ObjectID `bson:"_id" json:"_id"`
	Foto   string             `bson:"foto" json:"foto"`
	Nombre string             `bson:"nombre" json:"nombre"`
	Raza   Raza               `bson:"raza" json:"raza"`
	//Activo            bool               `bson:"activo" json:"activo"`
	Peso              float64   `bson:"peso" json:"peso"`
	DuenoID           uint      `bson:"dueno_id" json:"dueno_id"`
	FechaDeNacimiento time.Time `bson:"fecha_de_nacimiento" json:"fecha_de_nacimiento"`
	CreatedAt         time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt         time.Time `bson:"updated_at" json:"updated_at"`
	DeletedAt         time.Time `bson:"deleted_at,omitempty" json:"deleted_at,omitempty"`
}

type Mascotas []Mascota // Mascotas es un slice de Mascota
