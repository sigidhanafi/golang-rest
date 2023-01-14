package book

type BookResponse struct {
	ID    int    `json: "id"`
	Title string `json: "title"`
}
