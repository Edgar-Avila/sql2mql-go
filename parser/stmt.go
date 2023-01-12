package parser

import (
	"participle-test/parser/create"
	"participle-test/parser/drop"
	"participle-test/parser/insert"
	"participle-test/parser/query"

	"github.com/alecthomas/participle/v2"
)

// Una sentencia generica
type Statement interface{ StmtType() string }

// Union de todos los tipos de sentencias
func StatementUnion() participle.Option {
	return participle.Union[Statement](
		create.CreateStmt{},
		drop.DropStmt{},
		insert.InsertStmt{},
		query.SelectStmt{},
	)
}
