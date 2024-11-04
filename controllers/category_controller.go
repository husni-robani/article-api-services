package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/husni-robani/article-api-services/config"
	"github.com/husni-robani/article-api-services/models"
	"github.com/husni-robani/article-api-services/response"
)

func CreateCategory(c *gin.Context){
	var category models.Category
	err := c.ShouldBindJSON(&category)
	if err != nil{
		response.Error(c, http.StatusBadRequest, "Failed to create new category", err.Error())
		return
	}

	statement, err := config.DB.Prepare("INSERT INTO categories (name, description) VALUES (?, ?)")
	if err != nil{
		response.Error(c, http.StatusInternalServerError, "Failed to prepare statement", err.Error())
		return
	}

	result, err := statement.Exec(category.Name, category.Description)
	if err != nil{
		response.Error(c, http.StatusInternalServerError, "Failed to execute statement", err.Error())
		return 
	}

	id, _ := result.LastInsertId()
	category.Id = int(id)
	response.Success(c, http.StatusCreated, category, "Category Created Successfully")
}