package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/husni-robani/article-api-services/config"
	"github.com/husni-robani/article-api-services/models"
	"github.com/husni-robani/article-api-services/response"
)

func GetAllArticles(c *gin.Context){
	rows, err := config.DB.Query("SELECT * FROM articles")	

	if err != nil{
		response.Error(c, http.StatusInternalServerError, "Failed to get articles", err.Error())
		return
	}

	defer rows.Close()

	var articles []models.Article
	for rows.Next(){
		var article models.Article

		err := rows.Scan(&article.Id, &article.Title, &article.Content, &article.AuthorId, &article.CategoryId, &article.CreatedAt, &article.UpdatedAt)

		if err != nil{
			response.Error(c, http.StatusInternalServerError, "Failed to scan articles", err.Error())
			return
		}

		articles = append(articles, article)
	}

	response.Success(c, http.StatusOK, articles, "Success to get articles")
}

func CreateArticle(c *gin.Context){
	var new_article models.Article
	err := c.ShouldBindJSON(&new_article)
	if err != nil{
		response.Error(c, http.StatusBadRequest, "Failed to create new article", err.Error())
		return
	}

	statement, err := config.DB.Prepare("INSERT INTO articles (title, content, author_id, category_id) VALUES (?, ?, ?, ?)")
	if err != nil{
		response.Error(c, http.StatusInternalServerError, "Failed to prepare statement", err.Error())
		return
	}

	result, err := statement.Exec(new_article.Title, new_article.Content, new_article.AuthorId, new_article.CategoryId)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to execute statement", err.Error())
	}

	id, _ := result.LastInsertId()
	new_article.Id = int(id)

	response.Success(c, http.StatusCreated, new_article, "Article created successfully!!")
}