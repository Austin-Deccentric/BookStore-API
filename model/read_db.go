package model

import "book-api/views"

// ReadAll displays all data in the table
func ReadAll() ([]views.PostRequest, error) {
	rows, err := connection.Query("SELECT * FROM BOOKSHOP")
	if err != nil {
		return nil, err
	}
	books := []views.PostRequest{}
	for rows.Next() {
		libray := views.PostRequest{}
		rows.Scan(&libray.ID, &libray.Name, &libray.Author, &libray.PublishedAt)
		books = append(books, libray)
	}
	return books, nil
}

