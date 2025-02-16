package acme

import (
	"testing"
	"time"

	"github.com/mahsamnz/check24/internal/models"
	"github.com/mahsamnz/check24/internal/serializers"
	"github.com/mahsamnz/check24/internal/utils"
	"github.com/stretchr/testify/assert"
)

func TestACMEServiceProvider_MapData(t *testing.T) {
	provider := NewACMEServiceProvider(serializers.NewXMLSerializer())

	input := models.CarInsuranceRequest{
		Holder:                       models.MainDriver,
		OccasionalDriver:             models.Yes,
		PrevInsurance_years:          utils.IntPtr(8),
		PrevInsurance_exists:         models.Yes,
		PrevInsurance_expirationDate: utils.StrPtr(time.Now().AddDate(0, 1, 0).Format(time.DateOnly)), // not expired insurance
	}

	err := provider.mapData(input)
	assert.NoError(t, err)

	// Verify the mapped data
	assert.Equal(t, Yes, provider.data.Datos.DatosGenerales.CondPpalEsTomador)
	assert.Equal(t, Yes, provider.data.Datos.DatosGenerales.ConductorUnico)
	assert.Equal(t, 8, provider.data.Datos.DatosGenerales.AnosSegAnte)
	assert.Equal(t, 1, provider.data.Datos.DatosGenerales.NroCondOca)
	assert.Equal(t, Yes, provider.data.Datos.DatosAseguradora.SeguroEnVigor)
}

func TestACMEServiceProvider_SerializeData(t *testing.T) {
	provider := NewACMEServiceProvider(serializers.NewXMLSerializer())

	input := models.CarInsuranceRequest{
		Holder:                       models.MainDriver,
		OccasionalDriver:             models.Yes,
		PrevInsurance_years:          utils.IntPtr(8),
		PrevInsurance_exists:         models.Yes,
		PrevInsurance_expirationDate: utils.StrPtr(time.Now().AddDate(0, 1, 0).Format(time.DateOnly)), // not expired insurance
	}

	xmlData, err := provider.SerializeData(input)
	assert.NoError(t, err)
	assert.Contains(t, string(xmlData), "<CondPpalEsTomador>S</CondPpalEsTomador>")
	assert.Contains(t, string(xmlData), "<ConductorUnico>S</ConductorUnico>")
	assert.Contains(t, string(xmlData), "<AnosSegAnte>8</AnosSegAnte>")
	assert.Contains(t, string(xmlData), "<NroCondOca>1</NroCondOca>")
	assert.Contains(t, string(xmlData), "<SeguroEnVigor>S</SeguroEnVigor>")
}
