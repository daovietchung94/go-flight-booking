package handlers

import (
	"go-training/grpc/flight/repository"
	"go-training/pb"
)

type FlightHandler struct {
	pb.UnimplementedMyFlightServer
	flightRepository repository.FlightRepository
}

func NewFlightHandler(flightRepository repository.FlightRepository) (*FlightHandler, error) {
	return &FlightHandler{flightRepository: flightRepository}, nil
}
