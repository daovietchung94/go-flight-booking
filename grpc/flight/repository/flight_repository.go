package repository

import (
	"context"
	"errors"
	"go-training/grpc/flight/models"
	"go-training/pb"
	"go-training/pkg/pagination"
	"math"
	"time"

	"github.com/google/uuid"
)

type FlightRepository interface {
	GetFlights(context context.Context, pagination pagination.Pagination) (*pagination.Pagination, error)
	CreateFlight(context context.Context, model *models.Flight) (*models.Flight, error)
	UpdateFlight(context context.Context, model *models.Flight) (*models.Flight, error)
	FlightDetails(context context.Context, id uuid.UUID) (*models.Flight, error)
}

func (conn *dbmanager) GetFlights(context context.Context, pagination pagination.Pagination) (*pagination.Pagination, error) {
	var totalRows int64
	var flights []*models.Flight
	query := conn.Model(&models.Flight{}).Where("is_landed = ?", false)
	if pagination.Filter != nil {
		filter := pagination.Filter.(*pb.FlightFilter)
		if filter.City != "" {
			query = query.Where("from_city LIKE ? OR to_city LIKE ?", "%"+filter.City+"%", "%"+filter.City+"%")
		}
		if filter.Time != nil {
			startOfDay := time.Date(int(filter.Time.Year), time.Month(filter.Time.Month), int(filter.Time.Day), 0, 0, 0, 0, time.Local)
			endOfDay := startOfDay.Add(24 * time.Hour)
			query = query.Where("(dep_time >= ? AND dep_time < ?) OR (arr_time >= ? AND arr_time < ?)", startOfDay, endOfDay, startOfDay, endOfDay)
		}
	}

	err := query.Count(&totalRows).Error
	if err != nil {
		return nil, err
	}
	totalPages := int(math.Ceil(float64(totalRows) / float64(pagination.Limit)))
	pagination.TotalRows = totalRows
	pagination.TotalPages = totalPages

	err = query.Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order(pagination.GetSort()).Find(&flights).Error
	if err != nil {
		return nil, err
	}
	pagination.Rows = flights

	return &pagination, nil
}

func (conn *dbmanager) CreateFlight(context context.Context, model *models.Flight) (*models.Flight, error) {
	model.Id = uuid.New()
	if err := conn.Create(model).Error; err != nil {
		return nil, err
	}

	return model, nil
}

func (conn *dbmanager) UpdateFlight(context context.Context, model *models.Flight) (*models.Flight, error) {
	var cs []*models.Flight
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
