package marshaller

import (
	"encoding/json"
)

interface Marshaller {
	Marshal(interface{}) ([]byte, error)
	Unmarshal(data []byte, v interface{}) error
}

func Marshal(data interface{}) ([]byte, error) {
	return json.Marshal(data)
}

func Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}