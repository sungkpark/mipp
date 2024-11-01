package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"mipp.com/app/service/idea"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	ideaStore := idea.NewStore(s.db)
	ideaHandler := idea.NewHandler(ideaStore)
	ideaHandler.RegisterRoutes(subrouter)

	log.Println("Listening on", s.addr)

	return http.ListenAndServe(s.addr, router)
}
