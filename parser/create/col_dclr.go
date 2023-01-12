package create

// Declaracion de una columna
type ColumnDclr struct {
	// Nombre (Identificador)
	Name string `parser:"@Ident"`

	// Tipo (Se parsea en la estructura ColType)
	Type ColType `parser:"@@"`

	// Modificadores, (Se parsean en Modifier)
	// (Tiene varias opciones)
	Modifiers []Modifier `parser:"@(PrimaryKey|NotNull|Unique)*"`
}
