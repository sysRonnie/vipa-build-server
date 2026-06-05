package activity

import (
	"context"
	"database/sql"
	"errors"
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
	QueryActivityListNotes(ctx context.Context) ([]ActivityRow, error)
	QueryActivityListByID(ctx context.Context, id int) (ActivityRowList, error)
	QueryActivityListRecycled(ctx context.Context) ([]ActivityRow, error)
	QueryActivityById(ctx context.Context, id int) (*ActivityRow, error)
	QueryActivityByIdExpense(ctx context.Context, id int) (*ExpenseActivityRow, error)
	QueryActivityType(ctx context.Context, id int) (ACTIVITY_TYPE, error)
	DeleteActivitySoft(ctx context.Context, id int) error
	DeleteActivityErase(ctx context.Context, id int) error
	InsertActivityExpense(ctx context.Context, email string, newActivity ExpenseActivityRow) error
	UpdateActivityExpense(ctx context.Context, email string, updatedActivity ExpenseActivityRow) error
	InsertActivity(ctx context.Context, email string, newActivity ActivityRow) error
	UpdateActivity(ctx context.Context, email string, updatedActivity ActivityRow) (*ActivityRow, error)

	QueryProjectList(ctx context.Context) ([]string, error)
	QueryVendorList(ctx context.Context) ([]string, error)
	QueryEventCategoryList(ctx context.Context) ([]string, error)
	QueryCostCategoryList(ctx context.Context) ([]string, error)
	QueryIncomeCategoryList(ctx context.Context) ([]string, error)
}

func (s *Store) QueryActivityListByID(ctx context.Context, id int) (ActivityRowList, error) {
	advisor := advisor.FromContext(ctx)
	advisor.Log("store_attached_query_activity_list_by_id" + strconv.Itoa(id))

	rows, err := s.db.QueryContext(ctx, baseActivityListByProjectIdQuery, id)
	if err != nil {
		advisor.Error("failed to query activity list by id", err)
		return ActivityRowList{}, err
	}
	defer rows.Close()

	activities := make([]ActivityRow, 0)

	for rows.Next() {
		var activity ActivityRow

		err := rows.Scan(
			&activity.ID,
			&activity.UserID,
			&activity.ProjectID,
			&activity.ProjectName,
			&activity.CustomerID,
			&activity.CustomerName,
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
			&activity.IncomeCategoryID,
			&activity.IncomeCategoryName,
			&activity.PhotoURL,
			&activity.PhotoThumbnailURL,
			&activity.FlagIsCompleted,
			&activity.FlagIsDeleted,
			&activity.CreatedAt,
			&activity.UpdatedAt,
		)
		if err != nil {
			advisor.Error("failed to scan activity list by project id", err)
			return ActivityRowList{}, err
		}

		activities = append(activities, activity)
	}

	if err := rows.Err(); err != nil {
		advisor.Error("failed during activity list rows iteration", err)
		return ActivityRowList{}, err
	}

	return ActivityRowList{
		Activities: activities,
	}, nil
}

