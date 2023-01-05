package parser

import (
	"participle-test/parser/create"
	"participle-test/parser/drop"
	"participle-test/parser/insert"
	"participle-test/parser/query"

	"github.com/alecthomas/participle/v2"
)

// Statements
type Statement interface{ StmtType() string }

func StatementUnion() participle.Option {
	return participle.Union[Statement](create.CreateStmt{}, drop.DropStmt{}, insert.InsertStmt{}, query.SelectStmt{})
}
