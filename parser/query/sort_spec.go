package query

type Direction string

func (d *Direction) Capture(values []string) error {
	*d = "ASC"
	if len(values) > 0 && values[0] == "DESC" {
		*d = "DESC"
	}
	return nil
}

// Sort specification
type SortSpec struct {
	Col string    `parser:"@Ident"`
	Dir Direction `parser:"@(('ASC'|'DESC')?)"`
}
