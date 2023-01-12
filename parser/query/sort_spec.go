package query

type Direction string

// Direccion (Ascendente o descendente)
func (d *Direction) Capture(values []string) error {
	*d = "ASC"
	if len(values) > 0 && values[0] == "DESC" {
		*d = "DESC"
	}
	return nil
}

// Especificacion del orden que va en ORDER BY
type SortSpec struct {
	// Nombre de la columna (Identificador)
	Col string `parser:"@Ident"`

	// Es ascendente o descendente
	Dir Direction `parser:"@(('ASC'|'DESC')?)"`
}
