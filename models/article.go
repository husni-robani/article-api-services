package models

type Article struct{
	Id int `json:"id"`
	Title string `json:"title"`
	Content string `json:"content"`
	AuthorId int `json:"author_id"`
	CategoryId int `json:"category_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}