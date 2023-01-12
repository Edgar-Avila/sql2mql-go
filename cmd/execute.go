/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"participle-test/parser"
	"participle-test/parser/create"
	"participle-test/parser/drop"
	"participle-test/parser/insert"
	"participle-test/parser/query"
	"participle-test/storage/mongodb"
	"participle-test/util"
	"time"

	"github.com/spf13/cobra"
)

// executeCmd represents the execute command
var executeCmd = &cobra.Command{
	Use:   "execute",
	Short: "Execute a Mongo query using a SQL file",
	Run: func(cmd *cobra.Command, args []string) {
		// Sacar los flags (Archivo, nombre de BD y URI)
		filename, _ := cmd.Flags().GetString("file")
		dbname, _ := cmd.Flags().GetString("dbname")
		uri, _ := cmd.Flags().GetString("mongouri")

		// Conectar a mongo
		ctx, f := context.WithTimeout(context.Background(), 10*time.Second)
		defer f()
		mongoClient := mongodb.NewClient(ctx, uri)
		defer mongoClient.Disconnect(context.Background())
		db := mongoClient.Database(dbname)

		// Parsear el archivo
		parsed := parser.ParseFile(filename)

		// Seleccionar como traducir segun el tipo de sentencia
		for _, stmt := range parsed.Statements {
			if stmt.StmtType() == "CREATE" {
				createStmt := stmt.(create.CreateStmt)
				mongodb.Create(db, createStmt)
			} else if stmt.StmtType() == "DROP" {
				dropStmt := stmt.(drop.DropStmt)
				mongodb.Delete(db, dropStmt)
			} else if stmt.StmtType() == "SELECT" {
				selectStmt := stmt.(query.SelectStmt)
				found := mongodb.Find(db, selectStmt)
				fmt.Println(util.Prettify(found))
			} else if stmt.StmtType() == "INSERT" {
				insertStmt := stmt.(insert.InsertStmt)
				mongodb.Insert(db, insertStmt)
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(executeCmd)

	// Archivo a ejecutar
	executeCmd.Flags().StringP("file", "f", "", "Name of the SQL file")

	// Uri de la base de datos (Por defecto "mongodb://localhost:27017/")
	executeCmd.Flags().StringP("mongouri", "u", "mongodb://localhost:27017/", "Uri for the Mongo database")

	// Nombre de la base de datos (Por defecto "sqlql-test")
	executeCmd.Flags().StringP("dbname", "d", "sqlmql-test", "Name of the database")

	// El archivo es obligatorio
	executeCmd.MarkFlagRequired("file")
}
