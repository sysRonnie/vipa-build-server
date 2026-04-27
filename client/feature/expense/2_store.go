package expense

import (
	"database/sql"
)

type Store struct {
	db *sql.DB
}

func NewExpenseStore(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

type ExpenseStore interface {
	InsertNewRecordClient(data []ExpenseFormModelClient) error
}

func (s *Store) InsertNewRecordClient(data []ExpenseFormModelClient) error {
    tx, err := s.db.Begin()
    if err != nil {
        return err
    }

    defer func() {
        if err != nil {
            tx.Rollback()
        }
    }()

    for _, client := range data {
        _, err = tx.Exec(
            `INSERT INTO master_client_list (client_name, client_email, client_phone, client_address)
             VALUES ($1, $2, $3, $4)`,
            client.Name,
            client.Email,
            client.Phone,
            client.Address,
        )

        if err != nil {
            return err
        }
    }

    return tx.Commit()
}