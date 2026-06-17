package main

import (
	"context"
	"net/http"
	"stagefront-backend/internal/config"
	"stagefront-backend/internal/handlers"
)

func main() {
	// Env loading
	cfg := config.LoadConfig()

	// Database connection
	db := config.ConnectDb()
	defer db.Close(context.Background())

	// Endpoint routes
	handlers.RegisterRoutes(db)

	// Server setup
	port := ":" + cfg.PORT
	http.ListenAndServe(port, nil)

}
