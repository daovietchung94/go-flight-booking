package models

import (
	"github.com/google/uuid"
	"go-training/pb"
	"time"
)

type Reservation struct {
	Id              uuid.UUID
	CustomerId      uuid.UUID `gorm:"column:customer_id"`
	FlightId        uuid.UUID `gorm:"column:flight_id"`
	ReservationDate time.Time `gorm:"column:reservation_date"`
	Status          string
}

func (m *Reservation) ToPBModel() *pb.Reservation {
	return &pb.Reservation{
		Id: m.Id.String(),
		Flight: &pb.Flight{
			Id: m.FlightId.String(),
		},
		Customer: &pb.Customer{
			Id: m.CustomerId.String(),
		},
		ReservationDate: &pb.Date{
			Year:  int32(m.ReservationDate.Year()),
			Month: int32(m.ReservationDate.Month()),
			Day:   int32(m.ReservationDate.Day()),
		},
		Status: m.Status,
	}
}
