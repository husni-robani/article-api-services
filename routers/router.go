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
			authenticatedRoute.POST("/articles/:author_id", controllers.CreateArticle)
			authenticatedRoute.DELETE("/articles/:article_id", controllers.DeleteArticle)
			authenticatedRoute.PUT("/articles/:article_id", controllers.UpdateArticle)

			authenticatedRoute.POST("/comments/:article_id", controllers.CreateComment)
			authenticatedRoute.DELETE("/comments/:comment_id", controllers.DeleteComment)

			authenticatedRoute.POST("/categories", controllers.CreateCategory)
		}
		
		api.POST("/register", controllers.Register)
        api.POST("/login", controllers.Login)
		api.GET("/articles", controllers.GetAllArticles)
	}
	return router
}