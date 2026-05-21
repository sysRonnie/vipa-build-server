package activity

type ACTIVITY_TYPE string 

const (
	ACTIVITY_TYPE_EVENT   ACTIVITY_TYPE = "event"
	ACTIVITY_TYPE_EXPENSE ACTIVITY_TYPE = "expense"
	ACTIVITY_TYPE_INCOME  ACTIVITY_TYPE = "income"
	ACTIVITY_TYPE_NOTE    ACTIVITY_TYPE = "note"
	ACTIVITY_TYPE_PHOTO   ACTIVITY_TYPE = "photo"
)

type ActivityRow struct {
	ID int `json:"id"`

	UserID string `json:"user_id"`

	ProjectID int `json:"project_id"`
	ProjectName *string `json:"project_name,omitempty"`

	ActivityType string `json:"activity_type"`
	ActivityTitle string `json:"activity_title"`
	ActivityBody *string `json:"activity_body,omitempty"`

	Amount *float64 `json:"amount,omitempty"`

	ActivityDate *string `json:"activity_date,omitempty"`

	EventCategoryID *int `json:"event_category_id,omitempty"`
	EventCategoryName *string `json:"event_category_name,omitempty"`

	CostCategoryID *int `json:"cost_category_id,omitempty"`
	CostCategoryName *string `json:"cost_category_name,omitempty"`

	VendorID *int `json:"vendor_id,omitempty"`
	VendorName *string `json:"vendor_name,omitempty"`

	PhotoURL *string `json:"photo_url,omitempty"`
	PhotoThumbnailURL *string `json:"photo_thumbnail_url,omitempty"`

	FlagIsCompleted bool `json:"flag_is_completed"`
	FlagIsDeleted bool `json:"flag_is_deleted"`

	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type ActivityRowList struct {
	Activities []ActivityRow `json:"activities"`
}

type ActivityDeleteRequest struct {
	ID int `json:"id" validate:"required"`
}

type ExpenseActivityRow struct {
	ID int `json:"id"`
	ProjectName string `json:"project_name" validate:"required"`
	VendorName string `json:"vendor_name" validate:"required"`
	CostCategoryName string `json:"cost_category_name" validate:"required"`
	ExpenseTitle string `json:"expense_title" validate:"required"`
	ExpenseDesc *string `json:"expense_desc"`
	Amount float64 `json:"amount" validate:"required"`
	ExpenseDate string `json:"expense_date" validate:"required"`
	PhotoURL *string `json:"photo_url"`
}