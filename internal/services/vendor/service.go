package vendor


type Service struct {
	store VendorStore
}

func NewVendorService(store VendorStore) *Service {
	return &Service{
		store: store,
	}
}

type VendorService interface {	
}