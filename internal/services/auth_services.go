package services

import (
	"encoding/json"
	"net/http"

	"github.com/jackc/pgx/v5"
)

type RegisterBody struct {
	first
}

func RegisterHandler(db *pgx.Conn) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		err := json.NewDecoder(r.Body).Decode()
	}
}
