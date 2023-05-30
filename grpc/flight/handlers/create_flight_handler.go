package handlers

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"go-training/grpc/flight/models"
	"go-training/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

func (h *FlightHandler) CreateFlight(context context.Context, request *pb.CreateFlightRequest) (*pb.Flight, error) {
	plane, err := h.myPlaneClient.GetPlaneByNumber(context, &pb.GetPlaneByNumberRequest{Number: request.PlaneNumber})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	fmt.Print(plane)

	flight := &models.Flight{
		Id:             uuid.New(),
		PlaneNumber:    request.PlaneNumber,
		AvailableSeats: int(plane.NumOfSeats),
		FromCity:       request.FromCity,
		ToCity:         request.ToCity,
		DepTime:        time.Date(int(request.DepTime.Year), time.Month(request.DepTime.Month), int(request.DepTime.Day), 0, 0, 0, 0, time.Local),
		ArrTime:        time.Date(int(request.ArrTime.Year), time.Month(request.ArrTime.Month), int(request.ArrTime.Day), 0, 0, 0, 0, time.Local),
	}

	c, err := h.flightRepository.CreateFlight(context, flight)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return c.ToPBModel(), nil
}
