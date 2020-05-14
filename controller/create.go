package controller

import (
	"book-api/model"
	"encoding/json"
	"fmt"
	"net/http"
)


func create(w http.ResponseWriter, r *http.Request,) {
	if r.Method == http.MethodPost {
		// take data out of request and save it
		fmt.Println("recieved POST request")
		if err := model.CreateBook(value.ID, value.Name, value.Author); err != nil {
			w.Write([]byte("An error occured writing to database"))
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(value)

	} else if r.Method == http.MethodGet {
			books, err := model.ReadAll()
			if err != nil {
				w.Write([]byte(err.Error()))
				return
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(books)
	} 
}

