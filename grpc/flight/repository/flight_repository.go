package repository

import (
	"context"
	"errors"
	"go-training/grpc/flight/models"

	"github.com/google/uuid"
)

type FlightRepository interface {
	CreateFlight(context context.Context, model *models.Flight) (*models.Flight, error)
	UpdateFlight(context context.Context, model *models.Flight) (*models.Flight, error)
	FlightDetails(context context.Context, id uuid.UUID) (*models.Flight, error)
}

func (conn *dbmanager) CreateFlight(context context.Context, model *models.Flight) (*models.Flight, error) {
	model.Id = uuid.New()
	if err := conn.Create(model).Error; err != nil {
		return nil, err
	}

	return model, nil
}

func (conn *dbmanager) UpdateFlight(context context.Context, model *models.Flight) (*models.Flight, error) {
	cs := []*models.Flight{}
	err := conn.Where(&models.Flight{Id: model.Id}).Find(&cs).Updates(model).Error
	if err != nil {
		return nil, err
	}

	if len(cs) == 0 {
		return nil, errors.New("RECORD NOT FOUND")
	}

	return model, nil
}

func (conn *dbmanager) FlightDetails(context context.Context, id uuid.UUID) (*models.Flight, error) {
	cs := &models.Flight{}
	err := conn.First(&models.Flight{Id: id}).Find(&cs).Error

	if err != nil {
		return nil, err
	}

	return cs, nil
}
