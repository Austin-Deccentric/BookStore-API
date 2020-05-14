package controller

import "net/http"

// Register registers endpoints and does the routing.
func Register() *http.ServeMux {
	mux:= http.NewServeMux()
	mux.HandleFunc("/api/v1/books", create)
	mux.HandleFunc("/api/v1/books/", updateDel)
	return mux
}