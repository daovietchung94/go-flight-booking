package repository

import (
	"context"
	"errors"
	"go-training/grpc/plane/models"

	"github.com/google/uuid"
)

type PlaneRepository interface {
	CreatePlane(context context.Context, model *models.Plane) (*models.Plane, error)
	UpdatePlane(context context.Context, model *models.Plane) (*models.Plane, error)
	PlaneDetails(context context.Context, id uuid.UUID) (*models.Plane, error)
}

func (conn *dbmanager) CreatePlane(context context.Context, model *models.Plane) (*models.Plane, error) {
	model.Id = uuid.New()
	if err := conn.Create(model).Error; err != nil {
		return nil, err
	}

	return model, nil
}

func (conn *dbmanager) UpdatePlane(context context.Context, model *models.Plane) (*models.Plane, error) {
	cs := []*models.Plane{}
	err := conn.Where(&models.Plane{Id: model.Id}).Find(&cs).Updates(model).Error
	if err != nil {
		return nil, err
	}

	if len(cs) == 0 {
		return nil, errors.New("RECORD NOT FOUND")
	}

	return model, nil
}

func (conn *dbmanager) PlaneDetails(context context.Context, id uuid.UUID) (*models.Plane, error) {
	cs := &models.Plane{}
	err := conn.First(&models.Plane{Id: id}).Find(&cs).Error

	if err != nil {
		return nil, err
	}

	return cs, nil
}
