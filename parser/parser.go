package parser

import (
	"log"
	"os"
	"participle-test/parser/common"

	"github.com/alecthomas/participle/v2"
)

// Crear el parser
func NewParser() *participle.Parser[SqlFile] {
	return participle.MustBuild[SqlFile](
		SqlLexer(),
		StatementUnion(),
		common.ValueUnion(),
		participle.CaseInsensitive("Ident"),
	)
}

// Parsear un archivo
func ParseFile(path string) *SqlFile {
	parser := NewParser()

	// Leer archivo
	read, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	// Ejecutar el parseo
	sqlFile, err := parser.ParseBytes("", read)
	if err != nil {
		log.Fatal(err)
	}

	// Devolver estructura parseada
	return sqlFile
}
