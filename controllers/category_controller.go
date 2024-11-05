package controllers

import (
	"net/http"
	"strconv"

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

func DeleteCategory(c *gin.Context){
	category_id, err := strconv.Atoi(c.Param("category_id"))
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid category id", err.Error())
		return
	}

	statement, err := config.DB.Prepare("DELETE FROM categories WHERE id = ?")
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to prepare statmenet", err.Error())
		return
	}

	_, err = statement.Exec(category_id)
	if err != nil{
		response.Error(c, http.StatusInternalServerError, "Failed to execute statement", err.Error())
		return
	}

	response.Success(c, http.StatusOK, nil, "Deleted category successfull")
}