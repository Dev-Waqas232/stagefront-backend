package main

import (
	"fmt"
	"net/http"
	"stagefront-backend/config"
)

func main() {
	cfg := config.LoadConfig()

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Server is running!")
	})

	port := ":" + cfg.PORT
	http.ListenAndServe(port, nil)
}
