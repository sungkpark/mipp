package domain

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sikozonpc/ecom/utils"
	"mipp.com/app/types"
)

type Handler struct {
	store types.DomainStore
}

func NewHandler(store types.DomainStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/domains/limit/{limit}", h.handleGetDomains).Methods(http.MethodGet)
	router.HandleFunc("/domains/{domainID}", h.handleGetDomainByID).Methods(http.MethodGet)

	// NO NEED
	// router.HandleFunc("/domains", h.handleCreateDomain).Methods(http.MethodPost)
}

func (h *Handler) handleGetDomains(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	limit, err := strconv.Atoi(vars["limit"])

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("missing limit"))
	}
	domains, err := h.store.GetDomains(limit)
	// print(&domains[0].ID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, domains)
}

func (h *Handler) handleGetDomainByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	str, ok := vars["domainID"]

	if !ok {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("missing domain ID"))
		return
	}

	domainID, err := strconv.Atoi(str)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid domain ID"))
		return
	}

	domain, err := h.store.GetDomainByID(domainID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, domain)
}
