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
	QueryMasterClientList() ([]ExpenseFormModelClient, error)
	QueryMasterClientListRowById(id int) (ExpenseFormModelClient, error)
}

func (s *Store) QueryMasterClientListRowById(id int) (ExpenseFormModelClient, error) {
	row := s.db.QueryRow(`
		SELECT 
			id
			,client_name
			,client_email
			,client_phone
			,client_address
		FROM master_client_list
		WHERE id = $1
		AND flag_is_deleted = false
	`, id)

	var client ExpenseFormModelClient
	err := row.Scan(
		&client.Id,
		&client.Name,
		&client.Email,
		&client.Phone,
		&client.Address,
	)
	if err != nil {
		return ExpenseFormModelClient{}, err
	}

	return client, nil
}

func (s *Store) QueryMasterClientList() ([]ExpenseFormModelClient, error) {
	
	rows,err := s.db.Query(` 
		SELECT 
			id
			,client_name
			,client_email
			,client_phone
			,client_address
		FROM master_client_list
		WHERE 1=1
		AND flag_is_deleted = false
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// var listStringMap []map[string]string 
	// "message":"hello"
	// [{""}]
	var listClients []ExpenseFormModelClient
	for rows.Next() {
		var client ExpenseFormModelClient

		err := rows.Scan(
			&client.Id,
			&client.Name,
			&client.Email,
			&client.Phone,
			&client.Address,
		)
		if err != nil {
			return nil, err
		}
		
		listClients = append(listClients, client)
	}

	return listClients, nil

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