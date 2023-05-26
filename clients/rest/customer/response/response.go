package response

import (
	"time"

	"github.com/gofrs/uuid"
)

type Customer struct {
	Id      uuid.UUID
	Email   string
	Name    string
	Address string
	DoB     time.Time
}
