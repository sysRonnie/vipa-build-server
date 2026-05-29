package home

import "context"


type Service struct {
	store HomeStore
}


func NewHomeService(store HomeStore) *Service {
	return &Service{
		store: store,
	}
}


type HomeService interface {
	ServiceHomeProjectCards(ctx context.Context) (HomeDashboard, error)
}

func (s *Service) ServiceHomeProjectCards(ctx context.Context) (HomeDashboard, error) {
	return s.store.QueryHomeProjectCards(ctx)
}
