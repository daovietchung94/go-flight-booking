package handlers

import (
	"go-training/grpc/booking/repository"
	"go-training/pb"
)

type BookingHandler struct {
	pb.UnimplementedMyBookingServer
	bookingRepository repository.BookingRepository
	myPlaneClient     pb.MyPlaneClient
	myFlightClient    pb.MyFlightClient
	myCustomerClient  pb.MyCustomerClient
}

func NewBookingHandler(bookingRepository repository.BookingRepository, myPlaneClient pb.MyPlaneClient, myFlightClient pb.MyFlightClient, myCustomerClient pb.MyCustomerClient) (*BookingHandler, error) {
	return &BookingHandler{
		bookingRepository: bookingRepository,
		myPlaneClient:     myPlaneClient,
		myFlightClient:    myFlightClient,
		myCustomerClient:  myCustomerClient,
	}, nil
}
