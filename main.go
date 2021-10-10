package main

import (
	"os"

	"google-drive-service/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	//Create logging file
	f,_ := os.Create("gin.log")
	//Write the logs to file and console at the same time
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	router := gin.Default()

	//group these endpoints 
	v1 := router.Group("/v1")
	{
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})

		v1.GET("/authenticationURL", handlers.GetAuthenticationUrl)
		v1.POST("/token", handlers.CreateUserToken)
		v1.GET("/files/:userEmail", handlers.GetFilesMetadata)
	}

	router.Run() // listen and serve on 0.0.0.0:8080
}