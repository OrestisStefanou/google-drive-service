package handlers

import (
	"net/http"
	"fmt"

	"google-drive-service/goDrive"
	"github.com/gin-gonic/gin"
)

//Function to generate a url that the user can use to get an authentication code
func GetAuthenticationUrl(c *gin.Context) {
	authURL := goDrive.Get_user_auth_url()
	c.JSON(200, gin.H{
		"message": "Please go to this link to get an authentication code",
		"authURL": authURL,
	})	
}


//A function to create an authentication token using the authentication code that the user sends
func CreateUserToken(c *gin.Context) {
	type AuthCodeRequest struct {
		Email string `form:"email" binding:"required"`
		Code string `form:"code" binding:"required"`
	}
	var jsonAuthCodeReq AuthCodeRequest
	if err := c.ShouldBind(&jsonAuthCodeReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//Do basic logging here
	fmt.Println("Email:",jsonAuthCodeReq.Email)
	fmt.Println("Authentication code:",jsonAuthCodeReq.Code)
	tokenPath := fmt.Sprintf("goDrive/UserTokens/%s",jsonAuthCodeReq.Email)
	fmt.Println("Token Path is:",tokenPath)
	//First check if token already exists and if it exists check if is still valid
	exists,err := goDrive.TokenExists(tokenPath)
	if exists && err == nil {
		fmt.Println("Token exists")
		tokenIsValid,err := goDrive.TokenIsValid(tokenPath)
		if tokenIsValid && err == nil {
			c.JSON(http.StatusOK, gin.H{"message": "Token already exists"})
			return
		} 
	}
	//Create the user authentication token
	err = goDrive.CreateUserToken(jsonAuthCodeReq.Code,tokenPath)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Token created!"})	
}