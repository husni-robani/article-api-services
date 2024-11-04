package models

type Comment struct{
	Id int `json:"id"`
	ArticleId int `json:"article_id,omitempty"`
	UserId int `json:"user_id,omitempty"`
	Content string `json:"content"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}
