package cost

import (
	"context"
	"database/sql"
	"go-tailwind-test/internal/util/network"
	"strings"
)

type Store struct {
	db *sql.DB
}



func NewCostStore(db *sql.DB) *Store {
	return &Store{db: db}
}

type CostStore interface {
	QueryCostList(ctx context.Context) ([]CostRow, error)
	QueryCostListRecycled(ctx context.Context) ([]CostRow, error)
	QueryCostByID(ctx context.Context, id int) (*CostRow, error)
	InsertCost(ctx context.Context, newCost CostRow) error
	UpdateCost(ctx context.Context, updatedCost CostRow) error
	DeleteCost(ctx context.Context, id int) error
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


func (s *Store) InsertCost(
	ctx context.Context,
	newCost CostRow,
) error {

	_, err := s.db.ExecContext(
		ctx,
		baseCostInsert,
		newCost.Parent,
		newCost.Child,
	)

	if err != nil {
		if strings.Contains(
			err.Error(),
			"duplicate key value violates unique constraint",
		) {
			return network.ErrRecordExists
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


