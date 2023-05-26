package handlers

import (
	"go-training/grpc/customer/repository"
	"go-training/pb"
)

type CustomerHandler struct {
	pb.UnimplementedMyCustomerServer
	customerRepository repository.CustomerRepository
}

func NewCustomerHandler(customerRepository repository.CustomerRepository) (*CustomerHandler, error) {
	return &CustomerHandler{customerRepository: customerRepository}, nil
}
