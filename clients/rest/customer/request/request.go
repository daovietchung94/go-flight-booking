package request

import (
	"time"

	"github.com/gofrs/uuid"
)

type CreateCustomerRequest struct {
	Email    string
	Name     string
	Address  string
	DoB      time.Time
	Password string
}

type UpdateCustomerRequest struct {
	Id      uuid.UUID
	Email   string
	Name    string
	Address string
	DoB     time.Time
}

type ChangeCustomerPasswordRequest struct {
	Id          uuid.UUID
	OldPassword string
	NewPassword string
}
