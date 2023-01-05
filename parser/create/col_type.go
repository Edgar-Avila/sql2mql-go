package create

type ColType struct {
	Name   string `parser:"@(TextType|IntType|BoolType|DecimalType|TimeType)"`
	Params []int  `parser:"('(' @Int (',' @Int)* ')')?"`
}
