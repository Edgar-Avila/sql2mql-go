package mongodb

import (
	"context"
	"fmt"
	"log"
	"participle-test/parser/query"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Find(db *mongo.Database, selectStmt query.SelectStmt) []bson.M {
	collection := db.Collection(selectStmt.From)
	opts := options.Find()

	// Agregation

	// Limit
	opts.Limit = &selectStmt.Limit

	// Sort (ORDER BY)
	sort := bson.M{}
	for _, sortSpec := range selectStmt.OrderBy {
		if sortSpec.Dir == "ASC" {
			sort[sortSpec.Col] = 1
		} else {
			sort[sortSpec.Col] = -1
		}
	}
	opts.Sort = sort

	// Projection (SELECT <Columns>)
	if selectStmt.Cols[0] != "*" {
		projection := bson.M{}
		projection["_id"] = 0
		for _, col := range selectStmt.Cols {
			projection[col] = 1
		}
		opts.Projection = projection
	}

	// Aggregation (GROUP BY)

	// Find results
	cursor, err := collection.Find(context.TODO(), bson.M{}, opts)
	if err != nil {
		log.Fatal(err)
	}

	// Put results in a map and return them
	results := make([]bson.M, 0)
	err = cursor.All(context.TODO(), &results)
	if err != nil {
		log.Fatal(err)
	}
	return results
}

func TranslateSelect(selectStmt query.SelectStmt) string {
	// db.students.find({}, {_id: 0, name: 1}).sort({name: -1})
	name := selectStmt.From
	proj := ""
	if selectStmt.Cols[0] != "*" {
		projArr := []string{"_id: 0"}
		for _, col := range selectStmt.Cols {
			projArr = append(projArr, fmt.Sprintf("%s: 1", col))
		}
		proj = strings.Join(projArr, ", ")
	}
	translated := fmt.Sprintf("db.%s.find({%s})", name, proj)

	if len(selectStmt.OrderBy) > 0 {
		sortArr := make([]string, 0)
		for _, spec := range selectStmt.OrderBy {
			dir := 1
			if spec.Dir == "DESC" {
				dir = -1
			}
			sortArr = append(sortArr, fmt.Sprintf("%s: %v", spec.Col, dir))
		}
		translated += fmt.Sprintf(".sort({%s})", strings.Join(sortArr, ", "))
	}

	if selectStmt.Limit > 0 {
		translated += fmt.Sprintf(".limit(%v)", selectStmt.Limit)
	}
	translated += ";"

	return translated
}
