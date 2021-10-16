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
		AuthCode string `form:"code" binding:"required"`
	}
	var jsonAuthCodeReq AuthCodeRequest
	if err := c.ShouldBind(&jsonAuthCodeReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//Do basic logging here
	fmt.Println("Authentication code:",jsonAuthCodeReq.AuthCode)

	//Get the user authentication token
	accessToken,err := goDrive.GetUserToken(jsonAuthCodeReq.AuthCode)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"error":"Invalid auth code"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Token created!", 
		"AccessToken":accessToken,
	})	
}