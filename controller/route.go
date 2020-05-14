package controller

import (
	
	"book-api/views"
	"encoding/json"
	"fmt"
	"net/http"
)

var value views.PostRequest

func check(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		
		if r.Method == http.MethodPost {
			fmt.Println("Checking data")
			data:= views.PostRequest{}
			defer r.Body.Close()
			json.NewDecoder(r.Body).Decode(&data)
			fmt.Println(data)
			if &data.ID == nil || data.Name == "" || data.Author == "" {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("A value is missing"))
				return 
			}
			value = data
		}
		next(w, r)
	}
}


// Register registers endpoints and does the routing.
func Register() *http.ServeMux {
	mux:= http.NewServeMux()
	mux.HandleFunc("/api/v1/books",check(create))
	mux.HandleFunc("/api/v1/books/", updateDel)
	return mux
}