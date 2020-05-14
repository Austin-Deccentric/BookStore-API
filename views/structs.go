package views

type PostRequest struct {
	ID int			`json:"id"`
	Name string 	`json:"name"`
	Author string		`json:"author"`
	PublishedAt string `json:"published_at"`
}

type UpdateRequest struct {
	Name string 	`json:"name"`
	Author string		`json:"author"`
}
