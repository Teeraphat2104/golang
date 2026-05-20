package config

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

// InitDB initializes database connection
func InitDB(dsn string) (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %w", err)
	}

	// Test the connection
	err = conn.Ping(context.Background())
	if err != nil {
		return nil, fmt.Errorf("unable to ping database: %w", err)
	}

	fmt.Println("Database connection successful")
	return conn, nil
}

// CloseDB closes database connection
func CloseDB(conn *pgx.Conn) error {
	return conn.Close(context.Background())
}
