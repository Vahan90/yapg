package main

import (
	"log"
	"net/http"
	"os"

	"github.com/vahan90/yapg/internal/redisstore"
)

func main() {
	redisstore.InitializeRedis()

	http.HandleFunc("/push", redisstore.PushHandler)
	http.HandleFunc("/metrics", redisstore.MetricsHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "9091"
	}

	log.Printf("Server listening on port %s...", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
