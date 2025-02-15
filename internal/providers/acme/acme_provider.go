package acme

import (
	"github.com/mahsamnz/check24/internal/interfaces"
	"github.com/mahsamnz/check24/internal/models"
)

type ACMEServiceProvider struct {
	Identifier string                        `json:"identifier" default:"ACME"`
	Data       TarificacionThirdPartyRequest `json:"data"`
	mapper     *AcmeMapper
	serializer interfaces.Serializer
}

func NewACMEServiceProvider(serializer interfaces.Serializer) *ACMEServiceProvider {
	return &ACMEServiceProvider{
		Identifier: "ACME",
		mapper:     NewAcmeMapper(),
		serializer: serializer,
	}
}

func (acmeProvider *ACMEServiceProvider) SetSerializer(serializer interfaces.Serializer) {
	acmeProvider.serializer = serializer
}

func (acmeProvider *ACMEServiceProvider) GetIdentifier() string {
	return acmeProvider.Identifier
}

func (acmeProvider *ACMEServiceProvider) MapData(data models.CarInsuranceRequest) error {
	mappedData, err := acmeProvider.mapper.MapToProvider(data)
	if err != nil {
		return err
	}
	acmeProvider.Data = mappedData
	return nil
}

func (acmeProvider *ACMEServiceProvider) SerializeData() ([]byte, error) {
	return acmeProvider.serializer.Serialize(acmeProvider.Data)
}
