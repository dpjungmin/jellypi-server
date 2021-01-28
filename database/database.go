package database

import (
	"github.com/dpjungmin/jellypi-server/entities"
	"github.com/dpjungmin/jellypi-server/tools/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// DB is the underlying database connection
var DB *gorm.DB

// Connect initiates the database connection
func Connect() {
	cfg := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	}

	dsn := "host=localhost user=horimz password=passwd dbname=jellypi port=5432 sslmode=disable TimeZone=Asia/Seoul"
	db, err := gorm.Open(postgres.Open(dsn), cfg)

	if err != nil {
		logger.Error("[DATABASE]::CONNECTION_ERROR", err)
		panic(err)
	}

	DB = db
	logger.Info("Database connected")
}

// Close the database connection
func Close() {
	DB.Statement.ReflectValue.Close()
}

// AutoMigrate migrates all the database tables
func AutoMigrate() {
	var tables []interface{}

	tables = append(tables, &entities.User{})

	if err := autoMigrate(tables...); err != nil {
		logger.Error("[DATABASE]::MIGRATION_ERROR", err)
		panic(err)
	}

	logger.Info("Automatically migrated tables")
}

func autoMigrate(tables ...interface{}) error {
	return DB.AutoMigrate(tables...)
}
