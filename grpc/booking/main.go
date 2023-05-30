package main

import (
	"context"
	"go-training/config"
	"go-training/grpc/booking/handlers"
	"go-training/grpc/booking/repository"
	"go-training/pb"
	"net"
	"os"
	"os/signal"
	"syscall"

	log "go-training/logger"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	configPath = kingpin.Flag("config", "Location of config.json.").Default("./config.json").String()
)

func main() {
	ctx := context.Background()
	// Parse the CLI flags and load the config
	kingpin.CommandLine.HelpFlag.Short('h')
	kingpin.Parse()

	// Load the config
	conf, err := config.LoadConfig(*configPath)
	if err != nil {
		log.Fatal(err)
	}

	err = log.Setup(conf.Logging, "booking_grpc")
	if err != nil {
		log.Fatal(err)
	}
	listen, err := net.Listen("tcp", conf.GRPCConf.BookingGRPCConf.Host+":"+conf.GRPCConf.BookingGRPCConf.Port)
	if err != nil {
		log.Fatal(err)
	}
	server := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer()),
	)

	bookingRepository, err := repository.NewDBManager(conf)
	if err != nil {
		log.Fatal(err)
	}

	planeConn, err := grpc.Dial(conf.GRPCConf.PlaneGRPCConf.Host+":"+conf.GRPCConf.PlaneGRPCConf.Port, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	planeServiceClient := pb.NewMyPlaneClient(planeConn)

	flightConn, err := grpc.Dial(conf.GRPCConf.FlightGRPCConf.Host+":"+conf.GRPCConf.FlightGRPCConf.Port, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	flightServiceClient := pb.NewMyFlightClient(flightConn)

	customerConn, err := grpc.Dial(conf.GRPCConf.CustomerGRPCConf.Host+":"+conf.GRPCConf.CustomerGRPCConf.Port, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	customerServiceClient := pb.NewMyCustomerClient(customerConn)

	handler, err := handlers.NewBookingHandler(bookingRepository, planeServiceClient, flightServiceClient, customerServiceClient)
	if err != nil {
		log.Fatal(err)
	}

	reflection.Register(server)
	pb.RegisterMyBookingServer(server, handler)

	go func() {
		log.Infof("GRPC Server is listening on port: %v", conf.GRPCConf.FlightGRPCConf.Port)
		log.Fatal(server.Serve(listen))
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	select {
	case v := <-c:
		log.Errorf("signal.Notify: %v", v)
	case done := <-ctx.Done():
		log.Errorf("ctx.Done: %v", done)
	}

	server.GracefulStop()
	log.Info("Server Exited Properly")
}
