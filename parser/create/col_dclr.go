package create

// Column declaration
type ColumnDclr struct {
	Name      string     `parser:"@Ident"`
	Type      ColType    `parser:"@@"`
	Modifiers []Modifier `parser:"@(PrimaryKey|NotNull|Unique)*"`
}
