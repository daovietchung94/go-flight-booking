package handlers

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"go-training/grpc/booking/models"
	"go-training/pb"
	"go-training/pkg/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

func (h *BookingHandler) CreateReservation(ctx context.Context, input *pb.CreateReservationRequest) (*pb.Reservation, error) {
	h.mutex.Lock()
	defer h.mutex.Unlock()

	cRes, err := h.myCustomerClient.GetCustomerDetails(ctx, &pb.GetCustomerDetailsRequest{
		Id: input.CustomerId,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	fRes, err := h.myFlightClient.GetFlightDetails(ctx, &pb.GetFlightDetailsRequest{Id: input.FlightId})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	reservationDate := time.Now().UTC()
	expirationTime := utils.ToTime(fRes.DepTime).Add(-45 * time.Minute)
	if reservationDate.After(expirationTime) {
		return nil, errors.New("BOOKING TIME HAS EXPIRED")
	}
	if fRes.AvailableSeats == 0 {
		return nil, errors.New("SEATS ARE SOLD OUT")
	}

	availableSeats := fRes.AvailableSeats - 1
	_, err = h.myFlightClient.UpdateFlightAvailableSeats(ctx, &pb.UpdateFlightAvailableSeatsRequest{
		Id:             fRes.Id,
		AvailableSeats: availableSeats,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	cId, err := uuid.Parse(cRes.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	fId, err := uuid.Parse(fRes.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	rRes, err := h.bookingRepository.CreateReservation(ctx, &models.Reservation{
		CustomerId:      cId,
		FlightId:        fId,
		ReservationDate: reservationDate,
		Status:          "CREATED",
	})

	dto := &pb.Reservation{
		Id:              rRes.Id.String(),
		Customer:        cRes,
		Flight:          fRes,
		ReservationDate: utils.ToPbDateTime(rRes.ReservationDate),
		Status:          rRes.Status,
	}

	return dto, nil
}
