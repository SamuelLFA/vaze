package main

import (
	"log"

	"github.com/SamuelLFA/vaze/api/health"
	"github.com/SamuelLFA/vaze/api/user"
	"github.com/SamuelLFA/vaze/db"
	"github.com/SamuelLFA/vaze/http"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	conn := db.StartConnection()
	defer db.CloseConnection(conn)

	healthHandler := health.NewHandler()
	userHandler := user.NewHandler()

	http.RegisterAPIs(healthHandler, userHandler)

	log.Fatal(
		http.StartServer(8080),
	)
}
