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
		// Flags
		filename, _ := cmd.Flags().GetString("file")
		dbname, _ := cmd.Flags().GetString("dbname")
		uri, _ := cmd.Flags().GetString("mongouri")

		// Mongo database
		ctx, f := context.WithTimeout(context.Background(), 10*time.Second)
		defer f()
		mongoClient := mongodb.NewClient(ctx, uri)
		defer mongoClient.Disconnect(context.Background())
		db := mongoClient.Database(dbname)

		// Parse
		parsed := parser.ParseFile(filename)

		// Execute commands
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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// executeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	executeCmd.Flags().StringP("file", "f", "", "Name of the SQL file")
	executeCmd.Flags().StringP("mongouri", "u", "mongodb://localhost:27017/", "Uri for the Mongo database")
	executeCmd.Flags().StringP("dbname", "d", "sqlmql-test", "Name of the database")
	executeCmd.MarkFlagRequired("file")
}
