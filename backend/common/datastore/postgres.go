package datastore

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/anabiozz/goods/backend/models"
)

const (
	dbHost     = "localhost"
	dbPort     = 5432
	dbUser     = "postgres"
	dbPassword = "postgres"
	dbName     = "art"
)

// PostgresDatastore ...
type PostgresDatastore struct {
	*sql.DB
}

// NewPostgresDatastore ...
func NewPostgresDatastore() (*PostgresDatastore, error) {
	dbinfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)
	connection, err := sql.Open("postgres", dbinfo)
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

// GetImagesByType ...
func (p *PostgresDatastore) GetImagesByType(imageType int) ([]models.Image, error) {
	query := fmt.Sprintf("SELECT uuid, name, materials, year, size, type, is_for_sale FROM main.images WHERE type = %d", imageType)
	rows, err := p.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var images []models.Image

	for rows.Next() {
		image := models.Image{}

		err = rows.Scan(&image.UUID, &image.Name, &image.Materials, &image.Year, &image.Size, &image.Type, &image.IsForSale)
		if err != nil {
			return nil, err
		}
		images = append(images, image)

	}
	return images, nil
}

// SaveImages ...
func (p *PostgresDatastore) SaveImages(images []*models.Image) error {

	for _, image := range images {
		tx, err := p.Begin()
		if err != nil {
			log.Print(err)
		}

		defer tx.Rollback()

		stmt, err := tx.Prepare("INSERT INTO main.images(uuid, name, materials, year, size, type, is_for_sale) VALUES ($1, $2, $3, $4, $5, $6, $7)")
		if err != nil {
			return err
		}

		defer stmt.Close()

		// default permissions  by groups_info
		_, err = stmt.Exec(image.UUID, image.Name, image.Materials, image.Year, image.Size, image.Type, image.IsForSale)
		if err != nil {
			return err
		}

		err = tx.Commit()
		if err != nil {
			return err
		}
	}
	return nil
}

// CloseDB ..
func (p *PostgresDatastore) CloseDB() {
	p.DB.Close()
}
