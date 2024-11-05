package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/husni-robani/article-api-services/config"
	"github.com/husni-robani/article-api-services/models"
	"github.com/husni-robani/article-api-services/response"
)

func GetComment(c *gin.Context){
	article_id, err := strconv.Atoi(c.Param("article_id"))
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid Article Id", err.Error())
		return
	}

	rows, err := config.DB.Query("SELECT * FROM comments WHERE article_id = ?", article_id)

	if err != nil{
		response.Error(c, http.StatusInternalServerError, "Failed to retrieves comments", err.Error())
		return
	}

	defer rows.Close()

	var comments []models.Comment
	for rows.Next(){
		var comment models.Comment

		if err := rows.Scan(&comment.Id, &comment.ArticleId, &comment.UserId, &comment.Content, &comment.CreatedAt, &comment.UpdatedAt); err != nil{
			response.Error(c, http.StatusInternalServerError, "Failed to scan comments", err.Error())
			return
		}

		comments = append(comments, comment)
	}

	response.Success(c, http.StatusOK, comments, "Success to retrieves comments")
}

func CreateComment(c *gin.Context){
	article_id, err := strconv.Atoi(c.Param("article_id"))
	if err != nil{
		response.Error(c, http.StatusBadRequest, "Invalid Article Id", err.Error())
		return 
	}

	user_id, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User Not Authorized"})
		return 
	}

	var comment models.Comment
	comment.ArticleId = article_id
	if err := c.ShouldBindJSON(&comment); err != nil{
		response.Error(c, http.StatusBadRequest, "Failed to create comment", err.Error())
		return 
	}
	
	statement, err := config.DB.Prepare("INSERT INTO comments (article_id, user_id, content) VALUES (?, ?, ?)")
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to prepare statement", err.Error())
		return 
	}

	result, err := statement.Exec(article_id, user_id, comment.Content)
	if err != nil{
		response.Error(c, http.StatusInternalServerError, "Failed to execute statement", err.Error())
		return 
	}

	id, _ := result.LastInsertId()
	comment.Id = int(id)
	response.Success(c, http.StatusCreated, comment, "Comment Created Successfull")
}

func DeleteComment(c *gin.Context){
	comment_id, err := strconv.Atoi(c.Param("comment_id"))
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid Comment Id", err.Error())
		return
	}

	user_id, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authorized"})
		return
	}

	statement, err := config.DB.Prepare("DELETE FROM comments WHERE id = ? AND user_id = ?")
	if err != nil{
		response.Error(c, http.StatusInternalServerError, "Failed to prepare statment", err.Error())
		return 
	}

	_, err = statement.Exec(comment_id, user_id)
	if err != nil{
		response.Error(c, http.StatusInternalServerError, "Failed to execute statement", err.Error())
		return
	}

	response.Success(c, http.StatusOK, nil, "Comment Deleted Successful")
}