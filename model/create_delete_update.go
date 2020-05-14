package model

func CreateBook(id int, name, author string) error {
	insertQ, err := connection.Query("INSERT INTO `bookshop` (`id`, `name`, `author`) VALUES (? , ?, ?)", id, name, author)
	defer insertQ.Close() 
	if err!= nil {
		return err
	}
	return nil
}

func DeleteBook(id string) error {
	insertQ, err := connection.Query("DELETE FROM BOOKSHOP WHERE id=?", id)
	defer insertQ.Close() 
	if err!= nil {
		return err
	}
	return nil
}

func UpdateBook(id int, name, author string) error {
	insertQ, err := connection.Query("UPDATE `bookshop` SET `name` = ?, `author` = ? WHERE `bookshop`.`id` = ?", name, author, id)
	defer insertQ.Close() 
	if err!= nil {
		return err
	}
	return nil
}