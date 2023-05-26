package models

import "github.com/google/uuid"

type Plane struct {
	Id         uuid.UUID
	Number     string
	NumOfSeats int32 `gorm:"column:num_of_seats"`
	Status     string
}
