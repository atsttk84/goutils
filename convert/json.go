package text

import (
	"encoding/json"
	"log"
)

/*
  i: target struct or map

    cf. https://golang.org/pkg/encoding/json
*/
func ToString(i interface{}) string {
	j, err := json.Marshal(i)
	if err != nil {
		log.Printf("Convert String error. error: %s", err)
		return ""
	}
	return string(j)
}
