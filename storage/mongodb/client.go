package mongodb

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Cliente MongoDB
func NewClient(ctx context.Context, uri string) *mongo.Client {
	// Conectarlo usando el URI MongoDB
	opt := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, opt)

	if err != nil {
		log.Fatal(err)
	}

	// Verificar la conexion con un ping
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	return client
}
