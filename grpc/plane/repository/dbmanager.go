package repository

import (
	"go-training/database"
	"go-training/grpc/plane/models"

	"gorm.io/gorm"
)

type dbmanager struct {
	*gorm.DB
}

func NewDBManager() (PlaneRepository, error) {
	db, err := database.OpenDBConnection("plane")

	if err != nil {
		return nil, err
	}

	db = db.Debug()

	err = db.AutoMigrate(
		&models.Plane{},
	)

	if err != nil {
		return nil, err
	}

	return &dbmanager{db}, nil
}
