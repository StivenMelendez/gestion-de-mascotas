package mascota_repository

import (
	"context"
	"fmt"
	"gestion-de-mascotas/database"
	m "gestion-de-mascotas/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

var (
	Collection = database.GetCollection("mascotas")
	ctx        = context.Background()
)

func Set(Mascota m.Mascota) error {
	_, err := Collection.InsertOne(ctx, Mascota)

	if err != nil {
		return err
	}

	return nil

}

func Get() (m.Mascotas, error) {
	var Mascotas m.Mascotas

	filter := bson.D{}
	cursor, err := Collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {
		var Mascota m.Mascota
		err = cursor.Decode(&Mascota)
		if err != nil {
			return nil, err
		}
		Mascotas = append(Mascotas, &Mascota)
	}

	return Mascotas, nil
}

func Update(mascota m.Mascota, mascotaID uint) error {
	filter := bson.M{"id": mascotaID}

	update := bson.M{
		"$set": bson.M{
			"nombre":              mascota.Nombre,
			"peso":                mascota.Peso,
			"raza_id":             mascota.Raza_id,
			"fecha_de_nacimiento": mascota.Fecha_de_nacimiento,
			"updatedat":           time.Now(),
		},
	}

	_, err := Collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func Delete(mascota_id uint) error {
	filter := bson.M{"id": mascota_id}

	result, err := Collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return fmt.Errorf("no se encontró ningún documento con el id %d", mascota_id)
	}
	return nil
}
