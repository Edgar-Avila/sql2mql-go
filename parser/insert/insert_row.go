package insert

import "participle-test/parser/common"

type Row struct {
	Values []common.Value `parser:"'(' @@ (',' @@)* ')'"`
}
