package handlers

import (
	"context"
	"go-training/grpc/customer/requests"
	"go-training/pb"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *CustomerHandler) ChangeCustomerPassword(ctx context.Context, m *pb.ChangeCustomerPasswordRequest) (*pb.Empty, error) {
	id, err := uuid.Parse(m.Id)

	if err != nil {
		return nil, err
	}

	req := &requests.ChangeCustomerPasswordRequest{
		Id:          id,
		OldPassword: m.OldPassword,
		NewPassword: m.NewPassword,
	}

	err = h.customerRepository.UpdateCustomerPassword(ctx, req)

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if err != nil {
		return nil, err
	}

	return &pb.Empty{}, nil
}
