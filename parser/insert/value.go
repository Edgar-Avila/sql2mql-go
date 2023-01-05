package insert

import (
	"strings"

	"github.com/alecthomas/participle/v2"
)

type Boolean bool

func (b *Boolean) Capture(values []string) error {
	*b = Boolean(strings.EqualFold(values[0], "TRUE"))
	return nil
}

type Str string

func (s *Str) Capture(values []string) error {
	*s = Str(strings.Trim(values[0], "'\""))
	return nil
}

type Value interface{ GetVal() interface{} }

type Int struct {
	Value int `parser:"@Int"`
}

func (i Int) GetVal() interface{} { return i.Value }

type Decimal struct {
	Value float64 `parser:"@Decimal"`
}

func (d Decimal) GetVal() interface{} { return d.Value }

type String struct {
	Value Str `parser:"@String"`
}

func (s String) GetVal() interface{} { return s.Value }

type Bool struct {
	Value Boolean `parser:"@('TRUE'|'FALSE')"`
}

func (b Bool) GetVal() interface{} { return b.Value }

func ValueUnion() participle.Option {
	return participle.Union[Value](Bool{}, String{}, Int{}, Decimal{})
}
