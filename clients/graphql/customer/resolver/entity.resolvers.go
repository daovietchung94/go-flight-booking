package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.31

import (
	"context"
	"fmt"
	"go-training/clients/graphql/customer/generated"
	"go-training/clients/graphql/customer/model"
	"go-training/pb"
	"go-training/pkg/utils"
	"time"
)

// FindCustomerByID is the resolver for the findCustomerByID field.
func (r *entityResolver) FindCustomerByID(ctx context.Context, id string) (*model.Customer, error) {
	pReq := &pb.GetCustomerDetailsRequest{
		Id: id,
	}

	pRes, err := r.MyCustomerClient.GetCustomerDetails(ctx, pReq)
	if err != nil {
		panic(fmt.Errorf(err.Error()))
	}

	dto := &model.Customer{
		ID:          pRes.Id,
		Name:        pRes.Name,
		DateOfBirth: time.Date(int(pRes.DateOfBirth.Year), time.Month(pRes.DateOfBirth.Month), int(pRes.DateOfBirth.Day), 0, 0, 0, 0, time.UTC),
		Address:     pRes.Address,
		Email:       pRes.Email,
	}

	return dto, nil
}

// FindFlightByID is the resolver for the findFlightByID field.
func (r *entityResolver) FindFlightByID(ctx context.Context, id string) (*model.Flight, error) {
	pReq := &pb.GetFlightDetailsRequest{
		Id: id,
	}

	pRes, err := r.MyFlightClient.GetFlightDetails(ctx, pReq)
	if err != nil {
		panic(fmt.Errorf(err.Error()))
	}

	dto := &model.Flight{
		ID:             pRes.Id,
		PlaneNumber:    pRes.PlaneNumber,
		AvailableSeats: int(pRes.AvailableSeats),
		FromCity:       pRes.FromCity,
		ToCity:         pRes.ToCity,
		DepTime:        time.Date(int(pRes.DepTime.Year), time.Month(pRes.DepTime.Month), int(pRes.DepTime.Day), 0, 0, 0, 0, time.UTC),
		ArrTime:        time.Date(int(pRes.ArrTime.Year), time.Month(pRes.ArrTime.Month), int(pRes.ArrTime.Day), 0, 0, 0, 0, time.UTC),
		Status:         pRes.Status,
	}

	return dto, nil
}

// FindReservationByID is the resolver for the findReservationByID field.
func (r *entityResolver) FindReservationByID(ctx context.Context, id string) (*model.Reservation, error) {
	pReq := &pb.GetReservationDetailsRequest{
		Id: id,
	}
	pRes, err := r.MyBookingClient.GetReservationDetails(ctx, pReq)
	if err != nil {
		panic(fmt.Errorf(err.Error()))
	}

	cReq := &pb.GetCustomerDetailsRequest{
		Id: pRes.Customer.Id,
	}
	cRes, err := r.MyCustomerClient.GetCustomerDetails(ctx, cReq)
	if err != nil {
		panic(fmt.Errorf(err.Error()))
	}

	fReq := &pb.GetFlightDetailsRequest{
		Id: pRes.Flight.Id,
	}
	fRes, err := r.MyFlightClient.GetFlightDetails(ctx, fReq)
	if err != nil {
		panic(fmt.Errorf(err.Error()))
	}

	dto := &model.Reservation{
		ID: pRes.Id,
		Customer: &model.Customer{
			ID:          cRes.Id,
			Name:        cRes.Name,
			DateOfBirth: time.Date(int(cRes.DateOfBirth.Year), time.Month(cRes.DateOfBirth.Month), int(cRes.DateOfBirth.Day), 0, 0, 0, 0, time.UTC),
			Address:     cRes.Address,
			Email:       cRes.Email,
		},
		Flight: &model.Flight{
			ID:             fRes.Id,
			PlaneNumber:    fRes.PlaneNumber,
			AvailableSeats: int(fRes.AvailableSeats),
			FromCity:       fRes.FromCity,
			ToCity:         fRes.ToCity,
			DepTime:        utils.ToTime(fRes.DepTime),
			ArrTime:        utils.ToTime(fRes.ArrTime),
			Status:         fRes.Status,
		},
		ReservationDate: utils.ToTime(pRes.ReservationDate),
		Status:          pRes.Status,
	}

	return dto, nil
}

// Entity returns generated.EntityResolver implementation.
func (r *Resolver) Entity() generated.EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }
