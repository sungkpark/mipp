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
	store types.IdeaStore
}

func NewHandler(store types.IdeaStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/ideas/offset/{offset}/limit/{limit}", h.handleGetIdeas).Methods(http.MethodGet)
	router.HandleFunc("/ideas/{ideaID}", h.handleGetIdeaByID).Methods(http.MethodGet)
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
	ideas, err := h.store.GetIdeas(offset, limit)
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

	idea, err := h.store.GetIdeaByID(ideaID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, idea)
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

	idea, err := h.store.GetIdeasByID(ideaIDsParsed)
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

	err := h.store.CreateIdea(idea)
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
