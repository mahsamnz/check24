package interfaces

import "github.com/mahsamnz/check24/internal/models"

type Mapper interface {
	MapToProvider(data models.CarInsuranceRequest) interface{}
}
