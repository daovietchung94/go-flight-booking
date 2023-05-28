package handlers

import (
	"go-training/grpc/plane/repository"
	"go-training/pb"
)

type PlaneHandler struct {
	pb.UnimplementedMyPlaneServer
	myFlightClient  pb.MyFlightClient
	planeRepository repository.PlaneRepository
}

func NewPlaneHandler(myFlightClient pb.MyFlightClient, planeRepository repository.PlaneRepository) (*PlaneHandler, error) {
	return &PlaneHandler{planeRepository: planeRepository, myFlightClient: myFlightClient}, nil
}
