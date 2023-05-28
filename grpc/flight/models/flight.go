package models

import (
	"go-training/pb"
	"time"

	"github.com/google/uuid"
)

type Flight struct {
	Id          uuid.UUID
	PlaneNumber string    `gorm:"column:plane_number"`
	NumOfSeats  int       `gorm:"column:num_of_seats"`
	FromCity    string    `gorm:"column:from_city"`
	ToCity      string    `gorm:"column:to_city"`
	DepTime     time.Time `gorm:"column:dep_time"`
	ArrTime     time.Time `gorm:"column:arr_time"`
	IsLanded    bool      `gorm:"column:is_landed;default:false"`
}

func (m *Flight) ToPBModel() *pb.Flight {
	return &pb.Flight{
		Id:          m.Id.String(),
		PlaneNumber: m.PlaneNumber,
		NumOfSeats:  int32(m.NumOfSeats),
		FromCity:    m.FromCity,
		ToCity:      m.ToCity,
		DepTime: &pb.Date{
			Year:  int32(m.DepTime.Year()),
			Month: int32(m.DepTime.Month()),
			Day:   int32(m.DepTime.Day()),
		},
		ArrTime: &pb.Date{
			Year:  int32(m.ArrTime.Year()),
			Month: int32(m.ArrTime.Month()),
			Day:   int32(m.ArrTime.Day()),
		},
	}
}
