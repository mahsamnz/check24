package tests

import (
	"testing"
	"time"

	"github.com/mahsamnz/check24/internal/factories"
	"github.com/mahsamnz/check24/internal/models"
	"github.com/mahsamnz/check24/internal/providers/acme"
	"github.com/mahsamnz/check24/internal/serializers"
	"github.com/mahsamnz/check24/internal/utils"
	"github.com/stretchr/testify/assert"
)

func TestMainFlow(t *testing.T) {
	// Initialize factory
	factory := factories.NewServiceProviderFactory()

	// Register ACME provider
	acmeService := acme.NewACMEServiceProvider(serializers.NewXMLSerializer())
	factory.RegisterService(acmeService)

	// Test input
	input := models.CarInsuranceRequest{
		Holder:                       models.MainDriver,
		OccasionalDriver:             models.Yes,
		PrevInsurance_years:          utils.IntPtr(8),
		PrevInsurance_exists:         models.Yes,
		PrevInsurance_expirationDate: utils.StrPtr(time.Now().AddDate(0, 1, 0).Format(time.DateOnly)), // not expired insurance
	}

	// Get provider
	provider, err := factory.GetProvider("ACME")
	assert.NoError(t, err)

	// Serialize data
	xmlData, err := provider.SerializeData(input)
	assert.NoError(t, err)
	assert.Contains(t, string(xmlData), "<CondPpalEsTomador>S</CondPpalEsTomador>")
	assert.Contains(t, string(xmlData), "<ConductorUnico>S</ConductorUnico>")
	assert.Contains(t, string(xmlData), "<AnosSegAnte>8</AnosSegAnte>")
	assert.Contains(t, string(xmlData), "<NroCondOca>1</NroCondOca>")
	assert.Contains(t, string(xmlData), "<SeguroEnVigor>S</SeguroEnVigor>")
}
