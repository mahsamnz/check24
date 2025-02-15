package interfaces

type Serializer interface {
	Serialize(data interface{}) ([]byte, error)
	GetFormat() string
}
