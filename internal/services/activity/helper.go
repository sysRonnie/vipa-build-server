package activity


import "strings"

func ValidateFormExpense(activity ActivityRow) error {
	if activity.ProjectName == nil || strings.TrimSpace(*activity.ProjectName) == "" {
		return ErrProjectNameRequired
	}

	if activity.VendorName == nil || strings.TrimSpace(*activity.VendorName) == "" {
		return ErrVendorNameRequired
	}

	if activity.CostCategoryName == nil || strings.TrimSpace(*activity.CostCategoryName) == "" {
		return ErrCostCategoryNameRequired
	}

	if activity.Amount == nil || *activity.Amount <= 0 {
		return ErrAmountInvalid
	}

	if activity.ActivityDate == nil || strings.TrimSpace(*activity.ActivityDate) == "" {
		return ErrDateRequired
	}

	return nil
}

func ValidateFormEvent(activity ActivityRow) error {
	if activity.ProjectName == nil || strings.TrimSpace(*activity.ProjectName) == "" {
		return ErrProjectNameRequired
	}

	if activity.EventCategoryName == nil || strings.TrimSpace(*activity.EventCategoryName) == "" {
		return ErrEventTypeRequired
	}

	if strings.TrimSpace(activity.ActivityTitle) == "" {
		return ErrEventTitleRequired
	}

	return nil
}

func (a ActivityRow) ValidateForCreate() error {
	switch a.ActivityType {
	case string(ACTIVITY_TYPE_EVENT):
		return ValidateFormEvent(a)

	case string(ACTIVITY_TYPE_EXPENSE):
		return ValidateFormExpense(a)

	case string(ACTIVITY_TYPE_INCOME):
		return ErrNotFound

	default:
		return ErrInvalidActivityType
	}
}