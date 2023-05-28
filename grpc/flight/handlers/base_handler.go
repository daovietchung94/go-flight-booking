package handlers

import (
	"go-training/grpc/flight/repository"
	"go-training/pb"
)

type FlightHandler struct {
	pb.UnimplementedMyFlightServer
	flightRepository repository.FlightRepository
	myPlaneClient    pb.MyPlaneClient
}

func NewFlightHandler(flightRepository repository.FlightRepository, myPlaneClient pb.MyPlaneClient) (*FlightHandler, error) {
	return &FlightHandler{flightRepository: flightRepository, myPlaneClient: myPlaneClient}, nil
}
