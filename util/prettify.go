package util

import (
	"encoding/json"
	"log"
)

// Devolver una estructura bien formateada para imprimir
func Prettify(obj interface{}) string {
	s, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(s)
}
