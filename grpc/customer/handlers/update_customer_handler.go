package handlers

import (
	"context"
	"go-training/grpc/customer/models"
	log "go-training/logger"
	"go-training/pb"
	"time"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *CustomerHandler) UpdateCustomer(ctx context.Context, m *pb.Customer) (*pb.Customer, error) {

	id, err := uuid.Parse(m.Id)

	if err != nil {
		log.Fatal(err)
	}

	req := &models.Customer{
		Id:      id,
		Email:   m.Email,
		Name:    m.Name,
		Address: m.Address,
		DoB:     time.Date(int(m.DateOfBirth.Year), time.Month(m.DateOfBirth.Month), int(m.DateOfBirth.Day), 0, 0, 0, 0, time.UTC),
	}

	c, err := h.customerRepository.UpdateCustomer(ctx, req)

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if err != nil {
		return nil, err
	}

	return c.ToPBModel(), nil
}
