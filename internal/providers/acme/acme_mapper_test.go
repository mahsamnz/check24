package acme

import (
	"testing"
	"time"

	"github.com/mahsamnz/check24/internal/models"
	"github.com/mahsamnz/check24/internal/utils"
	"github.com/stretchr/testify/assert"
)

func TestAcmeMapper_MapToProvider(t *testing.T) {
	tests := []struct {
		name    string
		input   models.CarInsuranceRequest
		want    TarificacionThirdPartyRequest
		wantErr bool
	}{
		{
			name: "valid request with main driver and occasional driver",
			input: models.CarInsuranceRequest{
				Holder:                       models.MainDriver,
				OccasionalDriver:             models.Yes,
				PrevInsurance_years:          utils.IntPtr(8),
				PrevInsurance_exists:         models.Yes,
				PrevInsurance_expirationDate: utils.StrPtr(time.Now().AddDate(0, 1, 0).Format(time.DateOnly)), // not expired insurance
			},
			want: TarificacionThirdPartyRequest{
				Datos: Datos{
					DatosGenerales: DatosGenerales{
						CondPpalEsTomador: Yes,
						ConductorUnico:    Yes,
						AnosSegAnte:       8,
						NroCondOca:        1,
						FecCot:            utils.GetCurrentISODate(),
					},
					DatosAseguradora: DatosAseguradora{
						SeguroEnVigor: Yes,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "valid request with main driver and without occasional driver",
			input: models.CarInsuranceRequest{
				Holder:                       models.MainDriver,
				OccasionalDriver:             models.No,
				PrevInsurance_years:          utils.IntPtr(5),
				PrevInsurance_exists:         models.Yes,
				PrevInsurance_expirationDate: utils.StrPtr(time.Now().AddDate(0, 1, 0).Format(time.DateOnly)), // not expired insurance
			},
			want: TarificacionThirdPartyRequest{
				Datos: Datos{
					DatosGenerales: DatosGenerales{
						CondPpalEsTomador: Yes,
						ConductorUnico:    No,
						AnosSegAnte:       5,
						NroCondOca:        0,
						FecCot:            utils.GetCurrentISODate(),
					},
					DatosAseguradora: DatosAseguradora{
						SeguroEnVigor: Yes,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "valid request with main driver and with occasional driver with no previous insurance",
			input: models.CarInsuranceRequest{
				Holder:                       models.MainDriver,
				OccasionalDriver:             models.Yes,
				PrevInsurance_years:          utils.IntPtr(4),
				PrevInsurance_exists:         models.No,
				PrevInsurance_expirationDate: nil,
			},
			want: TarificacionThirdPartyRequest{
				Datos: Datos{
					DatosGenerales: DatosGenerales{
						CondPpalEsTomador: Yes,
						ConductorUnico:    Yes,
						AnosSegAnte:       4,
						NroCondOca:        1,
						FecCot:            utils.GetCurrentISODate(),
					},
					DatosAseguradora: DatosAseguradora{
						SeguroEnVigor: No,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "valid request with main driver and without occasional driver with expired previous insurance",
			input: models.CarInsuranceRequest{
				Holder:                       models.MainDriver,
				OccasionalDriver:             models.No,
				PrevInsurance_years:          utils.IntPtr(3),
				PrevInsurance_exists:         models.Yes,
				PrevInsurance_expirationDate: utils.StrPtr(time.Now().AddDate(-1, 0, 0).Format(time.DateOnly)), // expired insurance
			},
			want: TarificacionThirdPartyRequest{
				Datos: Datos{
					DatosGenerales: DatosGenerales{
						CondPpalEsTomador: Yes,
						ConductorUnico:    No,
						AnosSegAnte:       3,
						NroCondOca:        0,
						FecCot:            utils.GetCurrentISODate(),
					},
					DatosAseguradora: DatosAseguradora{
						SeguroEnVigor: No,
					},
				},
			},
			wantErr: false,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			mapper := NewAcmeMapper()
			got, err := mapper.MapToProvider(testCase.input)

			if testCase.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			// Skip FecCot comparison
			got.Datos.DatosGenerales.FecCot = testCase.want.Datos.DatosGenerales.FecCot
			assert.Equal(t, testCase.want, got)
		})
	}
}
