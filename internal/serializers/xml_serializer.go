package serializers

import (
	"encoding/xml"
)

type XMLSerializer struct{}

func NewXMLSerializer() *XMLSerializer {
	return &XMLSerializer{}
}

func (s *XMLSerializer) Serialize(data interface{}) ([]byte, error) {
	return xml.Marshal(data)
}

func (s *XMLSerializer) GetFormat() string {
	return "xml"
}
