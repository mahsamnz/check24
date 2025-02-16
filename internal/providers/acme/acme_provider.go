package acme

import (
	"github.com/mahsamnz/check24/internal/interfaces"
	"github.com/mahsamnz/check24/internal/models"
)

type ACMEServiceProvider struct {
	Identifier string `json:"identifier" default:"ACME"`
	data       TarificacionThirdPartyRequest
	mapper     *AcmeMapper
	serializer interfaces.Serializer
}

func NewACMEServiceProvider(serializer interfaces.Serializer) *ACMEServiceProvider {
	if serializer == nil {
		return nil
	}
	return &ACMEServiceProvider{
		Identifier: "ACME",
		mapper:     NewAcmeMapper(),
		serializer: serializer,
	}
}

func (acmeProvider *ACMEServiceProvider) GetIdentifier() string {
	return acmeProvider.Identifier
}

func (acmeProvider *ACMEServiceProvider) SetSerializer(serializer interfaces.Serializer) {
	acmeProvider.serializer = serializer
}

func (acmeProvider *ACMEServiceProvider) GetSerializer() interfaces.Serializer {
	return acmeProvider.serializer
}

func (acmeProvider *ACMEServiceProvider) mapData(data models.CarInsuranceRequest) error {
	mappedData, err := acmeProvider.mapper.MapToProvider(data)
	if err != nil {
		return err
	}
	acmeProvider.data = mappedData
	return nil
}

func (acmeProvider *ACMEServiceProvider) SerializeData(data models.CarInsuranceRequest) ([]byte, error) {
	err := acmeProvider.mapData(data)
	if err != nil {
		return nil, err
	}
	return acmeProvider.serializer.Serialize(acmeProvider.data)
}
