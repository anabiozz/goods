package datastore

import (
	"database/sql"

	"github.com/anabiozz/goods/backend/models"
)

// PostgresDatastore ...
type PostgresDatastore struct {
	*sql.DB
}

// NewPostgresDatastore ...
func NewPostgresDatastore() (*PostgresDatastore, error) {
	connection, err := sql.Open("postgres", "")
	if err != nil {
		return nil, err
	}
	return &PostgresDatastore{
		DB: connection,
	}, nil
}

// CreateUser ...
func (p *PostgresDatastore) CreateUser(user *models.User) error {
	return nil
}

// SaveImage ..
func (p *PostgresDatastore) SaveImage(image *models.Image) error {
	return nil
}

// CloseDB ..
func (p *PostgresDatastore) CloseDB() {
	p.DB.Close()
}
