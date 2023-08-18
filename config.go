package main

import (
	"errors"
	"os"
	"reflect"
)

type Config struct {
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string
	DbSslMode  string
	HttpHost   string
	HttpPort   string
	HttpUrl    string
}

func loadConfig() (*Config, error) {

	config := Config{
		DbHost:     os.Getenv("DB_HOST"),
		DbPort:     os.Getenv("DB_PORT"),
		DbUser:     os.Getenv("DB_USER"),
		DbPassword: os.Getenv("DB_PASSWORD"),
		DbName:     os.Getenv("DB_NAME"),
		DbSslMode:  os.Getenv("DB_SSLMODE"),
		HttpHost:   os.Getenv("HTTP_HOST"),
		HttpPort:   os.Getenv("HTTP_PORT_PLAIN"),
		HttpUrl:    os.Getenv("HTTP_URL"),
	}

	v := reflect.ValueOf(config)
	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).String() == "" {
			return nil, errors.New("missing configuration for " + v.Type().Field(i).Name)
		}
	}
	return &config, nil
}
