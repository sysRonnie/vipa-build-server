package project

import (
	"context"
	"database/sql"
	"go-tailwind-test/internal/util/advisor"
	"log"
	"strconv"
)

type Store struct {
	db *sql.DB
}


func NewProjectStore(db *sql.DB) *Store {
	return &Store{db: db}
}

type ProjectStore interface {
	QueryProjectList(ctx context.Context) ([]ProjectRow, error)
	QueryProjectListRecycled(ctx context.Context) ([]ProjectRow, error)
	QueryProjectByID(ctx context.Context, id int) (*ProjectRow, error)
	InsertProject(ctx context.Context, newProject ProjectRow) error
	UpdateProject(ctx context.Context, updatedProject ProjectRow) (error)
	DeleteProject(ctx context.Context, id int) error
	QueryProjectListNames(ctx context.Context) ([]string, error)
	QueryProjectExpenseListByID(ctx context.Context, id int) (VendorExpenseList, error)
	QueryProjectIncomeListByID(ctx context.Context, id int) (VendorExpenseList, error)

	QueryProjectNameLatest(ctx context.Context) (string, error)
}

func (s *Store) QueryProjectNameLatest(ctx context.Context) (string, error) {
	row := s.db.QueryRowContext(ctx, baseProjectNameLatestQuery)
	var projectName string
	err := row.Scan(&projectName)
	if err != nil {
		return "", err
	}
	return projectName, nil
}
func (s *Store) QueryProjectIncomeListByID(ctx context.Context, id int) (VendorExpenseList, error) {
	rows, err := s.db.QueryContext(ctx, baseProjectIncomeListByIDQuery, id)
	if err != nil {
		return VendorExpenseList{}, err
	}
	defer rows.Close()
	
	incomeList := VendorExpenseList{
		VendorExpenseList: []VendorExpense{},
	}
	for rows.Next() {
		var income VendorExpense
		err := rows.Scan(
			&income.VendorName,
			&income.CostCategoryParent,
			&income.CostCategoryChild,
			&income.Amount,
		)
		if err != nil {
			return VendorExpenseList{}, err
		}
		incomeList.VendorExpenseList = append(incomeList.VendorExpenseList, income)
	}
	
	if err := rows.Err(); err != nil {
		log.Println("ERROR ERROR ERROR", err)
		return VendorExpenseList{}, err
	}
	
	return incomeList, nil
}

func (s *Store) QueryProjectExpenseListByID(ctx context.Context, id int) (VendorExpenseList, error) {
	rows, err := s.db.QueryContext(ctx, baseProjectExpenseListByIDQuery, id)
	if err != nil {
		return VendorExpenseList{}, err
	}
	defer rows.Close()
	
	expenseList := VendorExpenseList{
		VendorExpenseList: []VendorExpense{},
	}
	for rows.Next() {
		var expense VendorExpense
		err := rows.Scan(
			&expense.VendorName,
			&expense.CostCategoryParent,
			&expense.CostCategoryChild,
			&expense.Amount,
		)
		if err != nil {
			return VendorExpenseList{}, err
		}
		expenseList.VendorExpenseList = append(expenseList.VendorExpenseList, expense)
	}
	
	if err := rows.Err(); err != nil {
		log.Println("ERROR ERROR ERROR", err)
		return VendorExpenseList{}, err
	}
	
	return expenseList, nil
}

func (s *Store) QueryProjectListNames(ctx context.Context) ([]string, error) {
	rows, err := s.db.QueryContext(ctx, baseProjectListNamesQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	var names []string
	for rows.Next() {
		var name string
		err := rows.Scan(&name)
		if err != nil {
			return nil, err
		}
		names = append(names, name)
	}
	
	if err := rows.Err(); err != nil {
		return nil, err
	}
	if len(names) == 0 {
		return []string{}, nil
	}
	return names, nil
}


func (s *Store) DeleteProject(ctx context.Context, id int) error {
	_, err := s.db.ExecContext(ctx, baseProjectDelete, id)
	return err
}


func (s *Store) UpdateProject(ctx context.Context, updatedProject ProjectRow) (err error) {
	res, err := s.db.ExecContext(ctx, baseProjectUpdate,
		updatedProject.CustomerName,
		updatedProject.Name,
		nullableDate(updatedProject.StartDate),
		nullableDate(updatedProject.EndDateEst),
		nullableDate(updatedProject.EndDateActual),
		updatedProject.Price,
		updatedProject.Budget,
		updatedProject.Note,
		updatedProject.ID,
	)

	if err != nil {
		return err
	}
	advisor	 := advisor.FromContext(ctx)
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	advisor.Log("Executed update project query, checking rows affected" + strconv.FormatInt(rowsAffected, 10) + " rows affected")

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}


func (s *Store) QueryProjectByID(ctx context.Context, id int) (*ProjectRow, error) {
	row := s.db.QueryRowContext(ctx, baseProjectByIDQuery, id)
	
	var project ProjectRow
	err := row.Scan(
		&project.ID,
		&project.CustomerName,
		&project.Name,
		&project.StartDate,
		&project.EndDateEst,
		&project.EndDateActual,
		&project.Price,
		&project.Budget,
		&project.Note,
		&project.IsDeleted,
		&project.CreatedAt,
		&project.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &project, nil
}

func nullableDate(value string) any {
	if value == "" {
		return nil
	}

	return value
}
func (s *Store) InsertProject(ctx context.Context, newProject ProjectRow) error {
	_, err := s.db.ExecContext(ctx, baseProjectInsert,
		newProject.CustomerName,
		newProject.Name,
		nullableDate(newProject.StartDate),
		nullableDate(newProject.EndDateEst),
		nullableDate(newProject.EndDateActual),
		newProject.Price,
		newProject.Budget,
		newProject.Note,
	)
	return err
}



func (s *Store) QueryProjectList(ctx context.Context) ([]ProjectRow, error) {
	rows, err := s.db.QueryContext(ctx, buildProjectListQuery())
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	var projects []ProjectRow
	for rows.Next() {
		var project ProjectRow
		err := rows.Scan(
			&project.ID,
			&project.CustomerName,
			&project.Name,
			&project.StartDate,
			&project.EndDateEst,
			&project.EndDateActual,
			&project.Price,
			&project.Budget,
			&project.Note,
			&project.IsDeleted,
			&project.CreatedAt,
			&project.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		projects = append(projects, project)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	if len(projects) == 0 {
		return []ProjectRow{}, nil
	}

	return projects, nil	
}

func (s *Store) QueryProjectListRecycled(ctx context.Context) ([]ProjectRow, error) {
	rows, err := s.db.QueryContext(ctx, buildProjectListRecycledQuery())
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	var projects []ProjectRow
	for rows.Next() {
		var project ProjectRow
		err := rows.Scan(
			&project.ID,
			&project.CustomerName,
			&project.Name,
			&project.StartDate,
			&project.EndDateEst,
			&project.EndDateActual,
			&project.Price,
			&project.Budget,
			&project.Note,
			&project.IsDeleted,
			&project.CreatedAt,
			&project.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		projects = append(projects, project)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	
	if len(projects) == 0 {
		return []ProjectRow{}, nil
	}

	return projects, nil
}