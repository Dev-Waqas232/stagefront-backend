package main

import (
	"context"
	"fmt"
	"net/http"
	"stagefront-backend/config"
)

func main() {
	cfg := config.LoadConfig()

	db := config.ConnectDb()
	defer db.Close(context.Background())

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		var result int
		err := db.QueryRow(context.Background(), "SELECT 1").Scan(&result)

		if err != nil {
			http.Error(w, "DB connection failed", http.StatusServiceUnavailable)
			return
		}

		fmt.Fprint(w, "Server is running!")
	})

	port := ":" + cfg.PORT
	http.ListenAndServe(port, nil)

}
