package handlers

import (
	"net/http"
	"stagefront-backend/internal/services"

	"github.com/jackc/pgx/v5"
)

func registerAuthRoutes(db *pgx.Conn) {
	http.HandleFunc("POST /auth/register", services.RegisterHandler(db))
}
