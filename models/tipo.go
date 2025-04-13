package models

type Tipo struct {
	ID          uint   `bson:"id" json:"id"`
	Nombre      string `bson:"nombre" json:"nombre"`
	Descripcion string `bson:"descripcion" json:"descripcion"`
	Activo      bool   `bson:"activo" json:"activo"`
}

type Tipos []Tipo
