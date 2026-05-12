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
	InsertCustomer(ctx context.Context, newCustomer CustomerRow) error
	QueryCustomerList(ctx context.Context) ([]CustomerRow, error)
	QueryCustomerListRecycled(ctx context.Context) ([]CustomerRow, error)
	QueryCustomerById(ctx context.Context, id int) (*CustomerRow, error)
	UpdateCustomer(ctx context.Context, customer CustomerInsertRequest) error
	DeleteCustomer(ctx context.Context, id int) error
	CheckCustomerExists(ctx context.Context, firstName, lastName string) (bool, error)
	CheckCustomerExistsRecycled(ctx context.Context, firstName, lastName string) (bool, error)
	QueryCustomerNames(ctx context.Context) ([]string, error)
}

func (s *Store) QueryCustomerNames(ctx context.Context) ([]string, error) {
	rows, err := s.db.QueryContext(ctx, baseCustomerNamesQuery)
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

	return names, nil
}

func (s *Store) DeleteCustomer(ctx context.Context, id int) error {
	_, err := s.db.ExecContext(ctx, baseCustomerSoftDelete, id)
	return err
}

func (s *Store) CheckCustomerExistsRecycled(ctx context.Context, firstName, lastName string) (bool, error) {
	var exists bool
	err := s.db.QueryRowContext(ctx, baseCustomerExistsRecycledQuery, firstName, lastName).Scan(&exists)
	return exists, err
}

func (s *Store) CheckCustomerExists(ctx context.Context, firstName, lastName string) (bool, error) {
	var exists bool
	err := s.db.QueryRowContext(ctx, baseCustomerExistsQuery, firstName, lastName).Scan(&exists)
	return exists, err
}


func (s *Store) UpdateCustomer(ctx context.Context, customer CustomerInsertRequest) error {

	result, err := s.db.ExecContext(ctx, baseCustomerUpdate,
		customer.Customer.FirstName,
		customer.Customer.LastName,
		customer.Customer.Phone,
		customer.Customer.Email,
		customer.Customer.Comment,
		customer.Customer.AddressStreet,
		customer.Customer.AddressUnit,
		customer.Customer.AddressCity,
		customer.Customer.AddressState,
		customer.Customer.AddressZip,
	)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (s *Store) QueryCustomerById(ctx context.Context, id int) (*CustomerRow, error) {
	row := s.db.QueryRowContext(
		ctx,
		buildCustomerDetailQuery(),
		id,
	)

	customer, err := scanCustomerRow(row)
	if err != nil {
		return nil, err
	}

	return &customer, nil
}

func (s *Store) InsertCustomer(ctx context.Context, newCustomer CustomerRow) error {
	_, err := s.db.ExecContext(ctx, baseCustomerInsert,
		newCustomer.FirstName,
		newCustomer.LastName,
		newCustomer.Phone,
		newCustomer.Email,
		newCustomer.AddressStreet,
		newCustomer.AddressUnit,
		newCustomer.AddressCity,
		newCustomer.AddressState,
		newCustomer.AddressZip,
		newCustomer.Comment,
	)

	return err
}



func (s *Store) QueryCustomerList(ctx context.Context) ([]CustomerRow, error) {
	return s.queryCustomers(ctx, false)
}

func (s *Store) QueryCustomerListRecycled(ctx context.Context) ([]CustomerRow, error) {
	return s.queryCustomers(ctx, true)
}

func (s *Store) queryCustomers(ctx context.Context, recycled bool) ([]CustomerRow, error) {
	rows, err := s.db.QueryContext(
		ctx,
		buildCustomerListQuery(recycled),
		recycled,
	)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	customers := []CustomerRow{}

	for rows.Next() {
		customer, err := scanCustomerRow(rows)
		if err != nil {
			return nil, err
		}

		customers = append(customers, customer)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return customers, nil
}
type customerScanner interface {
	Scan(dest ...any) error
}

func scanCustomerRow(scanner customerScanner) (CustomerRow, error) {
	var customer CustomerRow

	err := scanner.Scan(
		&customer.ID,
		&customer.FirstName,
		&customer.LastName,
		&customer.Phone,
		&customer.Email,
		&customer.AddressStreet,
		&customer.AddressUnit,
		&customer.AddressCity,
		&customer.AddressState,
		&customer.AddressZip,
		&customer.Comment,
		&customer.IsDeleted,
		&customer.CreatedAt,
		&customer.UpdatedAt,
	)

	return customer, err
}