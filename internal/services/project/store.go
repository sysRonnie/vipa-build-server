package project

import (
	"context"
	"database/sql"
)

type Store struct {
	db *sql.DB
}


func NewProjectStore(db *sql.DB) *Store {
	return &Store{db: db}
}

type ProjectStore interface {
	QueryProjectList(ctx context.Context) ([]ProjectRow, error)
	InsertProject(ctx context.Context, newProject ProjectRow) error
}

func (s *Store) InsertProject(ctx context.Context, newProject ProjectRow) error {
	_, err := s.db.ExecContext(ctx, baseProjectInsert,
		newProject.CustomerName,
		newProject.Name,
		newProject.StartDate,
		newProject.EndDateEst,
		newProject.EndDateActual,
	)
	return err
}



func (s *Store) QueryProjectList(ctx context.Context) ([]ProjectRow, error) {
	rows, err := s.db.QueryContext(ctx, baseProjectListQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	var projects []ProjectRow
	for rows.Next() {
		var project ProjectRow
		err := rows.Scan(
			&project.ID,
			&project.CustomerName,
			&project.Name,
			&project.StartDate,
			&project.EndDateEst,
			&project.EndDateActual,
			&project.IsDeleted,
			&project.CreatedAt,
			&project.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		projects = append(projects, project)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	if len(projects) == 0 {
		return []ProjectRow{}, nil
	}

	return projects, nil	
}

