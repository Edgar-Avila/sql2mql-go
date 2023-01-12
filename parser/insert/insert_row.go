package insert

import "participle-test/parser/common"

// Parsear una fila
type Row struct {
	// Array de valores que se le da a la fila, debe haber 1 al menos
	Values []common.Value `parser:"'(' @@ (',' @@)* ')'"`
}
