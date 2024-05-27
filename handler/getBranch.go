package handler

import (
	"errors"
	"log"
	"net/http"
	"pagi/repo"

	"github.com/julienschmidt/httprouter"
)

type GameStoreHandler struct {
	Repo repo.GameStoreRepo
}

func handleError(err error, w http.ResponseWriter) {
	log.Println(err)

	switch {
	case errors.Is(err, repo.ErrQuery):
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	case errors.Is(err, repo.ErrScan):
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	case errors.Is(err, repo.ErrRowsAffected):
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	case errors.Is(err, repo.ErrLastInsertId):
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	case errors.Is(err, repo.ErrNoAffectedRow):
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	case errors.Is(err, repo.ErrNoRows):
		http.Error(w, "No customers found", http.StatusNotFound)
	case errors.Is(err, repo.ErrInvalidId):
		http.Error(w, "Invalid ID", http.StatusBadRequest)
	case errors.Is(err, repo.ErrUserExists):
		http.Error(w, "User Already Exists", http.StatusBadRequest)
	default:
		http.Error(w, "Unknown error", http.StatusInternalServerError)
	}
}

func (h *GameStoreHandler) GetBranches(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}
