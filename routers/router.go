package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/husni-robani/article-api-services/controllers"
	"github.com/husni-robani/article-api-services/middleware"
	"github.com/husni-robani/article-api-services/response"
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

		// endpoint that gives documentation URL in response
		api.GET("/docs", func (c *gin.Context){
			response.Success(c, http.StatusOK, gin.H{"documentation_url": "https://documenter.getpostman.com/view/26114103/2sAY4yeLhF"}, "Success")
		})
	}
	return router
}