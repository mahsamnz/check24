package acme

type AcmeBoolean string

const (
	Yes AcmeBoolean = "S"
	No  AcmeBoolean = "N"
)

type DatosGenerales struct {
	CondPpalEsTomador AcmeBoolean `json:"CondPpalEsTomador"`
	ConductorUnico    AcmeBoolean `json:"ConductorUnico"`
	FecCot            string      `json:"FecCot"`
	AnosSegAnte       int         `json:"AnosSegAnte"`
	NroCondOca        int         `json:"NroCondOca"`
}

type DatosAseguradora struct {
	SeguroEnVigor AcmeBoolean `json:"SeguroEnVigor"`
}

type Datos struct {
	DatosGenerales   DatosGenerales   `json:"DatosGenerales"`
	DatosAseguradora DatosAseguradora `json:"DatosAseguradora"`
}

type TarificacionThirdPartyRequest struct {
	Cotizacion float64 `json:"Cotizacion"`
	Datos      Datos   `json:"Datos"`
}
