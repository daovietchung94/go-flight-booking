package handlers

import (
	"context"
	"go-training/grpc/plane/models"
	"go-training/pb"
	"go-training/pkg/pagination"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *PlaneHandler) GetPlanes(ctx context.Context, request *pb.GetPlanesRequest) (*pb.GetPlanesResponse, error) {
	req := pagination.Pagination{
		Limit: int(request.Limit),
		Page:  int(request.Page),
		Sort:  request.Sort,
	}

	c, err := h.planeRepository.GetPlanes(ctx, req)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	rows := c.Rows.([]*models.Plane)
	planes := make([]*pb.Plane, len(rows))

	for i, v := range rows {
		planes[i] = v.ToPBModel()
	}

	res := &pb.GetPlanesResponse{
		Page:       int32(c.Page),
		Limit:      int32(c.Limit),
		Sort:       c.Sort,
		TotalPages: int32(c.TotalPages),
		TotalRows:  int32(c.TotalRows),
		Rows:       planes,
	}

	return res, nil
}
