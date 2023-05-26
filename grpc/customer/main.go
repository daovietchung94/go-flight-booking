package main

import (
	"go-training/config"
	"go-training/grpc/customer/handlers"
	"go-training/grpc/customer/repository"
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

	err = log.Setup(conf.Logging)
	if err != nil {
		log.Fatal(err)
	}
	listen, err := net.Listen("tcp", conf.GRPCConf.CustomerGRPCConf.Host+":"+conf.GRPCConf.CustomerGRPCConf.Port)
	if err != nil {
		log.Fatal(err)
	}
	server := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer()),
	)

	customerRepository, err := repository.NewDBManager()

	if err != nil {
		log.Fatal(err)
	}

	handler, err := handlers.NewCustomerHandler(customerRepository)

	if err != nil {
		log.Fatal(err)
	}

	reflection.Register(server)
	pb.RegisterMyCustomerServer(server, handler)

	server.Serve(listen)
}
