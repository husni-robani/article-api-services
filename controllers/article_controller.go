package controllers

import (
	"net/http"
	"strconv"

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
	author_id, err := strconv.Atoi(c.Param("author_id"))
	if err != nil{
		response.Error(c, http.StatusBadRequest, "Invalid Author ID", err.Error())
		return
	}

	var new_article models.Article
	new_article.AuthorId = author_id
	if err := c.ShouldBindJSON(&new_article); err != nil{
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

func DeleteArticle(c *gin.Context){
	
	article_id, err := strconv.Atoi(c.Param("article_id"))
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid Article Id", err.Error())
		return
	}

	statement, err := config.DB.Prepare("DELETE FROM articles WHERE id = ?")
	if err != nil{
		response.Error(c, http.StatusInternalServerError, "Failed to prepare statment", err.Error())
		return
	}

	_, err = statement.Exec(article_id)
	if err != nil{
		response.Error(c, http.StatusInternalServerError, "Failed to execute statement", err.Error())
		return
	}

	response.Success(c, http.StatusOK, nil, "Article Deleted Successfull")
}

func UpdateArticle(c *gin.Context){
	article_id, err := strconv.Atoi(c.Param("article_id"))
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid Article Id", err.Error())
		return 
	}
	
	var article models.Article
	if err := c.ShouldBindBodyWithJSON(&article); err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid Input", err.Error())
		return 
	}

	statement, err := config.DB.Prepare("UPDATE articles SET title = ?, content = ?, category_id = ? WHERE id = ?")
	if err != nil{
		response.Error(c, http.StatusInternalServerError, "Failed to prepare statement", err.Error())
		return
	}

	_, err = statement.Exec(article.Title, article.Content, article.CategoryId, article_id)
	if err != nil{
		response.Error(c, http.StatusInternalServerError, "Failed to Execute Statement", err.Error())
		return
	}
	article.Id = article_id
	response.Success(c, http.StatusOK, article, "Article Updated Successfull")
}