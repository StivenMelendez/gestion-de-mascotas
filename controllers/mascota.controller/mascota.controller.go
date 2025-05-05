package mascota_controller

import (
	"context"
	"fmt"
	"gestion-de-mascotas/database"
	m "gestion-de-mascotas/models"
	"log"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	Collection = database.GetCollection("mascotas")
	ctx        = context.Background()
)

func Set(mascota m.Mascota) error {
	mascota.ID = primitive.NewObjectID()
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

func GetByDuenoID(duenoID string) (m.Mascotas, error) {
	var mascotas m.Mascotas

	// Convertir duenoID a número si es necesario
	duenoIDInt, err := strconv.Atoi(duenoID)
	if err != nil {
		return nil, fmt.Errorf("dueno_id inválido: %s", duenoID)
	}

	filter := bson.M{"dueno_id": duenoIDInt}
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
	log.Println("Mascotas encontradas:", mascotas)
	return mascotas, nil
}

func Update(mascota m.Mascota, mascotaID primitive.ObjectID) error {
	filter := bson.M{"_id": mascotaID}
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

func Delete(mascotaID primitive.ObjectID) error {
	filter := bson.M{"_id": mascotaID}
	result, err := Collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return fmt.Errorf("no se encontró ningún documento con el id %s", mascotaID.Hex())
	}
	return nil
}
