package db

import (
	"database/sql"
	"log"
	"os"
)

func StartConnection() *sql.DB {
	db, err := sql.Open("sqlite3", "./db/sqlite.db")
	if err != nil {
		log.Fatal(err)
	}

	runMigrations(db)

	return db
}

func CloseConnection(db *sql.DB) {
	db.Close()
}

func runMigrations(db *sql.DB) {
	initFileBytes, err := os.ReadFile("./db/sql/init.sql")
	if err != nil {
		log.Fatalf("fail to read init file: %s", err)
	}

	initScript := string(initFileBytes[:])
	_, err = db.Exec(initScript)
	if err != nil {
		log.Fatalf("fail to execute init sql: %s", err)
	}
}
