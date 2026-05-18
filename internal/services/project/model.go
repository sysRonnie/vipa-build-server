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