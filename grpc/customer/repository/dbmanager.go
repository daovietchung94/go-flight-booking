package repository

import (
	"go-training/database"
	"go-training/grpc/customer/models"

	"gorm.io/gorm"
)

type dbmanager struct {
	*gorm.DB
}

func NewDBManager() (CustomerRepository, error) {
	db, err := database.OpenDBConnection("customer")

	if err != nil {
		return nil, err
	}

	db = db.Debug()

	err = db.AutoMigrate(
		&models.Customer{},
	)

	if err != nil {
		return nil, err
	}

	return &dbmanager{db}, nil
}
