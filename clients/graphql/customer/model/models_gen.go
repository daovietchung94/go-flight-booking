// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"
)

type CreateCustomerRequest struct {
	Name        string    `json:"name"`
	DateOfBirth time.Time `json:"dateOfBirth"`
	Address     string    `json:"address"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
}

type Customer struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	DateOfBirth time.Time `json:"dateOfBirth"`
	Address     string    `json:"address"`
	Email       string    `json:"email"`
}

func (Customer) IsEntity() {}

type Flight struct {
	ID             string    `json:"id"`
	PlaneNumber    string    `json:"planeNumber"`
	AvailableSeats int       `json:"availableSeats"`
	FromCity       string    `json:"fromCity"`
	ToCity         string    `json:"toCity"`
	DepTime        time.Time `json:"depTime"`
	ArrTime        time.Time `json:"arrTime"`
	Status         string    `json:"status"`
}

func (Flight) IsEntity() {}

type GetCustomerDetailsRequest struct {
	ID string `json:"id"`
}

type GetReservationDetailsRequest struct {
	ID string `json:"id"`
}

type MakeReservationRequest struct {
	CustomerID string `json:"customerId"`
	FlightID   string `json:"flightId"`
}

type Reservation struct {
	ID              string    `json:"id"`
	Customer        *Customer `json:"customer"`
	Flight          *Flight   `json:"flight"`
	ReservationDate time.Time `json:"reservationDate"`
	Status          string    `json:"status"`
}

func (Reservation) IsEntity() {}
