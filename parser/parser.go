package parser

import (
	"log"
	"os"
	"participle-test/parser/insert"

	"github.com/alecthomas/participle/v2"
)

func NewParser() *participle.Parser[SqlFile] {

	return participle.MustBuild[SqlFile](
		SqlLexer(),
		StatementUnion(),
		insert.ValueUnion(),
		participle.CaseInsensitive("Ident"),
	)
}

func ParseFile(path string) *SqlFile {
	parser := NewParser()

	// Open file
	read, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	// Parse string
	sqlFile, err := parser.ParseBytes("", read)
	if err != nil {
		log.Fatal(err)
	}
	return sqlFile
}
