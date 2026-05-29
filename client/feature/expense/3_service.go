package expense

import (
	"fmt"
	"log"
	"strconv"
)

type Service struct {
	store ExpenseStore	
}

func NewService(store ExpenseStore) *Service {
	return &Service{store: store}
}

type ExpenseService interface {
	ProcessNewRecordInsertion(viewType ExpenseViewType, data map[string][]string) error
	ProcessExpenseTablePayload(viewType ExpenseViewType) (ExpenseTableModel, error)
	FetchExpenseTableRow(viewType ExpenseViewType, rowId int) (map[string]string, error)
	ProcessRemoveRecordRequest(viewType ExpenseViewType, rowId int) error
}





func (s *Service) ProcessRemoveRecordRequest(viewType ExpenseViewType, rowId int) error {
	log.Println("ProcessRemoveRecordRequest is working now")

	switch viewType {
	case ViewTypeClient:
		log.Println("client branch, removing record")
		return s.store.RemoveRecordClient(rowId)
	default:
		return fmt.Errorf("unknown viewType: %s", viewType)
	}
}

func (s *Service) FetchExpenseTableRow(viewType ExpenseViewType, rowId int) (map[string]string, error) {
	log.Println("FetchExpenseTableRow is working now")

	switch viewType {
	case ViewTypeClient:
		listClients, err := s.store.QueryMasterClientListRowById(rowId)
		if err != nil {
			return nil, err
		}
		row := map[string]string{
			"id":      strconv.Itoa(listClients.Id),
			"name":    listClients.Name,
			"email":   listClients.Email,
			"phone":   listClients.Phone,
			"address": listClients.Address,
		}
		

		return row, nil
	default:
		return nil, fmt.Errorf("unknown viewType: %s", viewType)
	}
}

func (s *Service) ProcessExpenseTablePayload(viewType ExpenseViewType) (ExpenseTableModel, error) {
	log.Println("ProcessExpenseTable is working now")

	switch viewType {
	case ViewTypeClient:
		listClients, err := s.store.QueryMasterClientList()
		if err != nil {
			return ExpenseTableModel{}, err
		}
		tableModel := BuildExpenseTableModel(listClients)
		return tableModel, nil
	default:
		return ExpenseTableModel{}, fmt.Errorf("unknown viewType: %s", viewType)
	}
}

func (s *Service) ProcessNewRecordInsertion(
    viewType ExpenseViewType,
    data map[string][]string,
) error {

  

    switch viewType {

    case ViewTypeClient:
        log.Println("client branch")

        clients := parseClients(data)

		log.Println("client records")
		log.Println(clients)

        return s.store.InsertNewRecordClient(clients)

    default:
        return fmt.Errorf("unknown viewType: %s", viewType)
    }
}