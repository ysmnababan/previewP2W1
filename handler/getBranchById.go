package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (h *GameStoreHandler) GetBranchBYId(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	param_id := p.ByName("id")

	id, err := strconv.Atoi(param_id)
	if err != nil || id <= 0 {
		log.Println("invalid id")
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	b, err := h.Repo.GetBranchByID(id)
	if err != nil {
		handleError(err, w)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(b)
}
