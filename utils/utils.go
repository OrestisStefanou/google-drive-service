package utils

import (
	"fmt"
	"io"
	"io/ioutil"

)

//Function to create the token's path from an email
func GetTokenPath(userEmail string) string {
	tokenPath := fmt.Sprintf("goDrive/UserTokens/%s",userEmail)
	return tokenPath
}

func GetFileToDownloadPath(userEmail,fileID,fileExtension string) string {
	//Here check if there is a directory for this user temporary files
	fmt.Println("Should check if dir exists for user:",userEmail)
	path := fmt.Sprintf("goDrive/UserFiles/%s.%s",fileID,fileExtension)
	fmt.Println("Trying to create file:",path)
	return path
}

//Create a local file to store a file we download from google drive
func CreateLocalFile(filepath string,filedata io.Reader) error {
        buf,err := ioutil.ReadAll(filedata)
        if err != nil {
                return err 
        }
        //Create and write the content to the new file
        err = ioutil.WriteFile(filepath,buf,0644)
        if err != nil {
                return err 
        }
        return nil        
}