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
		ReservationDate: &pb.DateTime{
			Year:   int32(m.ReservationDate.UTC().Year()),
			Month:  int32(m.ReservationDate.UTC().Month()),
			Day:    int32(m.ReservationDate.UTC().Day()),
			Hour:   int32(m.ReservationDate.UTC().Hour()),
			Minute: int32(m.ReservationDate.UTC().Minute()),
			Second: int32(m.ReservationDate.UTC().Second()),
		},
		Status: m.Status,
	}
}
