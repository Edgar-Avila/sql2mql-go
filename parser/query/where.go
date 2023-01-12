package query

import "participle-test/parser/common"

// Comparacion de una columna con un valor
type Comparison struct {
	// Nombre de columna (Identificador)
	Col string `parser:"@Ident"`

	// Operador
	Op string `parser:"@('='|'<'|'>'|'<='|'>=')"`

	// Valor
	Lit common.Value `parser:"@@"`
}

// Factor de una expresion booleana
type BoolFactor struct {
	// Not (Opcional)
	Not bool `parser:"@'NOT'?"`

	// Comparacion
	Cmp Comparison `parser:"@@"`
}

// Termino booleano
type BoolTerm struct {
	// Factores (Parseado en estructura BoolFactor)
	Factors []BoolFactor `parser:"@@ ('AND' @@)*"`
}

// Sentencia WHERE
type Where struct {
	// Terminos (Parseado en estructura BoolTerm)
	Terms []BoolTerm `parser:"'WHERE' @@ ('OR' @@)*"`
}
