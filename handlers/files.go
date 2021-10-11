package handlers

import (
	"net/http"
	"fmt"

	"google-drive-service/goDrive"
	"google-drive-service/utils"
	"github.com/gin-gonic/gin"
)


//Function that returns the metadata of the user's files
func GetFilesMetadata(c *gin.Context){
	userEmail := c.Param("userEmail")
	tokenPath := utils.GetTokenPath(userEmail)
	fmt.Println("Token Path is:",tokenPath)
	//Get the token from file
	tok,err := goDrive.GetTokenFromFile(tokenPath)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "A token for this email doesn't exist.Please create a new one"})
		return	
	}
	accessToken := c.GetHeader("Authorization")
	if accessToken != tok.AccessToken {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Access token given not valid"})
		return			
	}
	files,err := goDrive.GetFileList(tok)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "A token for this email doesn't exist or has expired.Please create a new one"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Files": files})
}