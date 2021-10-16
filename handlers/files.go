package handlers

import (
	"net/http"
	"encoding/json"
	"log"

	"google-drive-service/goDrive"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)


//Function that returns the metadata of the user's files
func GetFilesMetadata(c *gin.Context) {
	accessToken := c.GetHeader("Authorization")
	var tok *oauth2.Token
	err := json.Unmarshal([]byte(accessToken), &tok)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid access token",
		})
		return			
	}
	files,err := goDrive.GetFileList(tok)
	if err != nil {
		c.JSON(403, gin.H{"error": "Token is invalid or expired"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Files": files})
}


//Function that sends the requested file to the user(only for files with binary content)
func DownloadBinaryFile(c *gin.Context) {
	fileID := c.Param("fileID")
	log.Println("fileID:",fileID)
	accessToken := c.GetHeader("Authorization")
	var tok *oauth2.Token
	err := json.Unmarshal([]byte(accessToken), &tok)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid access token",
		})
		return			
	}
	fileData,err := goDrive.DownloadFile(tok,fileID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return			
	}
	c.Data(200,"application/binary",fileData)
}


//Function that exports the requested file to format provided 
// and then downloads it and sends it to user
func DownloadExportedFile(c *gin.Context) {
	fileID := c.Param("fileID")
	mimeType := c.Query("mimeType")
	log.Println("fileID:",fileID)
	log.Println("mimeType:",mimeType)
	accessToken := c.GetHeader("Authorization")
	var tok *oauth2.Token
	err := json.Unmarshal([]byte(accessToken), &tok)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid access token",
		})
		return			
	}
	fileData,err := goDrive.DownloadExportedFile(tok,fileID,mimeType)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return			
	}
	c.Data(200,"application/binary",fileData)
}