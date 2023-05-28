package handlers

import (
	"context"
	"github.com/google/uuid"
	"go-training/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *PlaneHandler) DeletePlane(context context.Context, request *pb.DeletePlaneRequest) (*pb.DeletePlaneResponse, error) {
	id, err := uuid.Parse(request.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	err = h.planeRepository.DeletePlane(context, id)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.DeletePlaneResponse{IsDeleted: true}, nil
}
