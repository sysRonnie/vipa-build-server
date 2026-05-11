package customer

import (
	"context"
	"database/sql"
)

type Store struct {
	db *sql.DB
}

func NewCustomerStore(db *sql.DB) *Store {
	return &Store{db: db}
}

type CustomerStore interface {
	InsertCustomer(ctx context.Context, name, phone, email, comment string) error
	QueryCustomerList(ctx context.Context) ([]CustomerRow, error)
}

func (s *Store) InsertCustomer(ctx context.Context, name, phone, email, comment string) error {
	_, err := s.db.ExecContext(ctx, `
		INSERT INTO master_customer_list (
			customer_name,
			customer_phone,
			customer_email,
			comment
		)
		VALUES ($1, $2, $3, $4)
	`, name, phone, email, comment)

	return err
}

func (s *Store) QueryCustomerList(ctx context.Context) ([]CustomerRow, error) {
	rows, err := s.db.QueryContext(ctx, `
		SELECT
			id,
			customer_name,
			customer_phone,
			customer_email,
			comment
		FROM master_customer_list
		WHERE flag_is_deleted = false
		ORDER BY created_at DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	customers := []CustomerRow{}

	for rows.Next() {
		var customer CustomerRow

		if err := rows.Scan(
			&customer.ID,
			&customer.Name,
			&customer.Phone,
			&customer.Email,
			&customer.Comment,
		); err != nil {
			return nil, err
		}

		customers = append(customers, customer)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return customers, nil
}
