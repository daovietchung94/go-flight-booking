package handlers

import (
	"context"
	"go-training/pb"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *CustomerHandler) GetCustomerDetails(ctx context.Context, m *pb.GetCustomerDetailsRequest) (*pb.Customer, error) {
	id, err := uuid.Parse(m.Id)
	if err != nil {
		return nil, err
	}

	c, err := h.customerRepository.CustomerDetails(ctx, id)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return c.ToPBModel(), nil
}
