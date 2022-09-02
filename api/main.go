package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	"sprit3dan.dev/weather-reports/api/services"
)

func prepareDB(dbPath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	f, err := os.Open("db/create_bundle.sql")
	if err != nil {
		return nil, err
	}

	barr := make([]byte, 1024*4)
	f.Read(barr)
	bundle := string(barr)

	_, dbErr := db.Exec(bundle)
	if dbErr != nil {
		return db, dbErr
	}

	return db, nil
}

func main() {
	var db *sql.DB
	var err error

	// Handle a retry for local DB
	if db, err = prepareDB("/app/weather-reports.db"); err != nil {
		if db, err = prepareDB("../weather-reports.db"); err != nil {
			log.Fatal(err)
		}
	}

	mux := mux.NewRouter()

	services.NewHealthService(db, mux)

	log.Fatal(http.ListenAndServe(":8080", mux))
}
