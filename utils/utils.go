package utils

import (
	"fmt"

)

//Function to create the token's path from an email
func GetTokenPath(userEmail string) string {
	tokenPath := fmt.Sprintf("goDrive/UserTokens/%s",userEmail)
	return tokenPath
}