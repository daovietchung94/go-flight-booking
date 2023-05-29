package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.31

import (
	"context"
	"fmt"
	"go-training/clients/graphql/logistics_manager/generated"
	"go-training/clients/graphql/logistics_manager/model"
	"go-training/pb"
)

// FindPlaneByID is the resolver for the findPlaneByID field.
func (r *entityResolver) FindPlaneByID(ctx context.Context, id string) (*model.Plane, error) {
	pReq := &pb.GetPlaneDetailsRequest{
		Id: id,
	}

	c, err := r.MyPlaneClient.GetPlaneDetails(ctx, pReq)
	if err != nil {
		panic(fmt.Errorf(err.Error()))
	}

	dto := &model.Plane{
		ID:         c.Id,
		Number:     c.Number,
		NumOfSeats: int(c.NumOfSeats),
		Status:     model.PlaneStatus(c.Status),
	}

	return dto, nil
}

// Entity returns generated.EntityResolver implementation.
func (r *Resolver) Entity() generated.EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }