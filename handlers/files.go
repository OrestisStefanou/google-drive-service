package handlers

import (
	"net/http"
	"fmt"

	"google-drive-service/goDrive"
	"github.com/gin-gonic/gin"
)


//Function that returns the metadata of the user's files
func GetFilesMetadata(c *gin.Context){
	userEmail := c.Param("userEmail")
	tokenPath := fmt.Sprintf("goDrive/UserTokens/%s",userEmail)
	fmt.Println("Token Path is:",tokenPath)
	files,err := goDrive.GetFileList(tokenPath)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "A token for this email doesn't exist or has expired.Please create a new one"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Files": files})
}