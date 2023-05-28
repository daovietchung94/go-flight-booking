// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type CreatePlaneRequest struct {
	Number     string      `json:"number"`
	NumOfSeats int         `json:"numOfSeats"`
	Status     PlaneStatus `json:"status"`
}

type DeletePlaneRequest struct {
	ID string `json:"id"`
}

type DeletePlaneResponse struct {
	IsDeleted bool `json:"isDeleted"`
}

type GetPlaneStatusRequest struct {
	ID string `json:"id"`
}

type GetPlanesRequest struct {
	Page  int    `json:"page"`
	Limit int    `json:"limit"`
	Sort  string `json:"sort"`
}

type GetPlanesResponse struct {
	Page       int      `json:"page"`
	Limit      int      `json:"limit"`
	Sort       string   `json:"sort"`
	TotalRows  int      `json:"totalRows"`
	TotalPages int      `json:"totalPages"`
	Rows       []*Plane `json:"rows,omitempty"`
}

type Plane struct {
	ID         string      `json:"id"`
	Number     string      `json:"number"`
	NumOfSeats int         `json:"numOfSeats"`
	Status     PlaneStatus `json:"status"`
}

type UpdatePlaneRequest struct {
	ID         string      `json:"id"`
	Number     string      `json:"number"`
	NumOfSeats int         `json:"numOfSeats"`
	Status     PlaneStatus `json:"status"`
}

type UpdatePlaneStatusRequest struct {
	ID     string      `json:"id"`
	Status PlaneStatus `json:"status"`
}

type PlaneStatus string

const (
	PlaneStatusCleaning  PlaneStatus = "CLEANING"
	PlaneStatusRepairing PlaneStatus = "REPAIRING"
	PlaneStatusReady     PlaneStatus = "READY"
)

var AllPlaneStatus = []PlaneStatus{
	PlaneStatusCleaning,
	PlaneStatusRepairing,
	PlaneStatusReady,
}

func (e PlaneStatus) IsValid() bool {
	switch e {
	case PlaneStatusCleaning, PlaneStatusRepairing, PlaneStatusReady:
		return true
	}
	return false
}

func (e PlaneStatus) String() string {
	return string(e)
}

func (e *PlaneStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = PlaneStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid PlaneStatus", str)
	}
	return nil
}

func (e PlaneStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}