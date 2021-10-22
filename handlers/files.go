package handlers

import (
	"net/http"
	"log"

	"google-drive-service/goDrive"
	"github.com/gin-gonic/gin"
)


//Function that returns the metadata of the user's files
func GetFilesMetadata(c *gin.Context) {
	tok,err := getTokenFromHeader(c)
	if err != nil {
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
	tok,err := getTokenFromHeader(c)
	if err != nil {
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
	tok,err := getTokenFromHeader(c)
	if err != nil {
		return
	}
	fileData,err := goDrive.DownloadExportedFile(tok,fileID,mimeType)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return			
	}
	c.Data(200,"application/binary",fileData)
}


//Function that creates a new folder in user's drive
func CreateFolder(c *gin.Context) {
	type CreateFolderRequest struct {
		FolderName string `json:"folder_name" binding:"required"`
		ParentId string `json:"parent_id"`
	}
	tok,err := getTokenFromHeader(c)
	if err != nil {
		return
	}
	var jsonCreateFolderReq CreateFolderRequest
	if err := c.ShouldBindJSON(&jsonCreateFolderReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Println("Folder name:",jsonCreateFolderReq.FolderName)
	log.Println("Parent id:",jsonCreateFolderReq.ParentId)

	err = goDrive.CreateFolder(tok,jsonCreateFolderReq.FolderName,jsonCreateFolderReq.ParentId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return		
	}
	c.JSON(http.StatusOK, gin.H{"message": "Folder Created"})
}


//Function that uploads a file in the user's drive
func UploadFile(c *gin.Context) {
	tok,err := getTokenFromHeader(c)
	if err != nil {
		return
	}
	multipartForm,err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return				
	}
	var parentId string
	parent_value,exists := multipartForm.Value["parent_id"]
	if exists {
		parentId = parent_value[0]
	}
	log.Println("ParentId:",parentId)
	files,exists := multipartForm.File["file"]
	if exists == false {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File to upload not found"})
		return		
	}
	file := files[0]
	fileName := file.Filename
	log.Println("Filename:",fileName)
	fileObj,err := file.Open()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return		
	}
	err = goDrive.UploadFile(tok,fileObj,parentId,fileName)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return		
	}
	c.JSON(http.StatusOK, gin.H{"message": "File Uploaded"})
}