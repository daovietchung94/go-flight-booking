package handlers

import (
	"context"
	"go-training/grpc/flight/models"
	"go-training/pb"
	"go-training/pkg/pagination"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *FlightHandler) GetFlights(ctx context.Context, request *pb.GetFlightsRequest) (*pb.GetFlightsResponse, error) {
	req := pagination.Pagination{
		Limit:  int(request.Limit),
		Page:   int(request.Page),
		Sort:   request.Sort,
		Filter: request.Filter,
	}

	c, err := h.flightRepository.GetFlights(ctx, req)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	rows := c.Rows.([]*models.Flight)
	flights := make([]*pb.Flight, len(rows))

	for i, v := range rows {
		flights[i] = v.ToPBModel()
	}

	res := &pb.GetFlightsResponse{
		Page:       int32(c.Page),
		Limit:      int32(c.Limit),
		Sort:       c.Sort,
		TotalPages: int32(c.TotalPages),
		TotalRows:  int32(c.TotalRows),
		Rows:       flights,
	}

	return res, nil
}
