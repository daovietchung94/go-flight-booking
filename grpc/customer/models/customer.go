package models

import (
	"go-training/pb"
	"time"

	"github.com/google/uuid"
)

type Customer struct {
	Id          uuid.UUID
	Email       string
	Name        string
	Address     string
	DateOfBirth time.Time `gorm:"column:date_of_birth"`
	Password    string
}

func (m *Customer) ToPBModel() *pb.Customer {
	return &pb.Customer{
		Id:   m.Id.String(),
		Name: m.Name,
		DateOfBirth: &pb.Date{
			Year:  int32(m.DateOfBirth.UTC().Year()),
			Month: int32(m.DateOfBirth.UTC().Month()),
			Day:   int32(m.DateOfBirth.UTC().Day()),
		},
		Address: m.Address,
		Email:   m.Email,
	}
}
