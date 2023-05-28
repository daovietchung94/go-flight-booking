package main

import (
	"context"
	"go-training/config"
	"go-training/grpc/plane/handlers"
	"go-training/grpc/plane/repository"
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

	err = log.Setup(conf.Logging, "plane_grpc")
	if err != nil {
		log.Fatal(err)
	}

	listen, err := net.Listen("tcp", conf.GRPCConf.PlaneGRPCConf.Host+":"+conf.GRPCConf.PlaneGRPCConf.Port)
	if err != nil {
		log.Fatal(err)
	}
	server := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer()),
	)

	repository, err := repository.NewDBManager()
	if err != nil {
		log.Fatal(err)
	}

	flightConn, err := grpc.Dial(conf.GRPCConf.FlightGRPCConf.Host+":"+conf.GRPCConf.FlightGRPCConf.Port, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	flightServiceClient := pb.NewMyFlightClient(flightConn)

	handler, err := handlers.NewPlaneHandler(flightServiceClient, repository)
	if err != nil {
		log.Fatal(err)
	}

	reflection.Register(server)
	pb.RegisterMyPlaneServer(server, handler)

	go func() {
		log.Infof("GRPC Server is listening on port: %v", conf.GRPCConf.PlaneGRPCConf.Port)
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
