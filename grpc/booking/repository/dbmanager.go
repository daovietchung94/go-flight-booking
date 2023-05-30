package repository

import (
	"go-training/config"
	"go-training/database"
	"go-training/grpc/booking/models"
	"gorm.io/gorm"
)

type dbmanager struct {
	db *gorm.DB
}

func NewDBManager(config *config.Config) (BookingRepository, error) {
	db, err := database.OpenDBConnection(config, "booking")

	if err != nil {
		return nil, err
	}

	db = db.Debug()

	err = db.AutoMigrate(
		&models.Reservation{},
	)

	if err != nil {
		return nil, err
	}

	return &dbmanager{
		db: db,
	}, nil
}
