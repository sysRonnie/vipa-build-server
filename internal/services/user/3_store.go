package user

import "database/sql"

type Store struct {
	db *sql.DB
}

func NewUserStore(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

type UserStore interface {}