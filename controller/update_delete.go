package controller

import (
	"book-api/model"
	"book-api/views"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func updateDel(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodDelete {
		fmt.Println("delete request recieved")
		id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/api/v1/books/"))
		if err != nil {
				log.Fatal(err)
			} 
		fmt.Println(id)
		if err := model.DeleteBook(id); err != nil {
			w.Write([]byte("An error occurred while deleting from database"))
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(struct {
			Status string `json:"status"`
		}{"Item Deleted"})
		
	}else if r.Method == http.MethodPut {
		fmt.Println("Put request recieved on", r.URL.String())
		id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/api/v1/books/"))
		if err != nil {
				log.Fatal(err)
			} 
		fmt.Println(id)
		data := views.UpdateRequest{}
		json.NewDecoder(r.Body).Decode(&data)
		fmt.Println(data)
			if err := model.UpdateBook(id, data.Name, data.Author); err != nil {
				w.Write([]byte("Some error"))
				return
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(struct {
				Status string `json:"status"`
			}{"Book Updated"})
		}
}