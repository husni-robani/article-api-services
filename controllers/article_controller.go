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