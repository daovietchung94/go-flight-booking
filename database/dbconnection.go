package database

import (
	"fmt"
	"go-training/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func OpenDBConnection(config *config.Config, dbName string) (*gorm.DB, error) {
	connection := fmt.Sprint("host=", config.DatabaseConf.Host, " port=", config.DatabaseConf.Port, " user=postgres password=postgres dbname=", dbName, " sslmode=disable")
	return gorm.Open(postgres.Open(connection))
}
