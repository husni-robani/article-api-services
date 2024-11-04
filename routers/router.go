package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/husni-robani/article-api-services/controllers"
	"github.com/husni-robani/article-api-services/middleware"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")
	{
		authenticatedRoute := api.Group("")
		authenticatedRoute.Use(middleware.AuthMiddleware())
		{
			authenticatedRoute.POST("/articles", controllers.CreateArticle)
		}
		
		api.POST("/register", controllers.Register)
        api.POST("/login", controllers.Login)
		api.GET("/articles", controllers.GetAllArticles)
	}
	return router
}