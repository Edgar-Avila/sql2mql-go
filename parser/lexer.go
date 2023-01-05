package parser

import (
	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
)

func SqlLexer() participle.Option {
	sqlLexer := lexer.MustSimple([]lexer.SimpleRule{
		{
			Name:    `Into`,
			Pattern: `(?i)INTO`,
		},
		{
			Name:    `NotNull`,
			Pattern: `(?i)NOT\s+NULL`,
		},
		{
			Name:    `PrimaryKey`,
			Pattern: `(?i)PRIMARY\s+KEY`,
		},
		{
			Name:    `Unique`,
			Pattern: `(?i)UNIQUE`,
		},
		{
			Name:    `TextType`,
			Pattern: `(?i)CHAR|VARCHAR|TEXT|TINYTEXT|MEDIUMTEXT|LONGTEXT`,
		},
		{
			Name:    `IntType`,
			Pattern: `(?i)TINYINT|SMALLINT|MEDIUMINT|INT|INTEGER|BIGINT`,
		},
		{
			Name:    `DecimalType`,
			Pattern: `(?i)FLOAT|DOUBLE|DECIMAL|DEC`,
		},
		{
			Name:    `BoolType`,
			Pattern: `(?i)BOOLEAN|BOOL`,
		},
		{
			Name:    `TimeType`,
			Pattern: `(?i)DATE|DATETIME|TIMESTAMP|TIME|YEAR`,
		},
		{
			Name:    `Ident`,
			Pattern: `[a-zA-Z_][a-zA-Z0-9_]*`,
		},
		{
			Name:    `Decimal`,
			Pattern: `[-+]?\d*\.\d+([eE][-+]?\d+)?`,
		},
		{
			Name:    `Int`,
			Pattern: `[-+]?\d+([eE][-+]?\d+)?`,
		},
		{
			Name:    `String`,
			Pattern: `'[^']*'|"[^"]*"`,
		},
		{
			Name:    `Bool`,
			Pattern: `(?i)TRUE|FALSE`,
		},
		{
			Name:    `Operator`,
			Pattern: `<>|!=|<=|>=|[-+*/%,.()=<>]`,
		},
		{
			Name:    "Semicolon",
			Pattern: ";",
		},
		{
			Name:    "whitespace",
			Pattern: `\s+`,
		},
	})

	return participle.Lexer(sqlLexer)
}
