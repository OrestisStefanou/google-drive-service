package handlers

import (
	"net/http"
	"log"

	"google-drive-service/goDrive"
	"github.com/gin-gonic/gin"
)

//Function to add permission to a file
func AddPermission(c *gin.Context) {
	type PermissionData struct {
		FileID string   `json:"file_id" binding:"required"`
		Role string     `json:"role" binding:"required"`
		Type string     `json:"type" binding:"required"`
		Emails []string `json:"emails"`
	}

	tok,err := getTokenFromHeader(c)
	if err != nil {
		return
	}

	var requestData PermissionData
	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Println("FileID:",requestData.FileID)
	log.Println("Role:",requestData.Role)
	log.Println("Type:",requestData.Type)
	log.Println("Emails:",requestData.Emails)

	err = goDrive.AddFilePermission(tok,requestData.FileID,requestData.Role,requestData.Type,requestData.Emails)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return		
	}
	c.JSON(http.StatusOK, gin.H{"message": "Permission Created"})
}


//Function to list the permissions of a file
func ListFilePermissions(c *gin.Context) {
	fileID := c.Param("fileID")
	log.Println("FileID:",fileID)

	tok,err := getTokenFromHeader(c)
	if err != nil {
		return
	}
	permissions,err := goDrive.GetFilePermissions(tok,fileID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return			
	}
	c.JSON(http.StatusOK, gin.H{"Permissions": permissions})
}