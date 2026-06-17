package handlers

import "github.com/jackc/pgx/v5"

func RegisterRoutes(db *pgx.Conn) {
	registerAuthRoutes(db)
	registerHealthRoutes(db)
}
