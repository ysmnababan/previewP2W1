package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"pagi/model"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (h *GameStoreHandler) UpdateBranch(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	param_id := p.ByName("id")

	id, err := strconv.Atoi(param_id)
	if err != nil || id <= 0 {
		log.Println("invalid id")
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var b model.Branch

	err = json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		log.Println("error decoding data")
		http.Error(w, "error getting the data", http.StatusBadRequest)
		return
	}

	if b.Name == "" || b.Location == "" {
		log.Println("error or missing parameter")
		http.Error(w, "error or missing parameter", http.StatusBadRequest)
		return
	}

	err = h.Repo.UpdateBranch(id, b)
	if err != nil {
		handleError(err, w)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(model.BranchMessage{Message: "Branch updated successfully"})
}
