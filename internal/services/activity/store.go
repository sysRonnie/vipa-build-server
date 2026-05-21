package activity

import (
	"context"
	"database/sql"
	"go-tailwind-test/internal/util/advisor"
	"strconv"
)


type Store struct {
	db *sql.DB
}

func NewActivityStore(db *sql.DB) *Store {
	return &Store{db: db}
}

type ActivityStore interface {
	QueryActivityList(ctx context.Context) ([]ActivityRow, error)
	QueryActivityListRecycled(ctx context.Context) ([]ActivityRow, error)
	QueryActivityById(ctx context.Context, id int) (*ActivityRow, error)
	QueryActivityByIdExpense(ctx context.Context, id int) (*ExpenseActivityRow, error)
	QueryActivityType(ctx context.Context, id int) (ACTIVITY_TYPE, error)
	DeleteActivitySoft(ctx context.Context, id int) error
	DeleteActivityErase(ctx context.Context, id int) error
	InsertActivityExpense(ctx context.Context, email string, newActivity ExpenseActivityRow) error
	UpdateActivityExpense(ctx context.Context, email string, updatedActivity ExpenseActivityRow) error
	InsertActivity(ctx context.Context, email string, newActivity ActivityRow) error
	UpdateActivity(ctx context.Context, email string, updatedActivity ActivityRow) error
}

func (s *Store) UpdateActivity(ctx context.Context, email string, updatedActivity ActivityRow) error {
	advisor := advisor.FromContext(ctx)
	advisor.Log("store_attached_update_activity" + strconv.Itoa(updatedActivity.ID))
	
	_, err := s.db.ExecContext(
		ctx,
		baseActivityUpdate,
		email,
		updatedActivity.ProjectName,
		updatedActivity.ActivityType,
		updatedActivity.ActivityTitle,
		updatedActivity.ActivityBody,
		updatedActivity.Amount,
		updatedActivity.ActivityDate,
		
		updatedActivity.EventCategoryName,
		updatedActivity.CostCategoryName,
		updatedActivity.VendorName,
		updatedActivity.PhotoURL,
		
		updatedActivity.ID,
	)
	
	if err != nil {
		advisor.Error("failed to update activity", err)
		return err
	}
	return nil
}

func (s *Store) InsertActivity(ctx context.Context, email string, activity ActivityRow) error {
	advisor := advisor.FromContext(ctx)
	advisor.Log("store_insert_activity")

	_, err := s.db.ExecContext(
		ctx,
		baseActivityInsertNewActivity,
		email,
		activity.ProjectName,
		activity.ActivityType,
		activity.ActivityTitle,
		activity.ActivityBody,
		activity.Amount,
		activity.ActivityDate,
		activity.EventCategoryName,
		activity.CostCategoryName,
		activity.VendorName,
		activity.PhotoURL,
	)

	if err != nil {
		advisor.Error("failed to insert activity", err)
		return err
	}

	return nil
}


func (s *Store) DeleteActivityErase(ctx context.Context, id int) error {
	advisor := advisor.FromContext(ctx)
	advisor.Log("store_attached_delete_activity_erase" + strconv.Itoa(id))
	_, err := s.db.ExecContext(ctx, baseActivityDeleteErase, id)
	if err != nil {
		advisor.Error("failed to delete activity erase", err)
		return err
	}
	return nil
}

func (s *Store) DeleteActivitySoft(ctx context.Context, id int) error {
	advisor := advisor.FromContext(ctx)
	advisor.Log("store_attached_delete_activity_soft" + strconv.Itoa(id))
	_, err := s.db.ExecContext(ctx, baseActivityDeleteSoft, id)
	if err != nil {
		advisor.Error("failed to delete activity soft", err)
		return err
	}
	return nil
}

func (s *Store) UpdateActivityExpense(ctx context.Context, email string, updatedActivity ExpenseActivityRow) error {
	advisor := advisor.FromContext(ctx)
	advisor.Log("store_attached_update_activity_expense")
	_, err := s.db.ExecContext(
		ctx,
		baseActivityUpdateExpense,
		email,
		updatedActivity.ProjectName,
		updatedActivity.ExpenseTitle,
		updatedActivity.ExpenseDesc,
		updatedActivity.Amount,
		updatedActivity.ExpenseDate,
		updatedActivity.CostCategoryName,
		updatedActivity.VendorName,
		updatedActivity.PhotoURL,
		updatedActivity.ID,
	)
	
	if err != nil {
		advisor.Error("failed to update activity expense", err)
		return err
	}
	return nil
}

func (s *Store) QueryActivityByIdExpense(ctx context.Context, id int) (*ExpenseActivityRow, error) {
	advisor := advisor.FromContext(ctx)
	advisor.Log("store_attached_query_activity_list_by_id_expense" + strconv.Itoa(id))
	rows, err := s.db.QueryContext(ctx, baseActivityByIdExpenseQuery, id)
	if err != nil {
		advisor.Error("failed to query activity list by id expense", err)
		return nil, err
	}
	defer rows.Close()
	
	if rows.Next() {
		var activity ExpenseActivityRow
		err := rows.Scan(
			&activity.ID,
			&activity.ProjectName,
			&activity.VendorName,
			&activity.CostCategoryName,
			&activity.ExpenseTitle,
			&activity.ExpenseDesc,
			&activity.Amount,
			&activity.ExpenseDate,
			&activity.PhotoURL,
		)

		if err != nil {
			return nil, err
		}
		return &activity, nil
	}
	
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return nil, sql.ErrNoRows
}


