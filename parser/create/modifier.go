package create

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
)

// Mdoificador
type Modifier int

// Posibles valores del modificador (Enum)
const (
	PrimaryKey Modifier = iota
	NotNull
	Unique
)

// Capturar el modificador, toma el string y
// returna el enum Modifier
func (cm *Modifier) Capture(values []string) error {
	text := regexp.MustCompile(`\s+`).ReplaceAllString(values[0], " ")
	if len(values) == 1 {
		if strings.EqualFold(text, "UNIQUE") {
			*cm = Unique
			return nil
		}
		if strings.EqualFold(text, "PRIMARY KEY") {
			*cm = PrimaryKey
			return nil
		}
		if strings.EqualFold(text, "NOT NULL") {
			*cm = NotNull
			return nil
		}

	}
	return fmt.Errorf("\"%s\" is not valid modifier", text)
}

// Convert to string
func (cm Modifier) String() string {
	if cm == PrimaryKey {
		return "PRIMARY KEY"
	} else if cm == NotNull {
		return "NOT NULL"
	} else if cm == Unique {
		return "UNIQUE"
	} else {
		return "UNKNOWN"
	}
}

// Serializar el Modificador como JSON
func (cm Modifier) MarshalJSON() ([]byte, error) {
	return json.Marshal(cm.String())
}
