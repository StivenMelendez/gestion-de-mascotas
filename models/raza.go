package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Raza struct {
	ID     primitive.ObjectID `bson:"_id" json:"_id"`
	Nombre string             `bson:"nombre" json:"nombre"`
	Tipo   Tipo               `bson:"tipo" json:"tipo"`
}

type Razas []Raza
