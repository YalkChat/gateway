package main

import (
	"os"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	// Set up mock environment variable
	os.Setenv("DB_HOST", "test_host")
	// ..set other environment variables ..

	config, err := loadConfig()
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}

	if config.DbHost != "test_host" {
		t.Errorf("Expected DB_HOST to be 'test_host', got %s", config.DbHost)
	}
}
