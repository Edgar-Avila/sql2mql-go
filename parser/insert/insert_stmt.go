package insert

// Sentencia insert
type InsertStmt struct {
	// Nombre de la tabla (Identificador)
	Into string `parser:"'INSERT' Into @Ident"`

	// Array de nombres de columnas (Identificadores)
	Cols []string `parser:"'(' @Ident (',' @Ident)* ')'"`

	// Array de filas (Parseados en Row)
	Rows []Row `parser:"'VALUES' @@ (',' @@)* ';'?"`
}

// Tipo de sentencia (INSERT)
func (is InsertStmt) StmtType() string { return "INSERT" }
