package acme

import (
	"log"
	"sync"
	"time"

	"github.com/mahsamnz/check24/internal/models"
	"github.com/mahsamnz/check24/internal/utils"
)

// Mapper handles all mapping operations for ACME provider
type AcmeMapper struct{}

var (
	instance *AcmeMapper
	once     sync.Once
)

func NewAcmeMapper() *AcmeMapper {
	once.Do(func() {
		instance = &AcmeMapper{}
	})
	return instance
}

func (m *AcmeMapper) MapToProvider(data models.CarInsuranceRequest) (TarificacionThirdPartyRequest, error) {
	result := TarificacionThirdPartyRequest{
		Datos: Datos{
			DatosGenerales: DatosGenerales{
				CondPpalEsTomador: isMainDriver(data.Holder),
				ConductorUnico:    isOccasionalDriver(data.OccasionalDriver),
				AnosSegAnte:       *data.PrevInsurance_years,
				NroCondOca:        getNumberOfExtraDrivers(data.OccasionalDriver),
				FecCot:            utils.GetCurrentISODate(),
			},
			DatosAseguradora: DatosAseguradora{
				SeguroEnVigor: isPreviousInsuranceValid(data.PrevInsurance_exists, data.PrevInsurance_expirationDate),
			},
		},
	}

	return result, nil
}

func isMainDriver(holder models.Holder) AcmeBoolean {
	if holder == models.MainDriver {
		return Yes
	}
	return No
}

func isOccasionalDriver(occasionalDriver models.Boolean) AcmeBoolean {
	if occasionalDriver == models.Yes {
		return Yes
	}
	return No
}

func isPreviousInsuranceValid(preInsExists models.Boolean, preInsExpDate *string) AcmeBoolean {
	if preInsExists == models.No || preInsExpDate == nil {
		return No
	}

	// Check if previous insurance has expired
	expDate, err := time.Parse(time.DateOnly, *preInsExpDate)
	if err != nil {
		log.Printf("Error parsing expiration date: %v", err)
		return No
	}

	if expDate.After(time.Now()) {
		return Yes
	}

	return No
}

func getNumberOfExtraDrivers(occasionalDriver models.Boolean) int {
	if occasionalDriver == models.Yes {
		return 1
	}

	return 0
}
