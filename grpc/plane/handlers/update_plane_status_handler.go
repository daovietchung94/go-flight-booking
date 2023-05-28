package handlers

import (
	"context"
	"github.com/google/uuid"
	"go-training/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *PlaneHandler) UpdatePlaneStatus(context context.Context, request *pb.UpdatePlaneStatusRequest) (*pb.Plane, error) {
	planeId, err := uuid.Parse(request.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	c, err := h.planeRepository.UpdatePlaneStatus(context, planeId, request.Status)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return c.ToPBModel(), nil
}
