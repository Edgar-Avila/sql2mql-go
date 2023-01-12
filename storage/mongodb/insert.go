package mongodb

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"participle-test/parser/insert"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Ejecutar insert
func Insert(db *mongo.Database, insertStmt insert.InsertStmt) {
	// Sacar la collecion que corresponde a la tabla
	collection := db.Collection(insertStmt.Into)

	// Iterar las filas en la sentencia INSERT
	for i := 0; i < len(insertStmt.Rows); i++ {

		// Verificar que el numero de columnas sea el mismo que el esquema
		// INSERT (nombre, edad) -> 2 Valores
		// VALUES ('juan', 17) -> 2 Valores
		if len(insertStmt.Cols) != len(insertStmt.Rows[i].Values) {
			log.Fatal("Length of rows and columns do not match")
		}

		// Construir el documento para insertar
		// {nombre: 'juan', edad: 17}
		document := bson.M{}
		for j := 0; j < len(insertStmt.Cols); j++ {
			value := insertStmt.Rows[i].Values[j].GetVal()
			document[insertStmt.Cols[j]] = value
		}

		// Insertar en BD
		_, err := collection.InsertOne(context.Background(), document)
		if err != nil {
			log.Fatal(err)
		}
	}
}

// Traducir un INSERT
func TranslateInsert(insertStmt insert.InsertStmt) string {
	// String donde se traducira todo
	translated := fmt.Sprintf("db.%s.insertMany([\n", insertStmt.Into)

	// Array de documentos a insertar
	docs := make([]string, 0)

	// Iterar las filas
	for i := 0; i < len(insertStmt.Rows); i++ {

		// Verificar que el numero de columnas sea el mismo que el esquema
		// INSERT (nombre, edad) -> 2 Valores
		// VALUES ('juan', 17) -> 2 Valores
		if len(insertStmt.Cols) != len(insertStmt.Rows[i].Values) {
			log.Fatal("Length of rows and columns do not match")
		}

		// Construir el documento para insertar
		// {nombre: 'juan', edad: 17}
		// En forma de string
		pairs := make([]string, 0, len(insertStmt.Cols))
		for j := 0; j < len(insertStmt.Cols); j++ {
			col := insertStmt.Cols[j]
			value := insertStmt.Rows[i].Values[j].GetVal()
			val, _ := json.Marshal(value)
			pairs = append(pairs, fmt.Sprintf("%s: %s", col, val))
		}

		// Agregar el documento al array de documentos
		docs = append(docs, fmt.Sprintf("  { %s }", strings.Join(pairs, ", ")))
	}

	// Poner los documentos separados por comas
	translated += strings.Join(docs, ",\n") + "\n]);"
	return translated
}
