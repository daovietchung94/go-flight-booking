package handlers

import (
	"context"
	"github.com/google/uuid"
	"go-training/pb"
	"go-training/pkg/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *FlightHandler) UpdateFlightAvailableSeats(ctx context.Context, input *pb.UpdateFlightAvailableSeatsRequest) (*pb.Flight, error) {
	id, err := uuid.Parse(input.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	res, err := h.flightRepository.UpdateFlightAvailableSeats(ctx, id, int(input.AvailableSeats))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	flight := &pb.Flight{
		Id:             res.Id.String(),
		PlaneNumber:    res.PlaneNumber,
		AvailableSeats: int32(res.AvailableSeats),
		FromCity:       res.FromCity,
		ToCity:         res.ToCity,
		DepTime:        utils.ToPbDateTime(res.DepTime),
		ArrTime:        utils.ToPbDateTime(res.ArrTime),
		Status:         res.Status,
	}

	return flight, nil
}
