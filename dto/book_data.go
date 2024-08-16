package dto

type BookData struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Isbn        string `json:"isbn"`
}

type CreateBookRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Isbn        string `json:"isbn"`
}

type UpdateBookRequest struct {
	Id          string `json:"-"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Isbn        string `json:"isbn"`
}
