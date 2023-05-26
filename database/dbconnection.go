package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDBConnectionString(dbName string) string {
	return fmt.Sprint("host=localhost port=5432 user=postgres password=postgres dbname=", dbName, " sslmode=disable")
}

func OpenDBConnection(dbName string) (*gorm.DB, error) {
	connection := NewDBConnectionString(dbName)
	return gorm.Open(postgres.Open(connection))
}
