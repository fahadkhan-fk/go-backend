package postgres

import (
	"fmt"
	"log"

	"github.com/fahadkhan-fk/go-gin-backend/pkg/env"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	dbClient *gorm.DB
	err      error
)

// DBConfig stores database related configurations
type DBConfig struct {
	Host     string
	UserName string
	Password string
	DbName   string
	SSLMode  string
	Port     string
}

func EstablishDBConnection() {
	dbClient, err = gorm.Open(postgres.Open(getDSN()), &gorm.Config{})
	if err != nil {
		log.Println("Connection Failed to Open due to: ", err)
	} else {
		log.Println("Connection Established !!")
	}

}

func getDSN() string {
	dsn := env.GetConnectionString()
	if dsn == "" {
		config := getDBConfig()
		return fmt.Sprintf("host=%s user=%s dbname=%s port=%s sslmode=%s password=%s", config.Host, config.UserName, config.DbName, config.Port, config.SSLMode, config.Password)
	}
	return dsn
}

func getDBConfig() DBConfig {
	config := DBConfig{
		Host:     env.GetHost(),
		UserName: env.GetDbUser(),
		Password: env.GetDbPassword(),
		DbName:   env.GetDbName(),
		SSLMode:  env.GetSSLMode(),
		Port:     env.GetPort(),
	}
	return config
}

//GetDBClient returns a db client
func GetClient() *gorm.DB {
	if dbClient == nil {
		establishDBConnection()
		return dbClient
	} else {
		return dbClient
	}
}
