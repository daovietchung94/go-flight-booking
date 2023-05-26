package handlers

import (
	"context"
	"go-training/grpc/customer/models"
	"go-training/pb"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *CustomerHandler) CreateCustomer(ctx context.Context, m *pb.Customer) (*pb.Customer, error) {

	req := &models.Customer{
		Email:   m.Email,
		Name:    m.Name,
		Address: m.Address,
		DoB:     time.Date(int(m.DateOfBirth.Year), time.Month(m.DateOfBirth.Month), int(m.DateOfBirth.Day), 0, 0, 0, 0, time.UTC),
	}

	c, err := h.customerRepository.CreateCustomer(ctx, req)

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if err != nil {
		return nil, err
	}

	return c.ToPBModel(), nil
}
