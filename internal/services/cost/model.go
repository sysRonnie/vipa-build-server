package cost

import "database/sql"

type CostRow struct {
	ID        int    `json:"id"`
	Parent    string `json:"parent"`
	Child     string `json:"child"`
	IsDeleted bool   `json:"is_deleted"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}


type CostRowList struct {
	Costs []CostRow `json:"costs"`
}

func (v *CostRow) Scan(scanner interface {
	Scan(dest ...any) error
}) error {

	return scanner.Scan(
		&v.ID,
		&v.Parent,
		&v.Child,
		&v.IsDeleted,
		&v.CreatedAt,
		&v.UpdatedAt,
	)
}


func ScanCostRows(rows *sql.Rows) ([]CostRow, error) {
	var costs []CostRow

	for rows.Next() {
		var cost CostRow

		err := cost.Scan(rows)
		if err != nil {
			return nil, err
		}

		costs = append(costs, cost)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	if len(costs) == 0 {
		return []CostRow{}, nil
	}

	return costs, nil
}

func ScanCostRow(scanner interface {
	Scan(dest ...any) error
}) (*CostRow, error) {

	var cost CostRow

	err := cost.Scan(scanner)
	if err != nil {
		return nil, err
	}

	return &cost, nil
}