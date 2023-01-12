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

// Ejecutar un SELECT
func Find(db *mongo.Database, selectStmt query.SelectStmt) []bson.M {
	// Hallar collecion que se corresponde con tabla
	collection := db.Collection(selectStmt.From)

	// Array de opciones
	opts := options.Find()

	// Agregacion (GROUP BY) (No implementado)

	// Donde (WHERE) (No implementado)

	// Limite
	opts.Limit = &selectStmt.Limit

	// Orden (ORDER BY)
	sort := bson.M{}
	for _, sortSpec := range selectStmt.OrderBy {
		if sortSpec.Dir == "ASC" {
			sort[sortSpec.Col] = 1
		} else {
			sort[sortSpec.Col] = -1
		}
	}
	opts.Sort = sort

	// Projeccion (SELECT <Columnas>)
	if selectStmt.Cols[0] != "*" {
		projection := bson.M{}
		projection["_id"] = 0
		for _, col := range selectStmt.Cols {
			projection[col] = 1
		}
		opts.Projection = projection
	}

	// Hallar todos los resultados
	cursor, err := collection.Find(context.TODO(), bson.M{}, opts)
	if err != nil {
		log.Fatal(err)
	}

	// Poner resultados en un mapa y retornarlos
	results := make([]bson.M, 0)
	err = cursor.All(context.TODO(), &results)
	if err != nil {
		log.Fatal(err)
	}
	return results
}

// Traducir un factor del WHERE
func TranslateFactor(factor query.BoolFactor, surround bool) string {
	translated := ""
	// Rodearlo de {}
	if surround {
		translated += "{"
	}

	// Si contiene un NOT
	if factor.Not {
		translated += "$not:{"
	}

	// Nombre de la columna
	translated += factor.Cmp.Col

	// Operacion
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

	// Valor
	v := factor.Cmp.Lit.GetVal()
	val, _ := json.Marshal(v)
	translated += string(val)

	// Terminar de rodearlo
	translated += "}"
	if factor.Not {
		translated += "}"
	}
	if surround {
		translated += "}"
	}

	// Retornar
	return translated
}

// Traducir un termino del WHERE
func TranslateTerm(term query.BoolTerm, surround bool) string {
	translated := ""
	multipleFactors := len(term.Factors) > 1

	// Rodearlo de {}
	if surround {
		translated += "{"
	}

	// Array si hay mas de un factor
	if multipleFactors {
		translated += "$and:["
	}

	// Poner los factores
	andArr := make([]string, 0)
	for _, factor := range term.Factors {
		factorStr := TranslateFactor(factor, multipleFactors)
		andArr = append(andArr, factorStr)
	}
	translated += strings.Join(andArr, ",")

	// Terminar de rodearlo
	if multipleFactors {
		translated += "]"
	}
	if surround {
		translated += "}"
	}
	return translated
}

// Traducir un WHERE
func TranslateWhere(where *query.Where) string {
	translated := ""
	findArr := make([]string, 0)

	// Solo si es que el SELECT tiene un WHERE
	if where != nil {
		// Si tiene varios terminos hacer un array
		multipleTerms := len(where.Terms) > 1
		obj := ""
		if multipleTerms {
			obj += "$or:["
		}
		orArr := make([]string, 0)

		// Traducir e insertar terminos
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

	// Traducirlo todo en un documento
	translated = fmt.Sprintf("{%s}", strings.Join(findArr, ", "))

	return translated
}

// Traducir una proyeccion (SELECT <columnas>)
func TranslateProjection(cols []string) string {
	projArr := make([]string, 0)

	// Si no se selecciona todas las columnas
	if cols[0] != "*" {
		// Llenar las columnas seleccionadas
		projArr = append(projArr, "_id: 0")
		for _, col := range cols {
			projArr = append(projArr, fmt.Sprintf("%s: 1", col))
		}
	}
	return fmt.Sprintf("{%s}", strings.Join(projArr, ","))
}

// Traducir un ORDER BY
func TranslateSort(orderBy []query.SortSpec) string {
	// Si es que hay un ORDER BY en el SELECT
	if len(orderBy) > 0 {
		sortArr := make([]string, 0)
		// Ascendente o descendente y la columna correspondiente
		for _, spec := range orderBy {
			dir := 1
			if spec.Dir == "DESC" {
				dir = -1
			}
			sortArr = append(sortArr, fmt.Sprintf("%s: %v", spec.Col, dir))
		}
		// Ponerlo en documento
		return fmt.Sprintf(".sort({%s})", strings.Join(sortArr, ", "))
	}
	return ""
}

// Traducir limit
func TranslateLimit(limit int64) string {
	if limit > 0 {
		return fmt.Sprintf(".limit(%v)", limit)
	}
	return ""
}

// Traducir SELECT
func TranslateSelect(selectStmt query.SelectStmt) string {
	// Ejemplo: db.students.find({}, {_id: 0, name: 1}).sort({name: -1})
	name := selectStmt.From
	proj := TranslateProjection(selectStmt.Cols)
	find := TranslateWhere(selectStmt.Where)
	sort := TranslateSort(selectStmt.OrderBy)
	limit := TranslateLimit(selectStmt.Limit)
	translated := fmt.Sprintf("db.%s.find(%s, %s)%s%s;", name, find, proj, sort, limit)
	return translated
}
