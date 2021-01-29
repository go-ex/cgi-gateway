package json

import (
	"encoding/json"
	"log"
)

func StructToString(obj interface{}) []byte {
	data, err := json.Marshal(obj)

	if err != nil {
		log.Println(err)
	}

	return data
}
