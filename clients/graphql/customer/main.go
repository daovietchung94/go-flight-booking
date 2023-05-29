package main

import (
	"go-training/clients/graphql/customer/generated"
	"go-training/clients/graphql/customer/resolver"
	"go-training/config"
	log "go-training/logger"
	"go-training/pb"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"google.golang.org/grpc"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	configPath = kingpin.Flag("config", "Location of config.json.").Default("./config.json").String()
)

func main() {
	// Parse the CLI flags and load the config
	kingpin.CommandLine.HelpFlag.Short('h')
	kingpin.Parse()

	// Load the config
	conf, err := config.LoadConfig(*configPath)
	if err != nil {
		log.Fatal(err)
	}

	err = log.Setup(conf.Logging, "graphql")
	if err != nil {
		log.Fatal(err)
	}

	//Create grpc client connect
	customerConn, err := grpc.Dial(conf.GRPCConf.CustomerGRPCConf.Host+":"+conf.GRPCConf.CustomerGRPCConf.Port, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	planeConn, err := grpc.Dial(conf.GRPCConf.PlaneGRPCConf.Host+":"+conf.GRPCConf.PlaneGRPCConf.Port, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	flightConn, err := grpc.Dial(conf.GRPCConf.FlightGRPCConf.Host+":"+conf.GRPCConf.FlightGRPCConf.Port, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	//Singleton
	customerServiceClient := pb.NewMyCustomerClient(customerConn)
	planeServiceClient := pb.NewMyPlaneClient(planeConn)
	flightServiceClient := pb.NewMyFlightClient(flightConn)

	graphConf := generated.Config{
		Resolvers: &resolver.Resolver{
			MyCustomerClient: customerServiceClient,
			MyPlaneClient:    planeServiceClient,
			MyFlightClient:   flightServiceClient,
		},
	}
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(graphConf))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Infof("connect to http://%s:%s/ for GraphQL playground", conf.GraphQLConf.CustomerConf.Host, conf.GraphQLConf.CustomerConf.Port)
	log.Fatal(http.ListenAndServe(":"+conf.GraphQLConf.CustomerConf.Port, nil))
}
