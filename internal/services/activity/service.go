package activity

import (
	"context"
	"go-tailwind-test/internal/services/auth"
	"go-tailwind-test/internal/util/advisor"
	"strconv"
)


type Service struct {
	store ActivityStore
}


func NewActivityService(store ActivityStore) *Service {
	return &Service{store: store}
}


type ActivityService interface {
	GetActivityList(ctx context.Context) ([]ActivityRow, error)
	GetActivityListRecycled(ctx context.Context) ([]ActivityRow, error)
	GetActivityById(ctx context.Context, id int) (*ActivityRow, error)
	InsertActivityExpense(ctx context.Context, newActivity ExpenseActivityRow) error
	UpdateActivityExpense(ctx context.Context, updatedActivity ExpenseActivityRow) error
	DeleteActivitySoft(ctx context.Context, id int) error
	DeleteActivityErase(ctx context.Context, id int) error
	CreateActivity(ctx context.Context, newActivity ActivityRow) error
	UpdateActivity(ctx context.Context, updatedActivity ActivityRow) error
}

func (s *Service) UpdateActivity(ctx context.Context, updatedActivity ActivityRow) error {
	advisor := advisor.FromContext(ctx)
	advisor.Log("service_attached_update_activity" + strconv.Itoa(updatedActivity.ID))
	
	claims, err := auth.GetClaims(ctx)
	if err != nil {
		advisor.Error("failed to get claims", err)
		return err
	}

	err = s.store.UpdateActivity(ctx, claims.UserEmail, updatedActivity)
	if err != nil {
		advisor.Error("failed to update activity", err)
		return err
	}
	return nil
}


func (s *Service) CreateActivity(ctx context.Context, activity ActivityRow) error {
	advisor := advisor.FromContext(ctx)
	advisor.Log("service_create_activity")

	claims, err := auth.GetClaims(ctx)
	if err != nil {
		advisor.Error("failed to get claims", err)
		return err
	}

	if err := s.store.InsertActivity(ctx, claims.UserEmail, activity); err != nil {
		advisor.Error("failed to insert activity", err)
		return err
	}

	return nil
}



func (s *Service) DeleteActivityErase(ctx context.Context, id int) error {
	advisor := advisor.FromContext(ctx)
	advisor.Log("service_attached_delete_activity_erase" + strconv.Itoa(id))
	
	err := s.store.DeleteActivityErase(ctx, id)
	if err != nil {
		advisor.Error("failed to delete activity erase", err)
		return err
	}
	return nil
}

func (s *Service) DeleteActivitySoft(ctx context.Context, id int) error {
	advisor := advisor.FromContext(ctx)
	advisor.Log("service_attached_delete_activity_soft" + strconv.Itoa(id))
	
	err := s.store.DeleteActivitySoft(ctx, id)
	if err != nil {
		advisor.Error("failed to delete activity soft", err)
		return err
	}
	return nil
}

func (s *Service) UpdateActivityExpense(ctx context.Context, updatedActivity ExpenseActivityRow) error {
	
	advisor	:= advisor.FromContext(ctx)
	advisor.Log("service_attached_update_activity_expense")
	claims, err := auth.GetClaims(ctx)
	if err != nil {
		advisor.Error("failed to get claims from context", err)
		return err
	}
	err = s.store.UpdateActivityExpense(ctx, claims.UserEmail, updatedActivity)
	if err != nil {
		advisor.Error("failed to update activity expense", err)
		return err
	}
	return nil
}

func (s *Service) GetActivityById(ctx context.Context, id int) (*ActivityRow, error) {
	advisor := advisor.FromContext(ctx)
	advisor.Log("service_attached_get_activity_list_by_id" + strconv.Itoa(id))

	activityType, err := s.store.QueryActivityType(ctx, id)
	if err != nil {
		advisor.Error("failed to get activity type by id", err)
		return nil, err
	}

	switch activityType {
	case ACTIVITY_TYPE_EXPENSE:

	}
	
	activity, err := s.store.QueryActivityById(ctx, id)
	if err != nil {
		advisor.Error("failed to get activity list by id", err)
		return nil, err
	}
	return activity, nil
}

func (s *Service) InsertActivityExpense(ctx context.Context, newActivity ExpenseActivityRow) error {

	advisor	:= advisor.FromContext(ctx)
	advisor.Log("service_attached_insert_activity_expense")
	claims, err := auth.GetClaims(ctx)
	if err != nil {
		advisor.Error("failed to get claims from context", err)
		return err
	}
	err = s.store.InsertActivityExpense(ctx, claims.UserEmail, newActivity)
	if err != nil {
		advisor.Error("failed to insert activity expense", err)
		return err
	}
	return nil
}

func (s *Service) GetActivityList(ctx context.Context) ([]ActivityRow, error) {
	advisor := advisor.FromContext(ctx)
	advisor.Log("service_attached_get_activity_list")
	activities, err := s.store.QueryActivityList(ctx)
	if err != nil {
		advisor.Error("failed to get activity list", err)
		return nil, err
	}
	return activities, nil
}

func (s *Service) GetActivityListRecycled(ctx context.Context) ([]ActivityRow, error) {
	advisor := advisor.FromContext(ctx)
	advisor.Log("service_attached_get_activity_list_recycled")
	activities, err := s.store.QueryActivityListRecycled(ctx)
	if err != nil {
		advisor.Error("failed to get activity list recycled", err)
		return nil, err
	}
	return activities, nil
}