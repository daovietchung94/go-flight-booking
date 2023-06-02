package handlers

import (
	"go-training/grpc/booking/repository"
	"go-training/pb"
	"sync"
)

type BookingHandler struct {
	mutex *sync.Mutex
	pb.UnimplementedMyBookingServer
	bookingRepository repository.BookingRepository
	myPlaneClient     pb.MyPlaneClient
	myFlightClient    pb.MyFlightClient
	myCustomerClient  pb.MyCustomerClient
}

func NewBookingHandler(mutex *sync.Mutex, bookingRepository repository.BookingRepository, myPlaneClient pb.MyPlaneClient, myFlightClient pb.MyFlightClient, myCustomerClient pb.MyCustomerClient) (*BookingHandler, error) {
	return &BookingHandler{
		mutex:             mutex,
		bookingRepository: bookingRepository,
		myPlaneClient:     myPlaneClient,
		myFlightClient:    myFlightClient,
		myCustomerClient:  myCustomerClient,
	}, nil
}
