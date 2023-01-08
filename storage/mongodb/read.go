package mongodb

import (
	"context"
	"encoding/json"
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

func TranslateFactor(factor query.BoolFactor, surround bool) string {
	translated := ""
	if surround {
		translated += "{"
	}
	if factor.Not {
		translated += "$not:{"
	}
	translated += factor.Cmp.Col
	translated += ":{"
	if factor.Cmp.Op == "=" {
		translated += "$eq:"
	} else if factor.Cmp.Op == "<" {
		translated += "$lt:"
	} else if factor.Cmp.Op == ">" {
		translated += "$gt:"
	} else if factor.Cmp.Op == "<=" {
		translated += "$lte:"
	} else if factor.Cmp.Op == ">=" {
		translated += "$gte:"
	}
	v := factor.Cmp.Lit.GetVal()
	val, _ := json.Marshal(v)
	translated += string(val)
	translated += "}"
	if factor.Not {
		translated += "}"
	}
	if surround {
		translated += "}"
	}
	return translated
}

func TranslateTerm(term query.BoolTerm, surround bool) string {
	translated := ""
	multipleFactors := len(term.Factors) > 1
	if surround {
		translated += "{"
	}
	if multipleFactors {
		translated += "$and:["
	}
	andArr := make([]string, 0)
	for _, factor := range term.Factors {
		factorStr := TranslateFactor(factor, multipleFactors)
		andArr = append(andArr, factorStr)
	}
	translated += strings.Join(andArr, ",")
	if multipleFactors {
		translated += "]"
	}
	if surround {
		translated += "}"
	}
	return translated
}

func TranslateWhere(where *query.Where) string {
	translated := ""
	findArr := make([]string, 0)
	if where != nil {
		multipleTerms := len(where.Terms) > 1
		obj := ""
		if multipleTerms {
			obj += "$or:["
		}
		orArr := make([]string, 0)
		for _, term := range where.Terms {
			termStr := TranslateTerm(term, multipleTerms)
			orArr = append(orArr, termStr)
		}
		obj += strings.Join(orArr, ",")
		if multipleTerms {
			obj += "]"
		}
		findArr = append(findArr, obj)
	}
	translated = fmt.Sprintf("{%s}", strings.Join(findArr, ", "))

	return translated
}

func TranslateProjection(cols []string) string {
	projArr := make([]string, 0)
	if cols[0] != "*" {
		projArr = append(projArr, "_id: 0")
		for _, col := range cols {
			projArr = append(projArr, fmt.Sprintf("%s: 1", col))
		}
	}
	return fmt.Sprintf("{%s}", strings.Join(projArr, ","))
}

func TranslateSort(orderBy []query.SortSpec) string {
	if len(orderBy) > 0 {
		sortArr := make([]string, 0)
		for _, spec := range orderBy {
			dir := 1
			if spec.Dir == "DESC" {
				dir = -1
			}
			sortArr = append(sortArr, fmt.Sprintf("%s: %v", spec.Col, dir))
		}
		return fmt.Sprintf(".sort({%s})", strings.Join(sortArr, ", "))
	}
	return ""
}

func TranslateLimit(limit int64) string {
	if limit > 0 {
		return fmt.Sprintf(".limit(%v)", limit)
	}
	return ""
}

func TranslateSelect(selectStmt query.SelectStmt) string {
	// db.students.find({}, {_id: 0, name: 1}).sort({name: -1})
	name := selectStmt.From
	proj := TranslateProjection(selectStmt.Cols)
	find := TranslateWhere(selectStmt.Where)
	sort := TranslateSort(selectStmt.OrderBy)
	limit := TranslateLimit(selectStmt.Limit)
	translated := fmt.Sprintf("db.%s.find(%s, %s)%s%s;", name, find, proj, sort, limit)
	return translated
}
