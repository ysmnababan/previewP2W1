package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"pagi/model"

	"github.com/julienschmidt/httprouter"
)

func (h *GameStoreHandler) AddBranch(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var b model.Branch

	err := json.NewDecoder(r.Body).Decode(&b)
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

	bNew, err := h.Repo.AddNewBranch(b)
	if err != nil {
		handleError(err, w)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bNew)
}
