package services

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
)

type healthService struct {
	db *sql.DB
}

func NewHealthService(db *sql.DB, mux *mux.Router) *healthService {
	hs := &healthService{db}

	mux.HandleFunc("/api/health", hs.healthCheck)

	return hs
}

func (hs *healthService) healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
