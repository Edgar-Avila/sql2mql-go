package mongodb

import (
	"context"
	"fmt"
	"log"
	"participle-test/parser/drop"

	"go.mongodb.org/mongo-driver/mongo"
)

func Delete(db *mongo.Database, dropStmt drop.DropStmt) {
	for _, table := range dropStmt.Tables {
		collection := db.Collection(table)
		if err := collection.Drop(context.Background()); err != nil {
			log.Fatal(err)
		}
	}
}

func TranslateDrop(dropStmt drop.DropStmt) string {
	translated := ""
	for _, name := range dropStmt.Tables {
		translated += fmt.Sprintf("db.%s.drop();", name)
	}
	return translated
}
