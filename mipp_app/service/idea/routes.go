package idea

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/sikozonpc/ecom/utils"
	"mipp.com/app/types"
)

type Handler struct {
	ideaStore   types.IdeaStore
	domainStore types.DomainStore
}

func NewHandler(ideaStore types.IdeaStore, domainStore types.DomainStore) *Handler {
	return &Handler{ideaStore: ideaStore, domainStore: domainStore}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/ideas/offset/{offset}/limit/{limit}", h.handleGetIdeas).Methods(http.MethodGet)
	router.HandleFunc("/ideas/{ideaID}", h.handleGetIdeaByID).Methods(http.MethodGet)
	router.HandleFunc("/ideas/domain-id/{domainID}", h.handleGetIdeasByDomainID).Methods(http.MethodGet)
	router.HandleFunc("/ideas", h.handleGetIdeasByID).Methods(http.MethodGet)

	router.HandleFunc("/ideas", h.handleCreateIdea).Methods(http.MethodPost)
}

func (h *Handler) handleGetIdeas(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	offset, err := strconv.Atoi(vars["offset"])
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("missing offset"))
	}
	limit, err := strconv.Atoi(vars["limit"])
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("missing limit"))
	}
	ideas, err := h.ideaStore.GetIdeas(offset, limit)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, ideas)
}

func (h *Handler) handleGetIdeaByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	str, ok := vars["ideaID"]

	if !ok {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("missing idea ID"))
		return
	}

	ideaID, err := strconv.Atoi(str)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid idea ID"))
		return
	}

	idea, err := h.ideaStore.GetIdeaByID(ideaID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, idea)
}

func (h *Handler) handleGetIdeasByDomainID(w http.ResponseWriter, r *http.Request) {
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

	ideas, err := h.ideaStore.GetIdeasByDomainID(domainID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, ideas)
}

func (h *Handler) handleGetIdeasByID(w http.ResponseWriter, r *http.Request) {
	var ideaIDs types.Ideas
	if err := utils.ParseJSON(r, &ideaIDs); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	if err := utils.Validate.Struct(ideaIDs); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	ideaIDsParsed, err := sliceAtoi(ideaIDs.IDs)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid idea ID"))
		return
	}

	idea, err := h.ideaStore.GetIdeasByID(ideaIDsParsed)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, idea)
}

func (h *Handler) handleCreateIdea(w http.ResponseWriter, r *http.Request) {
	var idea types.CreateIdeaPayload
	if err := utils.ParseJSON(r, &idea); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	if err := utils.Validate.Struct(idea); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	domain, err := h.domainStore.GetDomainByName(idea.DomainName)
	if err != nil {
		return
	}

	err = h.ideaStore.CreateIdea(idea, domain.ID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, idea)
}

func sliceAtoi(sa []string) ([]int, error) {
	si := make([]int, 0, len(sa))
	for _, a := range sa {
		i, err := strconv.Atoi(a)
		if err != nil {
			return si, err
		}
		si = append(si, i)
	}
	return si, nil
}
