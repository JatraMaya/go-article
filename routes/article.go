package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jatraMaya/go-library/controllers"
	"github.com/jatraMaya/go-library/utils"
)

func SetupArticleRoutes(r *gin.Engine) {
	article := r.Group("/article")
	{
		article.GET("/", utils.RequiredAuth, controllers.GetArticles)
		article.GET("/:id", utils.RequiredAuth, controllers.GetArticleById)
		article.POST("/", utils.RequiredAuth, controllers.CreateArticle)
		article.PUT("/:id", utils.RequiredAuth, controllers.UpdateArticle)
		article.DELETE("/:id", utils.RequiredAuth, controllers.DeleteArticle)
	}

}
