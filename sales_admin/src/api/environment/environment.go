package environment

import (
	"github.com/sirupsen/logrus"
	"os"
)

type Environment struct {
	DbUser               string
	DbPass               string
	DbName               string
	DbHost               string
	DbPort               string
	DbMaxLifetime        string
	DbMaxOpenConnections string
	DbMaxIdleConnections string
	DbMaxAttempts        string
	AppPort              string
	AppHost              string
	ApiPort              string
	LogLevel             string
	JwtSecret            string
}

func NewEnvironment() *Environment {
	var logLevel logrus.Level
	switch os.Getenv("LOG_LEVEL") {
	case "info":
		logLevel = logrus.InfoLevel
	case "debug":
		logLevel = logrus.DebugLevel
	case "fatal":
		logLevel = logrus.FatalLevel
	default:
		logLevel = logrus.FatalLevel
	}

	logrus.SetLevel(logLevel)

	return &Environment{
		DbUser:               os.Getenv("DB_USER"),
		DbPass:               os.Getenv("DB_PASS"),
		DbName:               os.Getenv("DB_NAME"),
		DbHost:               os.Getenv("DB_HOST"),
		DbPort:               os.Getenv("DB_PORT"),
		AppPort:              os.Getenv("APP_PORT"),
		AppHost:              os.Getenv("APP_HOST"),
		ApiPort:              os.Getenv("API_PORT"),
		LogLevel:             os.Getenv("LOG_LEVEL"),
		DbMaxLifetime:        os.Getenv("DB_MAX_LIFETIME"),
		DbMaxOpenConnections: os.Getenv("DB_MAX_OPEN_CONNECTIONS"),
		DbMaxIdleConnections: os.Getenv("DB_MAX_IDLE_CONNECTIONS"),
		DbMaxAttempts:        os.Getenv("DB_MAX_ATTEMPTS"),
		JwtSecret:            os.Getenv("JWT_SECRET"),
	}
}
