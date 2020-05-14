package controller

import (
	"book-api/model"
	"book-api/views"
	"encoding/json"
	"fmt"
	"net/http"
)


func crud(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// take data out of request and save it
		fmt.Println("recieved request")
		data := views.PostRequest{}
		json.NewDecoder(r.Body).Decode(&data)
		fmt.Println(data)
		if err := model.CreateBook(data.ID, data.Name, data.Author); err != nil {
			w.Write([]byte("An error occured writing to database"))
			return
		}
		w.WriteHeader(http.StatusCreated)
		//json.NewEncoder(w).Encode(data)

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

// else if r.Method == http.MethodGet {
// 	data, err := model.ReadAll()
// 	if err != nil {
// 		w.Write([]byte(err.Error()))
// 	}
// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(data)