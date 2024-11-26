package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"mipp.com/app/service/domain"
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

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	domainStore := domain.NewStore(s.db)
	domainHandler := domain.NewHandler(domainStore)
	domainHandler.RegisterRoutes(subrouter)

	ideaStore := idea.NewStore(s.db)
	ideaHandler := idea.NewHandler(ideaStore, domainStore)
	ideaHandler.RegisterRoutes(subrouter)

	log.Println("Listening on", s.addr)

	return http.ListenAndServe(s.addr, c.Handler(router))
}
