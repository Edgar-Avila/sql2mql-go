package create

// Tipo de datos para una columna
type ColType struct {
	// Nombre del tipo
	Name string `parser:"@(TextType|IntType|BoolType|DecimalType|TimeType)"`

	// Parametros que se le pasa
	Params []int `parser:"('(' @Int (',' @Int)* ')')?"`
}
