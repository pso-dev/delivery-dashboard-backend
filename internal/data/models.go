package data

import "database/sql"

type Repositories struct {
}

func NewRepositories(db *sql.DB) *Repositories {
	return &Repositories{}
}

type Repository interface {
	Insert(any) error
	Get(id int64) (*any, error)
	GetAll(...any) ([]*any, Metadata, error)
	Update(*any) error
	Delete(id int64) error
}
