package main

import (
	"os"
	"io"
	"encoding/json"
	"fmt"

	"google-drive-service/handlers"
	"github.com/gin-gonic/gin"
    "golang.org/x/oauth2"
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
			headers := c.Request.Header 
			fmt.Println(headers)
			token := c.GetHeader("Authorization")
			var tok *oauth2.Token
			err := json.Unmarshal([]byte(token), &tok)
			if err != nil {
				c.JSON(200, gin.H{
					"message": "Something went wrong",
				})
				return			
			}
			c.JSON(200, gin.H{
				"message": "Token",
				"Token" : tok,
			})
		})

		v1.GET("/authenticationURL", handlers.GetAuthenticationUrl)
		v1.POST("/token", handlers.CreateUserToken)
		v1.GET("/files", handlers.GetFilesMetadata)
		v1.GET("/files/download/:fileID",handlers.DownloadBinaryFile)
	}

	router.Run() // listen and serve on 0.0.0.0:8080
}