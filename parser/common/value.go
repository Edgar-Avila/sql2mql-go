package common

// ****************************************************
// Capturar distintos valores int, bool, float, string
// ****************************************************

import (
	"strings"

	"github.com/alecthomas/participle/v2"
)

// Capturar un valor bool desde el string
type Boolean bool

func (b *Boolean) Capture(values []string) error {
	*b = Boolean(strings.EqualFold(values[0], "TRUE"))
	return nil
}

// Capturar un valor string
type Str string

func (s *Str) Capture(values []string) error {
	*s = Str(strings.Trim(values[0], "'\""))
	return nil
}

// Interfaz para los valores
type Value interface{ GetVal() interface{} }

// Capturar un entero
type Int struct {
	Value int `parser:"@Int"`
}

func (i Int) GetVal() interface{} { return i.Value }

// Capturar un decimal
type Decimal struct {
	Value float64 `parser:"@Decimal"`
}

func (d Decimal) GetVal() interface{} { return d.Value }

// Capturar un string
type String struct {
	Value Str `parser:"@String"`
}

func (s String) GetVal() interface{} { return s.Value }

// Capturar un bool
type Bool struct {
	Value Boolean `parser:"@('TRUE'|'FALSE')"`
}

func (b Bool) GetVal() interface{} { return b.Value }

// Union de todos los valores
func ValueUnion() participle.Option {
	return participle.Union[Value](Bool{}, String{}, Int{}, Decimal{})
}
