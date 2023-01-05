package drop

type DropStmt struct {
	Tables []string `parser:"'DROP' 'TEMPORARY'? 'TABLE' ('IF' 'EXISTS')? @Ident (',' @Ident)* ';'?"`
}

func (ds DropStmt) StmtType() string { return "DROP" }
