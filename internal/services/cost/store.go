package cost

import (
	"context"
	"database/sql"
	"errors"

	"github.com/lib/pq"
)

type Store struct {
	db *sql.DB
}



func NewCostStore(db *sql.DB) *Store {
	return &Store{db: db}
}

type CostStore interface {
	QueryCostList(ctx context.Context) ([]CostRow, error)
	QueryCostListNames(ctx context.Context) (CostNameList, error)
	QueryCostListRecycled(ctx context.Context) ([]CostRow, error)
	QueryCostByID(ctx context.Context, id int) (*CostRow, error)
	InsertCost(ctx context.Context, newCost CostRow) error
	UpdateCost(ctx context.Context, updatedCost CostRow) error
	DeleteCost(ctx context.Context, id int) error

	QueryCostNameLatest(ctx context.Context) (string, error)
}

func (s *Store) QueryCostNameLatest(ctx context.Context) (string, error) {
	row := s.db.QueryRowContext(ctx, baseCostNameLatestQuery)
	var costName string
	err := row.Scan(&costName)
	if err != nil {
		return "", err
	}
	return costName, nil
}

func (s *Store) QueryCostListNames(ctx context.Context) (CostNameList, error) {
	rows, err := s.db.QueryContext(ctx, baseCostListNamesQuery)
	if err != nil {
		return CostNameList{}, err
	}
	defer rows.Close()
	
	var names CostNameList
	for rows.Next() {
		var name string
		err := rows.Scan(&name)
		if err != nil {
			return CostNameList{}, err
		}
		names.Names = append(names.Names, name)
	}
	
	if err := rows.Err(); err != nil {
		return CostNameList{}, err
	}
	if len(names.Names) == 0 {
		return CostNameList{}, nil
	}
	return names, nil
}

func (s *Store) DeleteCost(ctx context.Context, id int) error {
	_, err := s.db.ExecContext(ctx, baseCostDelete, id)
	return err
}

func (s *Store) UpdateCost(ctx context.Context, updatedCost CostRow) (err error) {
	res, err := s.db.ExecContext(ctx, baseCostUpdate,
		updatedCost.Parent,
		updatedCost.Child,
		updatedCost.ID,
	)
	
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}
	return nil
}

func (s *Store) ExistsCostInRecycleBin(ctx context.Context,parent string,child string,) (bool, error) {
	row := s.db.QueryRowContext(
		ctx,
		baseCostExistsInRecycleBinQuery,
		parent,
		child,
	)

	var exists bool
	err := row.Scan(&exists)

	return exists, err
}

func (s *Store) InsertCost(ctx context.Context,newCost CostRow,) error {
	_, err := s.db.ExecContext(
		ctx,
		baseCostInsert,
		newCost.Parent,
		newCost.Child,
	)

	if err != nil {
		var pqErr *pq.Error

		if errors.As(err, &pqErr) && pqErr.Code == "23505" {
			exists, recycleErr := s.ExistsCostInRecycleBin(
				ctx,
				newCost.Parent,
				newCost.Child,
			)

			if recycleErr == nil && exists {
				return ErrDuplicateCostInRecycleBin
			}

			return ErrDuplicateCostNotDeleted
		}

		return err
	}

	return nil
}

func (s *Store) QueryCostList(ctx context.Context) ([]CostRow, error) {
	rows, err := s.db.QueryContext(ctx, buildCostListQuery())
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return ScanCostRows(rows)
}

func (s *Store) QueryCostListRecycled(ctx context.Context) ([]CostRow, error) {
	rows, err := s.db.QueryContext(ctx, buildCostListRecycledQuery())
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return ScanCostRows(rows)
}

func (s *Store) QueryCostByID(ctx context.Context, id int) (*CostRow, error) {
	row := s.db.QueryRowContext(ctx, baseCostByIDQuery, id)
	return ScanCostRow(row)
}


