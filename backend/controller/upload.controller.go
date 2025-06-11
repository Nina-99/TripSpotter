package controller

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/Nina-99/TripSpotter/backend/config"
	"github.com/Nina-99/TripSpotter/backend/models"
	"github.com/gin-gonic/gin"
)

func UploadImage(c *gin.Context) {
	markerIDStr := c.PostForm("marker_id")
	markerID, err := strconv.ParseUint(markerIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid marker_id"})
		return
	}

	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file is received"})
		return
	}

	//TODO: Path on Saved image
	//TODO: Ruta donde se guardará la imagen
	filename := filepath.Base(file.Filename)
	path := fmt.Sprintf("uploads/images/%d_%s", markerID, filename)

	if err := c.SaveUploadedFile(file, path); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot save file"})
		return
	}

	image := models.Image{
		Filename: filename,
		Path:     path,
		MarkerId: uint(markerID),
	}

	if err := config.DB.Create(&image).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot save image info"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Image uploaded successfully", "path": path})
}

