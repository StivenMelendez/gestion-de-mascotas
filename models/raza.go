package models

type Raza struct {
	ID          uint   `bson:"id" json:"id"`
	Nombre      string `bson:"nombre" json:"nombre"`
	TipoID      uint   `bson:"tipo_id" json:"tipo_id"`
	Descripcion string `bson:"descripcion" json:"descripcion"`
}

type Razas []Raza
