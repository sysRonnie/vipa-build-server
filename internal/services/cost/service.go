package cost 

import (

)
type Service struct {
	store CostStore
}

func NewCostService(store CostStore) *Service {
	return &Service{store: store}
}

type CostService interface {

}

