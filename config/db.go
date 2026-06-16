package config

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

func ConnectDb() *pgx.Conn {
	cfg := LoadConfig()

	conn, err := pgx.Connect(context.Background(), cfg.DB_URL)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	return conn
}
