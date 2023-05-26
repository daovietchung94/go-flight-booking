package models

import (
	"time"

	"github.com/google/uuid"
)

type Flight struct {
	Id         uuid.UUID
	PlaneID    uuid.UUID `gorm:"column:plane_id"`
	NumOfSeats int       `gorm:"column:num_of_seats"`
	FromCity   string    `gorm:"column:from_city"`
	ToCity     string    `gorm:"column:to_city"`
	DepTime    time.Time `gorm:"column:dep_time"`
	ArrTime    time.Time `gorm:"column:arr_time"`
}
