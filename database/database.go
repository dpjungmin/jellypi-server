package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB DB definition
var (
	DB *gorm.DB
)

// ConnectDB Create a connection to PostgreSQL and return the connection
func ConnectDB() error {
	dsn := "host=localhost user=horimz password=gorm dbname=jellypi port=5432 sslmode=disable TimeZone=Asia/Seoul"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return err
}
