package expense

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
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
	RemoveRecordClient(id int) error
	QueryMasterClientList() ([]ExpenseFormModelClient, error)
	QueryMasterClientListRowById(id int) (ExpenseFormModelClient, error)
	CheckClientNameExists(tx *sql.Tx, name string) (bool, error)
}

func (s *Store) CheckClientNameExists(tx *sql.Tx, name string) (bool, error) {
	var exists bool
	err := tx.QueryRow(`
		SELECT EXISTS (
			SELECT 1
			FROM master_client_list
			WHERE lower(client_name) = lower($1)
			AND flag_is_deleted = false
		)
	`, name).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}


func (s *Store) RemoveRecordClient(id int) error {
	result, err := s.db.Exec(`
		UPDATE master_client_list 
		SET flag_is_deleted = true 
		WHERE id = $1
	`, id)
	if err != nil {
		log.Println("Error in RemoveRecordClient:", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no client found with id %d", id)
	}

	return nil
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
		exists, err := s.CheckClientNameExists(tx, client.Name)
		if err != nil {
			return err
		}
		if exists {
			return &DuplicateClientNameError{Name: client.Name}
		}

        _, err = tx.Exec(
            `INSERT INTO master_client_list (client_name, client_email, client_phone, client_address)
             VALUES ($1, $2, $3, $4)`,
            client.Name,
            client.Email,
            client.Phone,
            client.Address,
        )

		if err != nil {
			log.Println("client info", client.Name)
			log.Println("Error inserting client:", err)
		
			if strings.Contains(err.Error(), "master_client_list_client_name_key") {
				return &DuplicateClientNameError{Name: client.Name}
			}
		
			return err
		}
    }

    return tx.Commit()
}