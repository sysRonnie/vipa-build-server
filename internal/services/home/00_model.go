package home

import "time"

type HomeDashboard struct {
	Projects []HomeProjectCard `json:"projects"`
}

type HomeProjectCard struct {
	ProjectID int `json:"project_id"`
	ProjectName string `json:"project_name"`
	CustomerName *string `json:"customer_name,omitempty"`
	TotalEvents int `json:"total_events"`
	TotalExpenses float64 `json:"total_expenses"`
	TotalIncome float64 `json:"total_income"`
	NetTotal float64 `json:"net_total"`
	TotalEventsComplete int `json:"total_events_complete"`
	TotalEventsPercent int `json:"total_events_percent"`
	TotalProjectDays int `json:"total_project_days"`
	LastActivityAt *time.Time `json:"last_activity_at,omitempty"`
	LastActivityPretty *string `json:"last_activity_pretty,omitempty"`
}