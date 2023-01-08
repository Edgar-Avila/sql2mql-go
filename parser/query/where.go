package query

import "participle-test/parser/common"

type Comparison struct {
	Col string       `parser:"@Ident"`
	Op  string       `parser:"@('='|'<'|'>'|'<='|'>=')"`
	Lit common.Value `parser:"@@"`
}

type BoolFactor struct {
	Not bool       `parser:"@'NOT'?"`
	Cmp Comparison `parser:"@@"`
}

type BoolTerm struct {
	Factors []BoolFactor `parser:"@@ ('AND' @@)*"`
}

type Where struct {
	Terms []BoolTerm `parser:"'WHERE' @@ ('OR' @@)*"`
}
