package serializers

import (
	"encoding/json"
)

type JSONSerializer struct{}

func NewJSONSerializer() *JSONSerializer {
	return &JSONSerializer{}
}

func (s *JSONSerializer) Serialize(data interface{}) ([]byte, error) {
	return json.Marshal(data)
}

func (s *JSONSerializer) GetFormat() string {
	return "json"
}
