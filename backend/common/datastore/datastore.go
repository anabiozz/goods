package datastore

import "github.com/anabiozz/goods/backend/models"

// Datastore ...
type Datastore interface {
	CreateUser(*models.User) error
	SaveImage(*models.Image) error
	CloseDB()
}

const (
	// POSTGRES ...
	POSTGRES = iota
)

// NewDatastore ...
func NewDatastore(datastoreType int) (Datastore, error) {
	switch datastoreType {
	case POSTGRES:
		return NewPostgresDatastore()
	}
	return nil, nil
}
