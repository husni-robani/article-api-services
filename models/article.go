package models

type Article struct{
	Id int `json:"id"`
	Title string `json:"title"`
	Content string `json:"content"`
	AuthorId int `json:"author_id,omitempty"`
	CategoryId int `json:"category_id"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}