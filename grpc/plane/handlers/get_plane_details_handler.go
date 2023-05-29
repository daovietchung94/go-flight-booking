package handlers

import (
	"context"
	"github.com/google/uuid"
	"go-training/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *PlaneHandler) GetPlaneDetails(ctx context.Context, request *pb.GetPlaneDetailsRequest) (*pb.Plane, error) {
	id, err := uuid.Parse(request.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	c, err := h.planeRepository.PlaneDetails(ctx, id)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return c.ToPBModel(), nil
}
