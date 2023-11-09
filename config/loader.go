package config

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"time"
)

type Config struct {
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string
	DbSslMode  string
	// HttpHost                 string
	// HttpPort                 string
	// HttpUrl                  string
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
		ClientTimeout:            time.Second * time.Duration(clientTimeout),
		WebSocketCompressionMode: os.Getenv("WS_COMPRESSION_MODE"),
		WebSocketReadLimit:       os.Getenv("WS_READ_LIMIT"),
		SessionLenght:            os.Getenv("SESSION_LENGHT"),
	}

	v := reflect.ValueOf(config)
	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).String() == "" {
			return nil, fmt.Errorf("missing configuration for " + v.Type().Field(i).Name)
		}
	}
	return &config, nil
}