func (s *Store) QueryIncomeCategoryList(ctx context.Context) ([]string, error) {
	rows, err := s.db.QueryContext(ctx, baseIncomeListQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	var incomeCategoryNames []string
	for rows.Next() {
		var name string
		err := rows.Scan(&name)
		if err != nil {
			return nil, err
		}
		incomeCategoryNames = append(incomeCategoryNames, name)
	}
	
	if err := rows.Err(); err != nil {
		return nil, err
	}
	if len(incomeCategoryNames) == 0 {
		return []string{}, nil
	}
	return incomeCategoryNames, nil
}

func (s *Store) QueryCostCategoryList(ctx context.Context) ([]string, error) {
	rows, err := s.db.QueryContext(ctx, baseCostListNamesQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	var costCategoryNames []string
	for rows.Next() {
		var name string
		err := rows.Scan(&name)
		if err != nil {
			return nil, err
		}
		costCategoryNames = append(costCategoryNames, name)
	}
	
	if err := rows.Err(); err != nil {
		return nil, err
	}
	if len(costCategoryNames) == 0 {
		return []string{}, nil
	}
	return costCategoryNames, nil
}

func (s *Store) QueryEventCategoryList(ctx context.Context) ([]string, error) {
	rows, err := s.db.QueryContext(ctx, baseEventListNamesQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	var eventCategoryNames []string
	for rows.Next() {
		var name string
		err := rows.Scan(&name)
		if err != nil {
			return nil, err
		}
		eventCategoryNames = append(eventCategoryNames, name)
	}
	
	if err := rows.Err(); err != nil {
		return nil, err
	}
	if len(eventCategoryNames) == 0 {
		return []string{}, nil
	}
	return eventCategoryNames, nil
}

func (s *Store) QueryVendorList(ctx context.Context) ([]string, error) {
	rows, err := s.db.QueryContext(ctx, baseVendorListNamesQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	var vendorNames []string
	for rows.Next() {
		var name string
		err := rows.Scan(&name)
		if err != nil {
			return nil, err
		}
		vendorNames = append(vendorNames, name)
	}
	
	if err := rows.Err(); err != nil {
		return nil, err
	}
	if len(vendorNames) == 0 {
		return []string{}, nil
	}
	return vendorNames, nil
}

func (s *Store) QueryProjectList(ctx context.Context) ([]string, error) {
	rows, err := s.db.QueryContext(ctx, baseProjectListNamesQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	var projectNames []string
	for rows.Next() {
		var name string
		err := rows.Scan(&name)
		if err != nil {
			return nil, err
		}
		projectNames = append(projectNames, name)
	}
	
	if err := rows.Err(); err != nil {
		return nil, err
	}
	if len(projectNames) == 0 {
		return []string{}, nil
	}
	return projectNames, nil
}

func (s *Store) UpdateActivity(ctx context.Context, email string, updatedActivity ActivityRow) (*ActivityRow, error) {
	advisor := advisor.FromContext(ctx)
	advisor.Log("store_attached_update_activity_" + strconv.Itoa(updatedActivity.ID))

	var row ActivityRow

	err := s.db.QueryRowContext(
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
		updatedActivity.IncomeCategoryName,
		updatedActivity.PhotoURL,
		updatedActivity.FlagIsCompleted,
		updatedActivity.ID,
	).Scan(
		&row.ID,
		&row.UserID,
		&row.ProjectID,
		&row.ProjectName,
		&row.ActivityType,
		&row.ActivityTitle,
		&row.ActivityBody,
		&row.Amount,
		&row.ActivityDate,
		&row.EventCategoryID,
		&row.EventCategoryName,
		&row.CostCategoryID,
		&row.CostCategoryName,
		&row.VendorID,
		&row.VendorName,
		&row.IncomeCategoryID,
		&row.IncomeCategoryName,
		&row.PhotoURL,
		&row.PhotoThumbnailURL,
		&row.FlagIsCompleted,
		&row.FlagIsDeleted,
		&row.CreatedAt,
		&row.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			advisor.Error("no rows affected when updating activity", err)
			return nil, err
		}

		advisor.Error("failed to update activity", err)
		return nil, err
	}

	return &row, nil
}

func (s *Store) InsertActivity(ctx context.Context, email string, activity ActivityRow) error {
	advisor := advisor.FromContext(ctx)
	advisor.Log("store_insert_activity")

	advisor.Log("activity_date: " + *activity.ActivityDate) // Log the activity date for debugging

	_, err := s.db.ExecContext(
		ctx,
		baseActivityInsertNewActivity,
		email,
		activity.ProjectName,
		activity.ActivityType,
		activity.ActivityTitle,
		activity.ActivityBody,
		activity.ActivityDate,
		activity.Amount,
		activity.EventCategoryName,
		activity.CostCategoryName,
		activity.VendorName,
		activity.IncomeCategoryName,
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
			&activity.CustomerID,
			&activity.CustomerName,
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
			&activity.IncomeCategoryID,
			&activity.IncomeCategoryName,
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
	advisor.Log("store_attached_query_activity_list_recycled")
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
			&activity.CustomerID,
			&activity.CustomerName,
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
			&activity.CustomerID,
			&activity.CustomerName,
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

			&activity.IncomeCategoryID,
			&activity.IncomeCategoryName,

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
func (s *Store) QueryActivityListNotes(ctx context.Context) ([]ActivityRow, error) {
	advisor := advisor.FromContext(ctx)
	advisor.Log("store_attached_query_activity_list_notes")
	rows, err := s.db.QueryContext(ctx, baseActivityListQueryNotes)
	if err != nil {
		advisor.Error("failed to query activity list notes", err)
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

			&activity.CustomerID,
			&activity.CustomerName,

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

			&activity.IncomeCategoryID,
			&activity.IncomeCategoryName,

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