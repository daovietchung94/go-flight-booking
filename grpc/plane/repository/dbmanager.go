package repository

import (
	"go-training/config"
	"go-training/database"
	"go-training/grpc/plane/models"

	"gorm.io/gorm"
)

type dbmanager struct {
	*gorm.DB
}

func NewDBManager(config *config.Config) (PlaneRepository, error) {
	db, err := database.OpenDBConnection(config, "plane")

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
