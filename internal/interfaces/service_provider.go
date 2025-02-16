package interfaces

import "github.com/mahsamnz/check24/internal/models"

type ServiceProvider interface {
	GetIdentifier() string
	SetSerializer(serializer Serializer)
	GetSerializer() (serializer Serializer)
	SerializeData(data models.CarInsuranceRequest) ([]byte, error)
}
