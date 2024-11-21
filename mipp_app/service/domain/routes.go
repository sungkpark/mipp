package domain

import (
	// "net/http"

	// "github.com/gorilla/mux"
	"mipp.com/app/types"
)

type Handler struct {
	store types.DomainStore
}

func NewHandler(store types.DomainStore) *Handler {
	return &Handler{store: store}
}

// func (h *Handler) RegisterRoutes(router *mux.Router) {
// 	router.HandleFunc("/domains/{domainName}", h.handleGetDomainByName).Methods(http.MethodGet)
// 	router.HandleFunc("/domains", h.handleGetDomains).Methods(http.MethodGet)

// NO NEED
// 	router.HandleFunc("/domains", h.handleCreateDomain).Methods(http.MethodPost)
// }