package controller

import (
	"book-api/model"
	"book-api/views"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func updateDel(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodDelete {
		fmt.Println("delete request recieved")
			id := r.URL.Path[2:]
			fmt.Println(id)
			if err := model.DeleteBook(id); err != nil {
				w.Write([]byte("Some error"))
				return
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(struct {
				Status string `json:"status"`
			}{"Item Deleted"})
		
	}else if r.Method == http.MethodPut {
		fmt.Println("Put request recieved", r.URL.Path)
		keys, ok :=r.URL.Query()["key"]
		if !ok || len(keys[0]) < 1 {
			log.Println("Url Pa")
		}
		id, err := strconv.Atoi(r.URL.Path[1:])
		if err != nil {
			log.Fatal(err)
		}
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