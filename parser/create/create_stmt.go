package create

// Sentencia create
type CreateStmt struct {
	// Nombre de la tabla (Identificador)
	Name string `parser:"'CREATE' 'TABLE' @Ident"`

	// Array de columnas (Parseado en ColumnDclr)
	Cols []ColumnDclr `parser:"'(' @@ (',' @@)* ')' ';'?"`
}

// Para reconocer el tipo se sentencia (CREATE)
func (cs CreateStmt) StmtType() string { return "CREATE" }
