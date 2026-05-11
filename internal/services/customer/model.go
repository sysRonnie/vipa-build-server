package customer



type CustomerRow struct {
	ID int `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Phone string `json:"phone"`
	Email string `json:"email"`
	AddressStreet string `json:"address_street"`
	AddressUnit string `json:"address_unit"`
	AddressCity string `json:"address_city"`
	AddressState string `json:"address_state"`
	AddressZip string `json:"address_zip"`
	Comment string `json:"comment"`
	IsDeleted bool `json:"is_deleted"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}



type CustomerListResponse struct {
	Customers []CustomerRow `json:"customers"`
}

type CustomerInsertRequest struct {
	Customer CustomerRow `json:"customer"`
}

type DeleteCustomerRequest struct {
	ID int `json:"id" validate:"required"`
}
