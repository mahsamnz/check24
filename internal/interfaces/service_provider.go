package interfaces

import "github.com/mahsamnz/check24/internal/models"

type ServiceProvider interface {
	GetIdentifier() string
	MapData(data models.CarInsuranceRequest) error
	SetSerializer(serializer Serializer)
	SerializeData() ([]byte, error)
}
