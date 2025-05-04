package raza_controller

import (
	"context"
	"fmt"
	"gestion-de-mascotas/database"
	"gestion-de-mascotas/models"

	"go.mongodb.org/mongo-driver/bson"
)

var (
	Collection = database.GetCollection("razas")
	ctx        = context.Background()
)

func Set(razas []models.Raza) error {
	for _, raza := range razas {
		filter := bson.M{"id": raza}
		count, err := Collection.CountDocuments(ctx, filter)
		if err != nil {
			return fmt.Errorf("error al verificar la existencia de la raza: %v", err)
		}

		if count == 0 {
			_, err := Collection.InsertOne(ctx, raza)
			if err != nil {
				return fmt.Errorf("error al insertar la raza: %v", err)
			}
		}
	}

	return nil
}
func Get(tipoID uint) (models.Razas, error) {
	var razas models.Razas

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
