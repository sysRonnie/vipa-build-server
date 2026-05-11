package customer 


import ( 

)



type Service struct {
	store CustomerStore
}

func NewCustomerService(store CustomerStore) *Service {
	return &Service{
		store: store,
	}
}

type CustomerService interface {

}