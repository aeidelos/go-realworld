package models

type Article struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Content     string   `json:"content"`
	Body        string   `json:"body"`
	Taglist     []string `json:"taglist"`
}
