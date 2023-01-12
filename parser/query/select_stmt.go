package query

// Sentencia SELECT
type SelectStmt struct {
	// Parsear DISTINCT si lo hay o no
	Distinct bool `parser:"'SELECT' 'DISTINCT'?"`

	// Array de nombres de columnas
	// Puede ser solo un * (Seleccionar todo) o varios identificadores
	Cols []string `parser:"((@'*')|(@Ident (',' @Ident)*))"`

	// Nombre de la tabla (Identificador)
	From string `parser:"'FROM' @Ident"`

	// Sentencia Where
	Where *Where `parser:"@@?"`

	// Parsear nombre de columnas por las que agrupar (Identificadores)
	GroupBy []string `parser:"('GROUP' 'BY' @Ident (',' @Ident)*)?"`

	// Parsear nombre de tablas por las que ordenar (Parseado en SortSpec)
	OrderBy []SortSpec `parser:"('ORDER' 'BY' @@ (',' @@)*)?"`

	// Parsear limite de registros a retornar (INT)
	Limit int64 `parser:"('LIMIT' @Int)? ';'?"`
}

func (ss SelectStmt) StmtType() string { return "SELECT" }
