package requests

import "github.com/google/uuid"

type ChangeCustomerPasswordRequest struct {
	Id          uuid.UUID
	OldPassword string
	NewPassword string
}
