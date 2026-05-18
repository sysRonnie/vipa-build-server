package event

import (
	"context"
	"database/sql"
	"go-tailwind-test/internal/util/advisor"
	"go-tailwind-test/internal/util/network"
	"strings"
)

type Store struct {
	db *sql.DB
}


func NewEventStore(db *sql.DB) *Store {
	return &Store{db: db}
}

type EventStore interface {
	QueryEventList(ctx context.Context) ([]EventRow, error)
	QueryEventListRecycled(ctx context.Context) ([]EventRow, error)
	QueryEventListNames(ctx context.Context) ([]string, error)
	QueryEventByID(ctx context.Context, id int) (*EventRow, error)
	InsertEvent(ctx context.Context, newEvent EventRow) error
	UpdateEvent(ctx context.Context, updatedEvent EventRow) error
	DeleteEvent(ctx context.Context, id int) error
}

func (s *Store) QueryEventListNames(ctx context.Context) ([]string, error) {
	advisor := advisor.FromContext(ctx)
	advisor.Log("Querying event list names from the database")
	advisor.Log(" ========== SQL QUERY START ==========")
	advisor.Log(baseEventListNamesQuery)
	advisor.Log(" ========== SQL QUERY END ==========")
	rows, err := s.db.QueryContext(ctx, baseEventListNamesQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var names []string
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, err
		}
		names = append(names, name)
	}
	advisor.Log("Successfully queried event list names from the database: "+strings.Join(names, ", ") + " names found")
	return names, nil
}

func (s *Store) DeleteEvent(ctx context.Context, id int) error {
	_, err := s.db.ExecContext(ctx, baseEventDelete, id)
	return err
}

func (s *Store) UpdateEvent(ctx context.Context, updatedEvent EventRow) (err error) {
	res, err := s.db.ExecContext(ctx, baseEventUpdate,
		updatedEvent.Parent,
		updatedEvent.Child,
		updatedEvent.ID,
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


func (s *Store) InsertEvent(
	ctx context.Context,
	newEvent EventRow,
) error {
	advisor := advisor.FromContext(ctx)
	_, err := s.db.ExecContext(
		ctx,
		baseEventInsert,
		newEvent.Parent,
		newEvent.Child,
	)

	if err != nil {
		if strings.Contains(
			err.Error(),
			"duplicate key value violates unique constraint",
		) {
			advisor.Log("Attempted to insert duplicate event with parent: " + newEvent.Parent + " and child: " + newEvent.Child)
			return network.ErrRecordExists
		}

		return err
	}

	return nil
}

func (s *Store) QueryEventList(ctx context.Context) ([]EventRow, error) {
	rows, err := s.db.QueryContext(ctx, buildEventListQuery())
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return ScanEventRows(rows)
}

func (s *Store) QueryEventListRecycled(ctx context.Context) ([]EventRow, error) {
	rows, err := s.db.QueryContext(ctx, buildEventListRecycledQuery())
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return ScanEventRows(rows)
}

func (s *Store) QueryEventByID(ctx context.Context, id int) (*EventRow, error) {
	row := s.db.QueryRowContext(ctx, baseEventByIDQuery, id)
	return ScanEventRow(row)
}




