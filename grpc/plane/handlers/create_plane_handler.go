package handlers

import (
	"context"
	"go-training/grpc/plane/models"
	"go-training/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *PlaneHandler) CreatePlane(ctx context.Context, request *pb.CreatePlaneRequest) (*pb.Plane, error) {
	req := &models.Plane{
		Number:     request.Number,
		NumOfSeats: request.NumOfSeats,
		Status:     request.Status,
	}

	c, err := h.planeRepository.CreatePlane(ctx, req)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return c.ToPBModel(), nil
}
