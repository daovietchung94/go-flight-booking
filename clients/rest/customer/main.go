package main

import (
	"go-training/clients/rest/customer/handler"
	"go-training/config"
	log "go-training/logger"
	"go-training/pb"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
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

	err = log.Setup(conf.Logging)
	if err != nil {
		log.Fatal(err)
	}

	//Create grpc client connect
	customerConn, err := grpc.Dial(conf.GRPCConf.CustomerGRPCConf.Host+
		":"+
		conf.GRPCConf.CustomerGRPCConf.Port, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	//Singleton
	customerServiceClient := pb.NewMyCustomerClient(customerConn)

	//Handler for GIN Gonic
	h := handler.NewCustomerApiHandler(customerServiceClient)
	os.Setenv("GIN_MODE", "debug")
	g := gin.Default()

	//Create routes
	gr := g.Group("/api/customers")
	gr.POST("", h.CreateCustomer)
	gr.PUT("/:id", h.UpdateCustomer)
	gr.PUT("/change-password", h.ChangeCustomerPassword)
	//Listen and serve
	http.ListenAndServe(conf.ServerConf.CustomerApiConf.Host+
		":"+
		conf.ServerConf.CustomerApiConf.Port, g)
}
