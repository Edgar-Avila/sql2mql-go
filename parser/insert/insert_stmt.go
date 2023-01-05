package insert

type InsertStmt struct {
	Into string   `parser:"'INSERT' Into @Ident"`
	Cols []string `parser:"'(' @Ident (',' @Ident)* ')'"`
	Rows []Row    `parser:"'VALUES' @@ (',' @@)* ';'?"`
}

func (is InsertStmt) StmtType() string { return "INSERT" }
