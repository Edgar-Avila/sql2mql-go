package mongodb

import (
	"context"
	"fmt"
	"log"
	"participle-test/parser/create"

	"go.mongodb.org/mongo-driver/mongo"
)

func Create(db *mongo.Database, createStmt create.CreateStmt) {
	err := db.CreateCollection(context.Background(), createStmt.Name)
	if err != nil {
		log.Fatal(err)
	}
}

func TranslateCreate(createStmt create.CreateStmt) string {
	return fmt.Sprintf("db.createCollection(\"%s\");", createStmt.Name)
}
