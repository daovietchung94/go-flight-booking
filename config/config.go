package config

import (
	"encoding/json"
	log "go-training/logger"
	"io/ioutil"
)

type Config struct {
	Logging      *log.Config    `json:"logging"`
	GraphQLConf  GraphQLServer  `json:"graphql"`
	GRPCConf     GRPCServer     `json:"grpc"`
	DatabaseConf DatabaseServer `json:"database"`
}

type GraphQLServer struct {
	CustomerConf         Server `json:"customer"`
	FlightManagerConf    Server `json:"flightManager"`
	LogisticsManagerConf Server `json:"logisticsManager"`
}

type GRPCServer struct {
	CustomerGRPCConf Server `json:"customer"`
	PlaneGRPCConf    Server `json:"plane"`
	FlightGRPCConf   Server `json:"flight"`
	BookingGRPCConf  Server `json:"booking"`
}

type DatabaseServer struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

type Server struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

func LoadConfig(filepath string) (*Config, error) {
	configFile, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	config := &Config{}
	err = json.Unmarshal(configFile, config)
	if err != nil {
		return nil, err
	}
	if config.Logging == nil {
		config.Logging = &log.Config{}
	}
	return config, nil
}
