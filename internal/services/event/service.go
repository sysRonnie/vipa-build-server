package event 


type Service struct {
	store EventStore
}

func NewEventService(store EventStore) *Service {
	return &Service{store: store}
}


type EventService interface {

}

