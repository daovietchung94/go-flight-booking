package config

import (
	"encoding/json"
	log "go-training/logger"
	"io/ioutil"
)

type Config struct {
	Logging    *log.Config `json:"logging"`
	GraphConf  Server      `json:"graph"`
	ServerConf ApiServer   `json:"api"`
	GRPCConf   GRPCServer  `json:"grpc"`
}

type ApiServer struct {
	CustomerApiConf Server `json:"customer"`
	PlaneApiConf    Server `json:"plane"`
	FlightApiConf   Server `json:"flight"`
	BookingApiConf  Server `json:"booking"`
}

type GRPCServer struct {
	CustomerGRPCConf Server `json:"customer"`
	PlaneGRPCConf    Server `json:"plane"`
	FlightGRPCConf   Server `json:"flight"`
	BookingGRPCConf  Server `json:"booking"`
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
