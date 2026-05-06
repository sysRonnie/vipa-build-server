package user 

import (

)

type Service struct {
	store UserStore
}

func NewService(store UserStore) *Service {
	return &Service{store: store}
}

type UserService interface {
}