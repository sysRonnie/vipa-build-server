package home

import (
	"context"
	"database/sql"
	"go-tailwind-test/internal/util/advisor"
)


type Store struct {
	db *sql.DB
}

func NewHomeStore(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}


type HomeStore interface {
	QueryHomeProjectCards(ctx context.Context) (HomeDashboard, error)
}

func (s *Store) QueryHomeProjectCards(ctx context.Context) (HomeDashboard, error) {
	advisor := advisor.FromContext(ctx)
	advisor.Log("Executing QueryHomeProjectCards")
	rows, err := s.db.QueryContext(ctx, baseHomeRead)
	if err != nil {
		advisor.Error("failed to execute home project cards query: ", err)
		return HomeDashboard{}, err
	}
	defer rows.Close()

	scannedRows, err := s.ScanHomeProjectCards(rows)
	if err != nil {
		advisor.Error("failed to scan home project cards: ", err)
		return HomeDashboard{}, err
	}


	

	return scannedRows, nil
}


