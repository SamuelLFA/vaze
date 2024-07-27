package main

import (
	"log"

	"github.com/SamuelLFA/vaze/api/health"
	"github.com/SamuelLFA/vaze/http"
)

func main() {
	healthHandler := health.New()

	http.RegisterAPIs(healthHandler)

	log.Fatal(
		http.StartServer(8080),
	)
}
