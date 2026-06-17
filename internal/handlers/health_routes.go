package handlers

import (
	"net/http"

	"stagefront-backend/internal/services"

	"github.com/jackc/pgx/v5"
)

func registerHealthRoutes(db *pgx.Conn) {
	http.HandleFunc("GET /health", services.HandleHealthCheck(db))
}
