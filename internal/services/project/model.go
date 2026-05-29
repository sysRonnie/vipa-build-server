package project

import "time"


type ProjectRow struct {
	ID            int       `json:"id"`
	CustomerName  string    `json:"customer_name"`
	Name          string    `json:"name"`
	StartDate     string    `json:"start_date"`
	EndDateEst    string    `json:"end_date_est"`
	EndDateActual string    `json:"end_date_actual"`
	Price         float64   `json:"price"`
	Budget 	      float64   `json:"budget"`
	Note 		string    `json:"note"`
	IsDeleted     bool      `json:"is_deleted"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type ProjectRowList struct {
	Projects []ProjectRow `json:"projects"`
}

type ProjectNameList struct {
	Names []string `json:"names"`
}

type VendorExpense struct {
	VendorName string `json:"vendor"`
	CostCategoryParent string `json:"cost_category_parent"`
	CostCategoryChild *string `json:"cost_category_child"`
	Amount float64 `json:"amount"`
}

type VendorExpenseList struct {
	VendorExpenseList []VendorExpense `json:"vendor_expense_list"`
}