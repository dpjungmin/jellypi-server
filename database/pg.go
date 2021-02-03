package database

import (
	"fmt"
	"sync"

	"github.com/dpjungmin/jellypi-server/config"
	"github.com/dpjungmin/jellypi-server/domain"
	"github.com/dpjungmin/jellypi-server/utils/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// PG is a struct that holds the PostgreSQL client
type PG struct {
	client *gorm.DB
}

var (
	once     sync.Once
	instance *PG // instance is the underlying database connection
)

// GetPGSingleton returns the singleton instance of PG
func GetPGSingleton() *PG {
	once.Do(func() {
		instance = &PG{client: connect()}
	})
	return instance
}

func connect() *gorm.DB {
	cfg := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	}

	host := config.PG.Host
	user := config.PG.User
	pass := config.PG.Password
	dbName := config.PG.DBName
	port := config.PG.Port

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Seoul", host, user, pass, dbName, port)
	db, err := gorm.Open(postgres.Open(dsn), cfg)

	if err != nil {
		logger.Error("[DATABASE]::CONNECTION_ERROR", err)
		panic(err)
	}

	logger.Info("Successfully connected to database")
	return db
}

// Client returns the database client
func (db *PG) Client() *gorm.DB {
	return db.client
}

// AutoMigrate will migrate all the database tables that are provided
func (db *PG) AutoMigrate() {
	tables := []interface{}{
		&domain.User{},
	}

	if err := db.client.AutoMigrate(tables...); err != nil {
		logger.Error("[DATABASE]::MIGRATION_ERROR", err)
		panic(err)
	}

	logger.Info("Automatically migrated tables")
}
