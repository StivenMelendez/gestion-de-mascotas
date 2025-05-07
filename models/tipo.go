package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Tipo struct {
	ID          primitive.ObjectID `bson:"_id" json:"_id"`
	Nombre      string             `bson:"nombre" json:"nombre"`
	Descripcion string             `bson:"descripcion" json:"descripcion"`
}

type Tipos []Tipo
