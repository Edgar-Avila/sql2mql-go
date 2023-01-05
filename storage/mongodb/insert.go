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

func Insert(db *mongo.Database, insertStmt insert.InsertStmt) {
	collection := db.Collection(insertStmt.Into)
	for i := 0; i < len(insertStmt.Rows); i++ {
		if len(insertStmt.Cols) != len(insertStmt.Rows[i].Values) {
			log.Fatal("Length of rows and columns do not match")
		}
		document := bson.M{}
		for j := 0; j < len(insertStmt.Cols); j++ {
			value := insertStmt.Rows[i].Values[j].GetVal()
			document[insertStmt.Cols[j]] = value
		}
		_, err := collection.InsertOne(context.Background(), document)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func TranslateInsert(insertStmt insert.InsertStmt) string {
	translated := fmt.Sprintf("db.%s.insertMany([\n", insertStmt.Into)
	docs := make([]string, 0)
	for i := 0; i < len(insertStmt.Rows); i++ {
		if len(insertStmt.Cols) != len(insertStmt.Rows[i].Values) {
			log.Fatal("Length of rows and columns do not match")
		}
		pairs := make([]string, 0, len(insertStmt.Cols))
		for j := 0; j < len(insertStmt.Cols); j++ {
			col := insertStmt.Cols[j]
			value := insertStmt.Rows[i].Values[j].GetVal()
			val, _ := json.Marshal(value)
			pairs = append(pairs, fmt.Sprintf("%s: %s", col, val))
		}
		docs = append(docs, fmt.Sprintf("  { %s }", strings.Join(pairs, ", ")))
	}
	translated += strings.Join(docs, ",\n") + "\n]);"
	return translated
}
