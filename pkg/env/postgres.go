package env

import (
	"os"

	log "github.com/sirupsen/logrus"
)

const (
	HOST        = "localhost"
	PORT        = "5432"
	DB_USER     = "postgres"
	DB_PASSWORD = "postgres"
	DB_NAME     = "go-backend"
	SSLMODE     = "disable"
	DATABASE    = "postgres"
)

func GetDatabaseType() string {
	database := os.Getenv("DATABASE")
	if database == "" {
		database = DATABASE
	}
	return database
}

func GetDbUser() string {
	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		dbUser = DB_USER
	}
	return dbUser
}

func GetDbPassword() string {
	dbPassword := os.Getenv("DB_PASSWORD")
	if dbPassword == "" {
		dbPassword = DB_PASSWORD
	}
	return dbPassword
}

func GetDbName() string {
	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = DB_NAME
	}
	return dbName
}

func GetSSLMode() string {
	sslmode := os.Getenv("SSLMODE")
	if sslmode == "" {
		sslmode = SSLMODE
	}
	return sslmode
}

func GetConnectionString() string {
	dsn := os.Getenv("DSN")
	if dsn == "" {
		log.Warn("Database connection string doesn't exist !!")
		return dsn
	}
	return dsn
}

func GetHost() string {
	host := os.Getenv("HOST")
	if host == "" {
		host = HOST
	}
	return host
}

func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = PORT
	}
	return port
}
