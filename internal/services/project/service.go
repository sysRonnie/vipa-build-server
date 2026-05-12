package project 


type Service struct {
	store ProjectStore
}


func NewProjectService(store ProjectStore) *Service {
	return &Service{
		store: store,
	}
}

type ProjectService interface {
}