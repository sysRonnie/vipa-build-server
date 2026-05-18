package event

import "database/sql"

type EventRow struct {
	ID        int    `json:"id"`
	Parent    string `json:"parent"`
	Child     string `json:"child"`
	IsDeleted bool   `json:"is_deleted"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}


type EventRowList struct {
	Events []EventRow `json:"events"`
}

func (v *EventRow) Scan(scanner interface {
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

type EventNameList struct {
	Names []string `json:"names"`
}


func ScanEventRows(rows *sql.Rows) ([]EventRow, error) {
	var costs []EventRow

	for rows.Next() {
		var cost EventRow

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
		return []EventRow{}, nil
	}

	return costs, nil
}

func ScanEventRow(scanner interface {
	Scan(dest ...any) error
}) (*EventRow, error) {

	var cost EventRow

	err := cost.Scan(scanner)
	if err != nil {
		return nil, err
	}

	return &cost, nil
}