func (s *Store) QueryActivityType(ctx context.Context, id int) (ACTIVITY_TYPE, error) {
	row, err := s.db.QueryContext(ctx, baseActivityTypeQuery, id)
	if err != nil {
		return "", err
	}
	defer row.Close()

	if row.Next() {
		var activityType ACTIVITY_TYPE
		err := row.Scan(&activityType)
		if err != nil {
			return "", err
		}
		return activityType, nil
	}

	if err := row.Err(); err != nil {
		return "", err
	}
	return "", sql.ErrNoRows
}

func (s *Store) QueryActivityById(ctx context.Context, id int) (*ActivityRow, error) {
	advisor := advisor.FromContext(ctx)
	advisor.Log("store_attached_query_activity_list_by_id" + strconv.Itoa(id))
	rows, err := s.db.QueryContext(ctx, baseActivityListByIdQuery, id)
	if err != nil {
		advisor.Error("failed to query activity list by id", err)
		return nil, err
	}
	defer rows.Close()
	
	if rows.Next() {
		var activity ActivityRow
		err := rows.Scan(
			&activity.ID,
			&activity.UserID,
			&activity.ProjectID,
			&activity.ProjectName,
			&activity.ActivityType,
			&activity.ActivityTitle,
			&activity.ActivityBody,
			&activity.Amount,
			&activity.ActivityDate,
			
			&activity.EventCategoryID,
			&activity.EventCategoryName,
			
			&activity.CostCategoryID,
			&activity.CostCategoryName,
			
			&activity.VendorID,
			&activity.VendorName,
			&activity.PhotoURL,
			&activity.PhotoThumbnailURL,
			&activity.FlagIsCompleted,
			&activity.FlagIsDeleted,
			&activity.CreatedAt,
			&activity.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		return &activity, nil
	}
	
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return nil, sql.ErrNoRows
}

func (s *Store) InsertActivityExpense(
	ctx context.Context,
	email string,
	newActivity ExpenseActivityRow,
) error {
	advisor := advisor.FromContext(ctx)
	advisor.Log("store_attached_insert_activity_expense")

	_, err := s.db.ExecContext(
		ctx,
		baseActivityInsertExpense,
		email,
		newActivity.ProjectName,
		newActivity.ExpenseTitle,
		newActivity.ExpenseDesc,
		newActivity.Amount,
		newActivity.ExpenseDate,
		newActivity.CostCategoryName,
		newActivity.VendorName,
		newActivity.PhotoURL,
	)

	if err != nil {
		advisor.Error("failed to insert activity expense", err)
		return err
	}

	return nil
}

func (s *Store) QueryActivityListRecycled(ctx context.Context) ([]ActivityRow, error) {
	advisor := advisor.FromContext(ctx)
	advisor.Log("store_attached_query_activity_list")
	rows, err := s.db.QueryContext(ctx, BuildBaseActivityListQueryWithDeletedFilter(true))
	if err != nil {
		advisor.Error("failed to query activity list", err)
		return nil, err
	}
	defer rows.Close()
	
	var activities []ActivityRow
	for rows.Next() {
		var activity ActivityRow
		err := rows.Scan(
			&activity.ID,
			&activity.UserID,
			&activity.ProjectID,
			&activity.ProjectName,
			&activity.ActivityType,
			&activity.ActivityTitle,
			&activity.ActivityBody,
			&activity.Amount,
			&activity.ActivityDate,

			&activity.EventCategoryID,
			&activity.EventCategoryName,

			&activity.CostCategoryID,
			&activity.CostCategoryName,
			
			&activity.VendorID,
			&activity.VendorName,
			&activity.PhotoURL,
			&activity.PhotoThumbnailURL,
			&activity.FlagIsCompleted,
			&activity.FlagIsDeleted,
			&activity.CreatedAt,
			&activity.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		activities = append(activities, activity)
	}
	
	if err := rows.Err(); err != nil {
		return nil, err
	}
	if len(activities) == 0 {
		return []ActivityRow{}, nil
	}
	return activities, nil
}
func (s *Store) QueryActivityList(ctx context.Context) ([]ActivityRow, error) {
	advisor := advisor.FromContext(ctx)
	advisor.Log("store_attached_query_activity_list")
	rows, err := s.db.QueryContext(ctx, BuildBaseActivityListQueryWithDeletedFilter(false))
	if err != nil {
		advisor.Error("failed to query activity list", err)
		return nil, err
	}
	defer rows.Close()
	
	var activities []ActivityRow
	for rows.Next() {
		var activity ActivityRow
		err := rows.Scan(
			&activity.ID,
			&activity.UserID,
			&activity.ProjectID,
			&activity.ProjectName,
			&activity.ActivityType,
			&activity.ActivityTitle,
			&activity.ActivityBody,
			&activity.Amount,
			&activity.ActivityDate,

			&activity.EventCategoryID,
			&activity.EventCategoryName,

			&activity.CostCategoryID,
			&activity.CostCategoryName,
			
			&activity.VendorID,
			&activity.VendorName,
			&activity.PhotoURL,
			&activity.PhotoThumbnailURL,
			&activity.FlagIsCompleted,
			&activity.FlagIsDeleted,
			&activity.CreatedAt,
			&activity.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		activities = append(activities, activity)
	}
	
	if err := rows.Err(); err != nil {
		return nil, err
	}
	if len(activities) == 0 {
		return []ActivityRow{}, nil
	}
	return activities, nil
}