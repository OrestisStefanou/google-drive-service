package main

import (
	"os"
	"io"
	"log"

	"google-drive-service/handlers"
	"github.com/gin-gonic/gin"
)


func main() {
	gin.SetMode(gin.ReleaseMode)
	//Create logging file
	f,_ := os.Create("gin.log")
	defer f.Close()
	//Write the logs to file and console at the same time
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	log.SetOutput(f)
	router := gin.Default()

	//group these endpoints 
	v1 := router.Group("/v1")
	{
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Pong",
			})
		})

		v1.GET("/authenticationURL", handlers.GetAuthenticationUrl)
		v1.POST("/token", handlers.CreateUserToken)
		v1.GET("/files", handlers.GetFilesMetadata)
		v1.GET("/files/download/:fileID",handlers.DownloadBinaryFile)
		v1.GET("/files/download_exported/:fileID",handlers.DownloadExportedFile)
		v1.POST("/files/folder", handlers.CreateFolder)
		v1.POST("/files/file", handlers.UploadFile)
		v1.POST("/permissions/permission", handlers.AddPermission)
		v1.GET("/permissions/:fileID", handlers.ListFilePermissions)
	}
	router.Run() // listen and serve on 0.0.0.0:8080
}