// filehandler/filehandler.go
package filehandler

import (
	"fmt"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

// HandleUpload handles file upload
func HandleUpload(c *gin.Context) (string, error) {
	file, err := c.FormFile("file")
	if err != nil {
		return "", fmt.Errorf("get form err: %s", err)
	}

	// Save the uploaded file
	dst := filepath.Base(file.Filename)
	if err := c.SaveUploadedFile(file, "uploads/"+dst); err != nil {
		return "", fmt.Errorf("upload file err: %s", err)
	}

	return dst, nil
}
