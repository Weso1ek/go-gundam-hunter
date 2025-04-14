package controller

import (
	"encoding/json"
	"github.com/Weso1ek/go-gundam-hunter/service"
	"net/http"
)

type SearchController struct{}

func NewSearchController() *SearchController {
	return &SearchController{}
}

func (s SearchController) Search(w http.ResponseWriter, r *http.Request) {
	searchTerm := r.URL.Query().Get("searchTerm")

	offers := service.GetOffers(searchTerm)

	statusJson, err := json.Marshal(offers)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(statusJson)
}
