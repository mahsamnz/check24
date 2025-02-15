package models

type Holder string

const (
	MainDriver Holder = "CONDUCTOR_PRINCIPAL"
	// TODO define other constants based on the actiual options
	NonMainDriver Holder = "NO_CONDUCTOR_PRINCIPAL"
)

type Boolean string

const (
	Yes Boolean = "SI"
	No  Boolean = "NO"
)

type CarInsuranceRequest struct {
	Holder                       Holder  `json:"holder" validate:"required,oneof=CONDUCTOR_PRINCIPAL NO_CONDUCTOR_PRINCIPAL"`
	OccasionalDriver             Boolean `json:"occasionalDriver" validate:"required,oneof=SI NO"`
	PrevInsurance_years          *int    `json:"prevInsurance_years" validate:"required,min=0,max=50"`
	PrevInsurance_exists         Boolean `json:"prevInsurance_exists" validate:"required,oneof=SI NO"`
	PrevInsurance_expirationDate *string `json:"prevInsurance_expirationDate" validate:"omitempty,datetime=2006-01-02"`
}
