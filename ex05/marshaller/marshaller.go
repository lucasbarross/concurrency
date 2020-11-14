package marshaller

import (
	"encoding/json"
)

type Marshaller interface  {
	Marshal(interface{}) ([]byte, error)
	Unmarshal(data []byte, v interface{}) error
}

type JsonMarshaller struct {} 

func (m JsonMarshaller) Marshal(data interface{}) ([]byte, error) {
	return json.Marshal(data)
}

func (m JsonMarshaller) Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}