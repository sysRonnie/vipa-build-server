package customer



type CustomerRow struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Phone string `json:"phone"`
	Email string `json:"email"`
	AddressStreet string `json:"address_street"`
	AddressCity string `json:"address_city"`
	AddressState string `json:"address_state"`
	AddressZip string `json:"address_zip"`
	Comment string `json:"comment"`
	IsDeleted bool `json:"is_deleted"`
}

type CustomerListResponse struct {
	Customers []CustomerRow `json:"customers"`
}

type CustomerInsertRequest struct {
	Name string `json:"name" validate:"required"`
	Phone string `json:"phone"`
	Email string `json:"email"` 
	Comment string `json:"comment"`
}