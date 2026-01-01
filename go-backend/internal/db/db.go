package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var Pool *pgxpool.Pool

func Connect() error {
	dbUrl := os.Getenv("DATABASE_URL")
	if dbUrl == "" {
		// Default to the docker-compose settings if not set
		dbUrl = "postgres://postgres:password@localhost:5432/hairone_db"
	}

	config, err := pgxpool.ParseConfig(dbUrl)
	if err != nil {
		return fmt.Errorf("unable to parse database config: %w", err)
	}

	Pool, err = pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return fmt.Errorf("unable to create connection pool: %w", err)
	}

	if err := Pool.Ping(context.Background()); err != nil {
		return fmt.Errorf("unable to connect to database: %w", err)
	}

	fmt.Println("Connected to PostgreSQL!")
	return nil
}

func Migrate(schemaPath string) error {
	schema, err := os.ReadFile(schemaPath)
	if err != nil {
		return fmt.Errorf("failed to read schema file: %w", err)
	}

	_, err = Pool.Exec(context.Background(), string(schema))
	if err != nil {
		return fmt.Errorf("failed to execute migration: %w", err)
	}

	fmt.Println("Database migration applied successfully.")
	return nil
}
