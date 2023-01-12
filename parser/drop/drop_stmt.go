package drop

// Sentencia DROP
type DropStmt struct {
	// Tablas, array de identificadores
	Tables []string `parser:"'DROP' 'TEMPORARY'? 'TABLE' ('IF' 'EXISTS')? @Ident (',' @Ident)* ';'?"`
}

// Tipo de sentencia (DROP)
func (ds DropStmt) StmtType() string { return "DROP" }
