package services

import (
	"context"
	"fmt"
	"net/http"

	"github.com/jackc/pgx/v5"
)

func HandleHealthCheck(db *pgx.Conn) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var result int
		err := db.QueryRow(context.Background(), "SELECT 1").Scan(&result)

		if err != nil {
			http.Error(w, "DB connection failed", http.StatusServiceUnavailable)
			return
		}

		fmt.Fprint(w, "Server is running!")
	}
}
