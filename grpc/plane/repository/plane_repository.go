package repository

import (
	"context"
	"errors"
	"go-training/grpc/plane/models"
	"go-training/pkg/pagination"
	"math"

	"github.com/google/uuid"
)

type PlaneRepository interface {
	GetPlanes(context context.Context, pagination pagination.Pagination) (*pagination.Pagination, error)
	CreatePlane(context context.Context, model *models.Plane) (*models.Plane, error)
	UpdatePlane(context context.Context, model *models.Plane) (*models.Plane, error)
	PlaneDetails(context context.Context, id uuid.UUID) (*models.Plane, error)
}

func (conn *dbmanager) GetPlanes(context context.Context, pagination pagination.Pagination) (*pagination.Pagination, error) {
	var totalRows int64
	var planes []*models.Plane
	err := conn.Model(&models.Plane{}).Count(&totalRows).Error
	if err != nil {
		return nil, err
	}
	totalPages := int(math.Ceil(float64(totalRows) / float64(pagination.Limit)))
	pagination.TotalRows = totalRows
	pagination.TotalPages = totalPages

	err = conn.Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order(pagination.GetSort()).Find(&planes).Error
	if err != nil {
		return nil, err
	}
	pagination.Rows = planes

	return &pagination, nil
}

func (conn *dbmanager) CreatePlane(context context.Context, model *models.Plane) (*models.Plane, error) {
	model.Id = uuid.New()
	if err := conn.Create(model).Error; err != nil {
		return nil, err
	}

	return model, nil
}

func (conn *dbmanager) UpdatePlane(context context.Context, model *models.Plane) (*models.Plane, error) {
	var cs []*models.Plane
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
