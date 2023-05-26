package handlers

import (
	"go-training/grpc/plane/repository"
	"go-training/pb"
)

type PlaneHandler struct {
	pb.UnimplementedMyPlaneServer
	planeRepository repository.PlaneRepository
}

func NewPlaneHandler(planeRepository repository.PlaneRepository) (*PlaneHandler, error) {
	return &PlaneHandler{planeRepository: planeRepository}, nil
}
