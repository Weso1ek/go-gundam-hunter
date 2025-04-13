package main

import (
	"fmt"
	"github.com/Weso1ek/go-gundam-hunter/controller"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	searchController := controller.NewSearchController()

	mux.HandleFunc("/search", searchController.Search)

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mux.ServeHTTP(w, r)
	})

	fmt.Println("Server Start")
	log.Fatal(http.ListenAndServe(":8099", handler))
}
