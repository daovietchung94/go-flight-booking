package handlers

import (
	"context"
	"go-training/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *PlaneHandler) GetPlaneByNumber(context context.Context, request *pb.GetPlaneByNumberRequest) (*pb.Plane, error) {
	c, err := h.planeRepository.GetPlaneByNumber(context, request.Number)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return c.ToPBModel(), nil
}
