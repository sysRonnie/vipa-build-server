package vendor

import (
	"context"
	"database/sql"
)

type Store struct {
	db *sql.DB
}


func NewVendorStore(db *sql.DB) *Store {
	return &Store{db: db}
}

type VendorStore interface {
	QueryVendorList(ctx context.Context) ([]VendorRow, error)
	QueryVendorListNames(ctx context.Context) (VendorNameList, error)
	QueryVendorListRecycled(ctx context.Context) ([]VendorRow, error)
	QueryVendorByID(ctx context.Context, id int) (*VendorRow, error)
	InsertVendor(ctx context.Context, newVendor VendorRow) error
	UpdateVendor(ctx context.Context, updatedVendor VendorRow) error
	DeleteVendor(ctx context.Context, id int) error

	QueryVendorNameLatest(ctx context.Context) (string, error) // New method for latest vendor name
}

func (s *Store) QueryVendorNameLatest(ctx context.Context) (string, error) {
	row := s.db.QueryRowContext(ctx, baseVendorNameLatestQuery)
	
	var latest string
	err := row.Scan(&latest)
	if err != nil {
		return "", err
	}
	
	return latest, nil
}

func (s *Store) QueryVendorListNames(ctx context.Context) (VendorNameList, error) {
	rows, err := s.db.QueryContext(ctx, baseVendorListNamesQuery)
	if err != nil {
		return VendorNameList{}, err
	}
	defer rows.Close()

	names := VendorNameList{
		Names: []string{},
	}

	for rows.Next() {
		var name string

		if err := rows.Scan(&name); err != nil {
			return VendorNameList{}, err
		}

		names.Names = append(names.Names, name)
	}

	if err := rows.Err(); err != nil {
		return VendorNameList{}, err
	}

	return names, nil
}


func (s *Store) DeleteVendor(ctx context.Context, id int) error {
	_, err := s.db.ExecContext(ctx, baseVendorDelete, id)
	return err
}

func (s *Store) UpdateVendor(ctx context.Context, updatedVendor VendorRow) (err error) {
	res, err := s.db.ExecContext(ctx, baseVendorUpdate,
		updatedVendor.Name,
		updatedVendor.PrimaryContact,
		updatedVendor.Phone,
		updatedVendor.Email,
		updatedVendor.AddressStreet,
		updatedVendor.AddressUnit,
		updatedVendor.AddressCity,
		updatedVendor.AddressState,
		updatedVendor.AddressZip,
		updatedVendor.Comment,
		updatedVendor.ID,
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

func (s *Store) InsertVendor(ctx context.Context, newVendor VendorRow) error {
	_, err := s.db.ExecContext(ctx, baseVendorInsert,
		newVendor.Name,
		newVendor.PrimaryContact,
		newVendor.Phone,
		newVendor.Email,
		newVendor.AddressStreet,
		newVendor.AddressUnit,
		newVendor.AddressCity,
		newVendor.AddressState,
		newVendor.AddressZip,
		newVendor.Comment,
	)
	return err
}

func (s *Store) QueryVendorList(ctx context.Context) ([]VendorRow, error) {
	rows, err := s.db.QueryContext(ctx, buildVendorListQuery())
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return ScanVendorRows(rows)
}

func (s *Store) QueryVendorListRecycled(ctx context.Context) ([]VendorRow, error) {
	rows, err := s.db.QueryContext(ctx, buildVendorListRecycledQuery())
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return ScanVendorRows(rows)
}

func (s *Store) QueryVendorByID(ctx context.Context, id int) (*VendorRow, error) {
	row := s.db.QueryRowContext(ctx, baseVendorByIDQuery, id)
	return ScanVendorRow(row)
}


