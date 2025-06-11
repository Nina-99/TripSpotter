package controller

import (
	"net/http"

	"github.com/Nina-99/TripSpotter/backend/config"
	"github.com/Nina-99/TripSpotter/backend/models"
	"github.com/gin-gonic/gin"
)

func SubmitReview(c *gin.Context) {
	var review models.Review

	if err := c.ShouldBindJSON(&review); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Create(&review).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al guardar la reseña"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Reseña enviada correctamente"})
}
