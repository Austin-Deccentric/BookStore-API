package controller

import "net/http"

func Register() *http.ServeMux {
	mux:= http.NewServeMux()
	mux.HandleFunc("/api/v1/books", crud)
	mux.HandleFunc("/api/v1/books/{id}", updateDel)
	return mux
}