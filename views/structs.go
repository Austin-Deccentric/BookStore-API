package views

// PostRequest implements the structure of the database.
//It provides the scheme for making POST requests to the database
type PostRequest struct {
	ID int			`json:"id"`
	Name string 	`json:"name"`
	Author string		`json:"author"`
	PublishedAt string `json:"published_at"`
}

// UpdateRequest provides the structure for deserializing JSON objects in a PUT request to the database.
type UpdateRequest struct {
	Name string 	`json:"name"`
	Author string		`json:"author"`
}
