package handlers

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

func getTokenFromHeader(c *gin.Context) (*oauth2.Token,error) {
	accessToken := c.GetHeader("Authorization")
	var tok *oauth2.Token
	err := json.Unmarshal([]byte(accessToken), &tok)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid access token",
		})
		return nil,err	
	}
	return tok,nil		
}