package model

import "fmt"

//CreateBook inserts a book into a table
func CreateBook(id int, name, author string) error {
	fmt.Println("Creating book")
	insertQ, err := connection.Query("INSERT INTO `bookshop` (`id`, `name`, `author`) VALUES (? , ?, ?)", id, name, author)
	defer insertQ.Close() 
	if err!= nil {
		return err
	}
	return nil
}

// DeleteBook deletes a book from the table. Takes the book ID as an argument.
func DeleteBook(id int) error {
	insertQ, err := connection.Query("DELETE FROM BOOKSHOP WHERE id=?", id)
	defer insertQ.Close() 
	if err!= nil {
		return err
	}
	return nil
}

// UpdateBook modifies items in the table.
func UpdateBook(id int, name, author string) error {
	insertQ, err := connection.Query("UPDATE `bookshop` SET `name` = ?, `author` = ? WHERE `bookshop`.`id` = ?", name, author, id)
	defer insertQ.Close() 
	if err!= nil {
		return err
	}
	return nil
}