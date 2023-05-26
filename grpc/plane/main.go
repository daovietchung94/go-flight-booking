package main

import (
	"go-training/config"
	"go-training/grpc/plane/handlers"
	"go-training/grpc/plane/repository"
	"go-training/pb"
	"net"

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

	handler, err := handlers.NewPlaneHandler(repository)
	if err != nil {
		log.Fatal(err)
	}

	reflection.Register(server)
	pb.RegisterMyPlaneServer(server, handler)

	log.Infof("Plane GRPC Server is listening on port: %v", conf.GRPCConf.PlaneGRPCConf.Port)
	log.Fatal(server.Serve(listen))
}
