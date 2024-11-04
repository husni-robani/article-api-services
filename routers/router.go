package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/husni-robani/article-api-services/controllers"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/api/articles", controllers.GetAllArticles)

	return router
}