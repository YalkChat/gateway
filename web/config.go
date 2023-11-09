package http_server

import (
	"fmt"
	"os"
	"reflect"
)

type Config struct {
	Host string
	Port string
	Url  string
}

// TODO: There is repetition in this logic, needs to be addressed
func LoadConfig() (*Config, error) {

	config := Config{
		Host: os.Getenv("HTTP_HOST"),
		Port: os.Getenv("HTTP_PORT_PLAIN"),
		Url:  os.Getenv("HTTP_URL"),
	}
	v := reflect.ValueOf(config)
	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).String() == "" {
			return nil, fmt.Errorf("missing configuration for " + v.Type().Field(i).Name)
		}
	}
	return &config, nil
}
