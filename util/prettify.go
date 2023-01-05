package util

import (
	"encoding/json"
	"log"
)

func Prettify(obj interface{}) string {
	s, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(s)
}
