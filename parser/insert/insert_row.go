package insert

type Row struct {
	Values []Value `parser:"'(' @@ (',' @@)* ')'"`
}
