package models

import (
	"go-training/pb"
	"time"

	"github.com/google/uuid"
)

type Flight struct {
	Id             uuid.UUID
	PlaneNumber    string    `gorm:"column:plane_number"`
	AvailableSeats int       `gorm:"column:available_seats"`
	FromCity       string    `gorm:"column:from_city"`
	ToCity         string    `gorm:"column:to_city"`
	DepTime        time.Time `gorm:"column:dep_time"`
	ArrTime        time.Time `gorm:"column:arr_time"`
	Status         string
}

func (m *Flight) ToPBModel() *pb.Flight {
	return &pb.Flight{
		Id:             m.Id.String(),
		PlaneNumber:    m.PlaneNumber,
		AvailableSeats: int32(m.AvailableSeats),
		FromCity:       m.FromCity,
		ToCity:         m.ToCity,
		DepTime: &pb.DateTime{
			Year:   int32(m.DepTime.UTC().Year()),
			Month:  int32(m.DepTime.UTC().Month()),
			Day:    int32(m.DepTime.UTC().Day()),
			Hour:   int32(m.DepTime.UTC().Hour()),
			Minute: int32(m.DepTime.UTC().Minute()),
			Second: int32(m.DepTime.UTC().Second()),
		},
		ArrTime: &pb.DateTime{
			Year:   int32(m.ArrTime.UTC().Year()),
			Month:  int32(m.ArrTime.UTC().Month()),
			Day:    int32(m.ArrTime.UTC().Day()),
			Hour:   int32(m.ArrTime.UTC().Hour()),
			Minute: int32(m.ArrTime.UTC().Minute()),
			Second: int32(m.ArrTime.UTC().Second()),
		},
		Status: m.Status,
	}
}
