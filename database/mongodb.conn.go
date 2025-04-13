package database

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	database = "gestion_de_mascotas"
	//user     = ""
	//password = ""
	host = "localhost"
	port = "27017"
)

func GetCollection(collectionName string) *mongo.Collection {
	uri := fmt.Sprintf("mongodb://%s:%s", host, port)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(fmt.Sprintf("Error al conectar con MongoDB: %v", err))
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		panic(fmt.Sprintf("No se pudo conectar a MongoDB: %v", err))
	}

	return client.Database(database).Collection(collectionName)
}
