package activity

import "strings"

func ValidateFormIncome(activity ActivityRow) error {
	if activity.ProjectName == nil || strings.TrimSpace(*activity.ProjectName) == "" {
		return ErrProjectNameRequired
	}

	if activity.IncomeCategoryName == nil || strings.TrimSpace(*activity.IncomeCategoryName) == "" {
		return ErrIncomeCategoryRequired
	}

	if activity.Amount == nil || *activity.Amount <= 0 {
		return ErrAmountInvalid
	}

	if activity.ActivityDate == nil || strings.TrimSpace(*activity.ActivityDate) == "" {
		return ErrDateRequired
	}

	if strings.TrimSpace(activity.ActivityTitle) == "" {
		return ErrIncomeTitleRequired
	}

	return nil
}

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

func ValidateFormNote(activity ActivityRow) error {
	if activity.ProjectName == nil || strings.TrimSpace(*activity.ProjectName) == "" {
		return ErrProjectNameRequired
	}

	if strings.TrimSpace(activity.ActivityTitle) == "" && (activity.ActivityBody == nil || strings.TrimSpace(*activity.ActivityBody) == "") {
		return ErrNoteBodyRequired
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
		return ValidateFormIncome(a)

	case string(ACTIVITY_TYPE_NOTE):
		return ValidateFormNote(a)

	default:
		return ErrInvalidActivityType
	}
}