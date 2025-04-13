package raza_repository

import (
	"context"
	"gestion-de-mascotas/database"
	"gestion-de-mascotas/models"

	"go.mongodb.org/mongo-driver/bson"
)

var (
	Collection = database.GetCollection("razas")
	ctx        = context.Background()
)

func Get(tipoID uint) ([]models.Raza, error) {
	var razas []models.Raza

	filter := bson.M{"tipo_id": tipoID}

	cursor, err := Collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var raza models.Raza
		if err := cursor.Decode(&raza); err != nil {
			return nil, err
		}
		razas = append(razas, raza)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return razas, nil
}
