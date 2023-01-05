package query

type SelectStmt struct {
	Distinct bool       `parser:"'SELECT' 'DISTINCT'?"`
	Cols     []string   `parser:"((@'*')|(@Ident (',' @Ident)*))"`
	From     string     `parser:"'FROM' @Ident"`
	GroupBy  []string   `parser:"('GROUP' 'BY' @Ident (',' @Ident)*)?"`
	OrderBy  []SortSpec `parser:"('ORDER' 'BY' @@ (',' @@)*)?"`
	Limit    int64      `parser:"('LIMIT' @Int)? ';'?"`
}

func (ss SelectStmt) StmtType() string { return "SELECT" }
