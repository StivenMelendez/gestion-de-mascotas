package mascota_controller

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

func Set(mascota m.Mascota) error {
	_, err := Collection.InsertOne(ctx, mascota)
	return err
}

func Get() (m.Mascotas, error) {
	var mascotas m.Mascotas
	cursor, err := Collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var mascota m.Mascota
		if err := cursor.Decode(&mascota); err != nil {
			return nil, err
		}
		mascotas = append(mascotas, mascota)
	}

	return mascotas, nil
}

/*func GetByName(mascotaID uint) (m.Mascota, error) {
	var Mascota m.Mascota

	filter := bson.M{"id": mascotaID}

	err := Collection.FindOne(ctx, filter).Decode(&Mascota)
	if err != nil {
		return Mascota, err
	}

	return Mascota, nil
}*/

func GetByDuenoID(duenoID uint) (m.Mascotas, error) {
	var mascotas m.Mascotas
	filter := bson.M{"dueno_id": duenoID}
	cursor, err := Collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var mascota m.Mascota
		if err := cursor.Decode(&mascota); err != nil {
			return nil, err
		}
		mascotas = append(mascotas, mascota)
	}

	return mascotas, nil
}

func Update(mascota m.Mascota, mascotaID uint) error {
	filter := bson.M{"id": mascotaID}
	update := bson.M{
		"$set": bson.M{
			"foto":                mascota.Foto,
			"nombre":              mascota.Nombre,
			"raza":                mascota.Raza,
			"peso":                mascota.Peso,
			"fecha_de_nacimiento": mascota.FechaDeNacimiento,
			"updated_at":          time.Now(),
		},
	}
	_, err := Collection.UpdateOne(ctx, filter, update)
	return err
}

func Delete(mascotaID uint) error {
	filter := bson.M{"id": mascotaID}
	result, err := Collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return fmt.Errorf("no se encontró ningún documento con el id %d", mascotaID)
	}
	return nil
}
