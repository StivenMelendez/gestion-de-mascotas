package raza_controller

import (
	"context"
	"fmt"
	"gestion-de-mascotas/database"
	"gestion-de-mascotas/models"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
			raza.ID = primitive.NewObjectID()
			_, err := Collection.InsertOne(ctx, raza)
			if err != nil {
				return fmt.Errorf("error al insertar la raza: %v", err)
			}
		}
	}

	return nil
}
func Get(tipoID string) (models.Razas, error) {
	var razas models.Razas

	// Convierte el tipoID a ObjectID
	objectID, err := primitive.ObjectIDFromHex(tipoID)
	if err != nil {
		return nil, fmt.Errorf("tipo_id no es un ObjectID válido: %v", err)
	}

	// Ajusta el filtro para buscar dentro de tipo._id
	filter := bson.M{"tipo._id": objectID}
	log.Printf("Filtro utilizado: %+v", filter)

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
	log.Println("Razas encontradas:", razas)
	return razas, nil
}
