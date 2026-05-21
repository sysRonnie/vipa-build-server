package vendor

import "database/sql"

type VendorRow struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	PrimaryContact string `json:"primary_contact"`
	Phone          string `json:"phone"`
	Email          string `json:"email"`
	AddressStreet  string `json:"address_street"`
	AddressUnit    string `json:"address_unit"`
	AddressCity    string `json:"address_city"`
	AddressState   string `json:"address_state"`
	AddressZip     string `json:"address_zip"`
	Comment        string `json:"comment"`
	IsDeleted      bool   `json:"is_deleted"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
}

type VendorRowList struct {
	Vendors []VendorRow `json:"vendors"`
}

type VendorNameList struct {
	Names []string `json:"names"`
}

func (v *VendorRow) Scan(scanner interface {
	Scan(dest ...any) error
}) error {

	return scanner.Scan(
		&v.ID,
		&v.Name,
		&v.PrimaryContact,
		&v.Phone,
		&v.Email,
		&v.AddressStreet,
		&v.AddressUnit,
		&v.AddressCity,
		&v.AddressState,
		&v.AddressZip,
		&v.Comment,
		&v.IsDeleted,
		&v.CreatedAt,
		&v.UpdatedAt,
	)
}

func ScanVendorRows(rows *sql.Rows) ([]VendorRow, error) {
	var vendors []VendorRow

	for rows.Next() {
		var vendor VendorRow

		err := vendor.Scan(rows)
		if err != nil {
			return nil, err
		}

		vendors = append(vendors, vendor)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	if len(vendors) == 0 {
		return []VendorRow{}, nil
	}

	return vendors, nil
}

func ScanVendorRow(scanner interface {
	Scan(dest ...any) error
}) (*VendorRow, error) {

	var vendor VendorRow

	err := vendor.Scan(scanner)
	if err != nil {
		return nil, err
	}

	return &vendor, nil
}