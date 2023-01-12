package mongodb

import (
	"context"
	"fmt"
	"log"
	"participle-test/parser/drop"

	"go.mongodb.org/mongo-driver/mongo"
)

// Ejecutar DROP TABLE
func Delete(db *mongo.Database, dropStmt drop.DropStmt) {
	// Iterar las tablas
	for _, table := range dropStmt.Tables {
		// Sacar la coleccion que le corresponde a la tabla
		collection := db.Collection(table)

		// Hacer DROP de la collecion
		if err := collection.Drop(context.Background()); err != nil {
			log.Fatal(err)
		}
	}
}

// Traducir DROP
func TranslateDrop(dropStmt drop.DropStmt) string {
	translated := ""
	// Iterar tablas
	for _, name := range dropStmt.Tables {
		// Construir string para hacer drop
		translated += fmt.Sprintf("db.%s.drop();", name)
	}
	return translated
}
