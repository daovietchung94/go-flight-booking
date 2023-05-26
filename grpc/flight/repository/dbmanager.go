package repository

import (
	"go-training/database"
	"go-training/grpc/flight/models"

	"gorm.io/gorm"
)

type dbmanager struct {
	*gorm.DB
}

func NewDBManager() (FlightRepository, error) {
	db, err := database.OpenDBConnection("flight")

	if err != nil {
		return nil, err
	}

	db = db.Debug()

	err = db.AutoMigrate(
		&models.Flight{},
	)

	if err != nil {
		return nil, err
	}

	return &dbmanager{db}, nil
}
