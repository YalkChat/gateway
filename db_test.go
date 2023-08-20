package main

import (
	"fmt"
	"testing"
)

func TestOpenDbConnection(t *testing.T) {
	// Use a test configuration or mock environment variables
	config := &Config{
		DbHost:     "192.168.188.34",
		DbPort:     "5432",
		DbUser:     "postgres",
		DbPassword: "yourpostgrespass",
		DbName:     "yalk",
		DbSslMode:  "disable",
	}

	db, err := openDbConnection(config)
	if err != nil {
		t.Fatalf("Failed to open DB connection: %v", err)
	}
	defer func() {
		sqlDB, err := db.DB()
		if err != nil {
			fmt.Printf("Error getting db interface")
		}
		sqlDB.Close()
	}()

	// Additional checks, e.g., ping the database to ensure the connection is alive

	sqlDB, err := db.DB()
	if err != nil {
		fmt.Printf("Error getting db interface")
	}

	if err := sqlDB.Ping(); err != nil {
		t.Errorf("Failed to ping DB: %v", err)
	}

}
