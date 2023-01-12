/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"participle-test/parser"
	"participle-test/parser/create"
	"participle-test/parser/drop"
	"participle-test/parser/insert"
	"participle-test/parser/query"
	"participle-test/storage/mongodb"

	"github.com/spf13/cobra"
)

// translateCmd represents the translate command
var translateCmd = &cobra.Command{
	Use:   "translate",
	Short: "Translate SQL to Mongo",
	Run: func(cmd *cobra.Command, args []string) {
		// Parsear el archivo
		filename, _ := cmd.Flags().GetString("file")
		parsed := parser.ParseFile(filename)

		// Traducir segun el tipo de sentencia
		for _, stmt := range parsed.Statements {
			if stmt.StmtType() == "CREATE" {
				createStmt := stmt.(create.CreateStmt)
				fmt.Println(mongodb.TranslateCreate(createStmt))
			} else if stmt.StmtType() == "DROP" {
				dropStmt := stmt.(drop.DropStmt)
				fmt.Println(mongodb.TranslateDrop(dropStmt))
			} else if stmt.StmtType() == "INSERT" {
				insertStmt := stmt.(insert.InsertStmt)
				fmt.Println(mongodb.TranslateInsert(insertStmt))
			} else if stmt.StmtType() == "SELECT" {
				selectStmt := stmt.(query.SelectStmt)
				fmt.Println(mongodb.TranslateSelect(selectStmt))
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(translateCmd)

	// Archivo a traducir
	translateCmd.Flags().StringP("file", "f", "", "Name of the SQL file")
	translateCmd.MarkFlagRequired("file")
}
