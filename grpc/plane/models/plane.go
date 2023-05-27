package models

import (
	"github.com/google/uuid"
	"go-training/pb"
)

type Plane struct {
	Id         uuid.UUID
	Number     string
	NumOfSeats int32 `gorm:"column:num_of_seats"`
	Status     string
}

func (m *Plane) ToPBModel() *pb.Plane {
	return &pb.Plane{
		Id:         m.Id.String(),
		Number:     m.Number,
		NumOfSeats: m.NumOfSeats,
		Status:     m.Status,
	}
}
