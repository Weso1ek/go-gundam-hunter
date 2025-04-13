package controller

import (
	"encoding/json"
	"net/http"
)

type SearchController struct{}

func NewSearchController() *SearchController {
	return &SearchController{}
}

func (s SearchController) Search(w http.ResponseWriter, r *http.Request) {
	var status = map[string]string{
		"status": "ok",
	}

	statusJson, err := json.Marshal(status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(statusJson)
}
