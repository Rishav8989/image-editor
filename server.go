// main.go
package main

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"

	"image-editor/packages/filehandler"
	"image-editor/packages/imageprocessor"
)

func main() {
	router := gin.Default()

	// Load HTML templates
	router.LoadHTMLGlob("templates/*")

	// Serve the upload form
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "upload.html", nil)
	})

	// Handle the file upload
	router.POST("/upload", func(c *gin.Context) {
		// Handle file upload
		uploadedFileName, err := filehandler.HandleUpload(c)
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("File upload error: %s", err))
			return
		}

		// Get form values
		widthStr := c.PostForm("width")
		heightStr := c.PostForm("height")
		blurLevelStr := c.PostForm("blur_level")

		width, err := strconv.Atoi(widthStr)
		if err != nil {
			c.String(http.StatusBadRequest, "Invalid width")
			return
		}

		height, err := strconv.Atoi(heightStr)
		if err != nil {
			c.String(http.StatusBadRequest, "Invalid height")
			return
		}

		blurLevel, err := strconv.ParseFloat(blurLevelStr, 64)
		if err != nil {
			c.String(http.StatusBadRequest, "Invalid blur level")
			return
		}

		// Process the uploaded image with the specified parameters
		inputImagePath := filepath.Join("uploads", uploadedFileName)
		outputImagePath := filepath.Join("downloads", uploadedFileName)

		if err := imageprocessor.ApplyBlur(inputImagePath, outputImagePath, blurLevel); err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("Image processing error: %s", err))
			return
		}

		// Resize the processed image
		resizedImagePath := outputImagePath + ".resized.jpg"
		if err := imageprocessor.ResizeTo(inputImagePath, resizedImagePath, width, height); err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("Image resizing error: %s", err))
			return
		}

		// Redirect to the output folder
		c.Redirect(http.StatusSeeOther, "/downloads")
	})

	// Serve the processed image for download
	router.Static("/downloads", "./downloads")

	// Start the server
	if err := router.Run(":8080"); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
