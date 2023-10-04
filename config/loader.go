package config

import (
	"errors"
	"os"
	"reflect"
	"strconv"
	"time"
)

type Config struct {
	DbHost                   string
	DbPort                   string
	DbUser                   string
	DbPassword               string
	DbName                   string
	DbSslMode                string
	HttpHost                 string
	HttpPort                 string
	HttpUrl                  string
	ClientTimeout            time.Duration
	WebSocketCompressionMode string
	WebSocketReadLimit       string
	SessionLenght            string
}

func LoadConfig() (*Config, error) {

	clientTimeoutStr := os.Getenv("CLIENT_TIMEOUT")

	clientTimeout, err := strconv.Atoi(clientTimeoutStr) // Convert string to int
	if err != nil {
		// handle error, set a default value
		clientTimeout = 5
	}
	config := Config{
		DbHost:                   os.Getenv("DB_HOST"),
		DbPort:                   os.Getenv("DB_PORT"),
		DbUser:                   os.Getenv("DB_USER"),
		DbPassword:               os.Getenv("DB_PASSWORD"),
		DbName:                   os.Getenv("DB_NAME"),
		DbSslMode:                os.Getenv("DB_SSLMODE"),
		HttpHost:                 os.Getenv("HTTP_HOST"),
		HttpPort:                 os.Getenv("HTTP_PORT_PLAIN"),
		HttpUrl:                  os.Getenv("HTTP_URL"),
		ClientTimeout:            time.Second * time.Duration(clientTimeout),
		WebSocketCompressionMode: os.Getenv("WS_COMPRESSION_MODE"),
		WebSocketReadLimit:       os.Getenv("WS_READ_LIMIT"),
		SessionLenght:            os.Getenv("SESSION_LENGHT"),
	}

	v := reflect.ValueOf(config)
	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).String() == "" {
			return nil, errors.New("missing configuration for " + v.Type().Field(i).Name)
		}
	}
	return &config, nil
}
