package models

type Raza struct {
	Nombre      string `bson:"nombre" json:"nombre"`
	Tipo        Tipo   `bson:"tipo" json:"tipo"`
	Descripcion string `bson:"descripcion" json:"descripcion"`
}

type Razas []Raza
