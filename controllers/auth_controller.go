package controllers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/husni-robani/article-api-services/config"
	"github.com/husni-robani/article-api-services/models"
	"github.com/husni-robani/article-api-services/response"
	"github.com/husni-robani/article-api-services/utils"
)

func Register(c *gin.Context) {
    var user models.User
    if err := c.ShouldBindJSON(&user); err != nil {
        response.Error(c, http.StatusBadRequest, "Invalid input", err.Error())
        return
    }

    statement, err := config.DB.Prepare("INSERT INTO users (username, email, password) VALUES (?, ?, ?)")
    if err != nil {
        response.Error(c, http.StatusInternalServerError, "Failed to prepare statement", err.Error())
        return
    }
    result, err := statement.Exec(user.Username, user.Email, user.Password)
    if err != nil {
        response.Error(c, http.StatusInternalServerError, "Failed to execute statement", err.Error())
        return
    }
    id, _ := result.LastInsertId()
    user.Id = int(id)

    response.Success(c, http.StatusCreated, user, "User registered successfully")
}

func Login(c *gin.Context) {
    var input struct {
        Email    string `json:"email"`
        Password string `json:"password"`
    }
    if err := c.ShouldBindJSON(&input); err != nil {
        response.Error(c, http.StatusBadRequest, "Invalid input", err.Error())
        return
    }

    var user models.User
    row := config.DB.QueryRow("SELECT id, username, email, password FROM users WHERE email = ?", input.Email)
    err := row.Scan(&user.Id, &user.Username, &user.Email, &user.Password)
    if err != nil {
        if err == sql.ErrNoRows {
            response.Error(c, http.StatusUnauthorized, "User not found", "")
        } else {
            response.Error(c, http.StatusInternalServerError, "Failed to fetch user", err.Error())
        }
        return
    }

    if user.Password != input.Password {
        response.Error(c, http.StatusUnauthorized, "Invalid password", "")
        return
    }

    token, err := utils.GenerateJWT(user.Id)
    if err != nil {
        response.Error(c, http.StatusInternalServerError, "Failed to generate token", err.Error())
        return
    }

    response.Success(c, http.StatusOK, gin.H{"token": token}, "Login successful")
}