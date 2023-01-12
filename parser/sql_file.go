package parser

// Archivo SQL (Consta de una o mas sentencias
// (CREATE, DROP, INSERT, SELECT))
type SqlFile struct {
	Statements []Statement `parser:"@@+"`
}
