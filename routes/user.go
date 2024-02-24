package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jatraMaya/go-library/controllers"
)

func SetupUserRoutes(r *gin.Engine) {
	user := r.Group("")
	{
		user.POST("/signup", controllers.SignUp)
		user.POST("/login", controllers.Login)
		user.GET("/login", controllers.Logout)
	}
}
