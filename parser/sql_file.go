package parser

// Sql file
type SqlFile struct {
	Statements []Statement `parser:"@@+"`
}
