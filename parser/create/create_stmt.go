package create

type CreateStmt struct {
	Name string       `parser:"'CREATE' 'TABLE' @Ident"`
	Cols []ColumnDclr `parser:"'(' @@ (',' @@)* ')' ';'?"`
}

func (cs CreateStmt) StmtType() string { return "CREATE" }
