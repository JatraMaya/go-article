package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jatraMaya/go-library/config"
	"github.com/jatraMaya/go-library/models"
	"github.com/jatraMaya/go-library/routes"
)

func main() {
	config.LoadConfig("config/config.yaml")

	models.InitDB()

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Message": "Server is up and running",
		})
	})

	routes.SetupArticleRoutes(r)
	routes.SetupUserRoutes(r)

	r.Run(config.AppConfig.Server.Port)

}
