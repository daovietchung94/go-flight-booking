package handlers

import (
	"context"
	"github.com/google/uuid"
	"go-training/grpc/plane/models"
	"go-training/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *PlaneHandler) UpdatePlane(context context.Context, request *pb.UpdatePlaneRequest) (*pb.Plane, error) {
	id, err := uuid.Parse(request.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	plane := &models.Plane{
		Id:         id,
		Number:     request.Number,
		NumOfSeats: request.NumOfSeats,
		Status:     request.Status,
	}

	c, err := h.planeRepository.UpdatePlane(context, plane)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return c.ToPBModel(), nil
}
