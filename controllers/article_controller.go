package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/husni-robani/article-api-services/config"
	"github.com/husni-robani/article-api-services/models"
)

func GetAllArticles(c *gin.Context){
	rows, err := config.DB.Query("SELECT * FROM articles")	

	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	defer rows.Close()

	var articles []models.Article
	for rows.Next(){
		var article models.Article

		err := rows.Scan(&article.Id, &article.Title, &article.Content, &article.AuthorId, &article.CategoryId, &article.CreatedAt, &article.UpdatedAt)

		if err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		articles = append(articles, article)
	}

	c.JSON(http.StatusOK, articles)
}