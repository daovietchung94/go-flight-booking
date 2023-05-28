package resolver

import "go-training/pb"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	MyCustomerClient pb.MyCustomerClient
	MyPlaneClient    pb.MyPlaneClient
	MyFlightClient   pb.MyFlightClient
}
