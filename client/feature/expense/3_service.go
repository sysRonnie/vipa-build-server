package expense

import (
	"fmt"
	"log"
)

type Service struct {
	store ExpenseStore	
}

func NewService(store ExpenseStore) *Service {
	return &Service{store: store}
}

type ExpenseService interface {
	ProcessNewRecordInsertion(viewType ExpenseViewType, data map[string][]string) error
}

func (s *Service) ProcessNewRecordInsertion(
    viewType ExpenseViewType,
    data map[string][]string,
) error {

    log.Println("ProcessNewRecordInsertion", viewType)
    log.Println("raw data:", data)

    switch viewType {

    case ViewTypeClient:
        log.Println("client branch")

        clients := parseClients(data)

        return s.store.InsertNewRecordClient(clients)

    default:
        return fmt.Errorf("unknown viewType: %s", viewType)
    }
}