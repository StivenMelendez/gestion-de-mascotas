package tipo_repository

import (
	"context"
	"fmt"
	"gestion-de-mascotas/database"
	"gestion-de-mascotas/models"

	"go.mongodb.org/mongo-driver/bson"
)

var (
	Collection = database.GetCollection("tipos")
	ctx        = context.Background()
)

func Set(tipos []models.Tipo) error {
	for _, tipo := range tipos {
		filter := bson.M{"id": tipo.ID}
		count, err := Collection.CountDocuments(ctx, filter)
		if err != nil {
			return fmt.Errorf("error al verificar la existencia del tipo: %v", err)
		}

		if count == 0 {
			_, err := Collection.InsertOne(ctx, tipo)
			if err != nil {
				return fmt.Errorf("error al insertar el tipo: %v", err)
			}
		}
	}

	return nil
}

func Get() (models.Tipos, error) {
	var tipos models.Tipos

	// Buscar todos los tipos en la colección
	cursor, err := Collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	// Iterar sobre los resultados y decodificarlos
	for cursor.Next(ctx) {
		var tipo models.Tipo
		if err := cursor.Decode(&tipo); err != nil {
			return nil, err
		}
		tipos = append(tipos, tipo) // Agregar el puntero al slice
	}

	// Verificar si hubo errores durante la iteración
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return tipos, nil
}
