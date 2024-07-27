package main

import (
	"log"

	"github.com/SamuelLFA/vaze/api/health"
	"github.com/SamuelLFA/vaze/db"
	"github.com/SamuelLFA/vaze/http"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	conn := db.StartConnection()
	defer db.CloseConnection(conn)

	healthHandler := health.New()

	http.RegisterAPIs(healthHandler)

	log.Fatal(
		http.StartServer(8080),
	)
}
