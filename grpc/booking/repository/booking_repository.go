package repository

import (
	"context"
	"github.com/google/uuid"
	"go-training/grpc/booking/models"
)

type BookingRepository interface {
	CreateReservation(context context.Context, model *models.Reservation) (*models.Reservation, error)
	ReservationDetails(context context.Context, id uuid.UUID) (*models.Reservation, error)
}

func (conn *dbmanager) CreateReservation(context context.Context, model *models.Reservation) (*models.Reservation, error) {
	model.Id = uuid.New()
	if err := conn.db.Create(model).Error; err != nil {
		return nil, err
	}

	return model, nil
}

func (conn *dbmanager) ReservationDetails(context context.Context, id uuid.UUID) (*models.Reservation, error) {
	cs := &models.Reservation{}
	err := conn.db.First(&models.Reservation{Id: id}).Find(&cs).Error

	if err != nil {
		return nil, err
	}

	return cs, nil
}
