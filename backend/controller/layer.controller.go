package controller

import (
	"encoding/json"
	"net/http"

	"github.com/Nina-99/TripSpotter/backend/config"
	"github.com/Nina-99/TripSpotter/backend/models"
	"github.com/gin-gonic/gin"
)

func GetAllLayers(c *gin.Context) {
	var layers []models.Shapefile
	if err := config.DB.Find(&layers).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot fetch layers"})
		return
	}

	//TODO: Create array of layer for response
	//TODO: Creamos un array de capas para devolver
	var result []map[string]interface{}

	for _, layer := range layers {
		var geojson map[string]interface{}
		if err := json.Unmarshal([]byte(layer.GeoJSON), &geojson); err != nil {
			//TODO: If the layer cannot be parsed, this layer is ignored
			//TODO: Si no se puede parsear, se ignora esta capa
			continue
		}

		layerItem := map[string]interface{}{
			"id":      layer.Id,
			"name":    layer.Name,
			"geojson": geojson,
		}

		result = append(result, layerItem)
	}

	//TODO: We send a list (array JSON)
	//TODO: Enviamos una lista (array JSON)
	c.JSON(http.StatusOK, result)
}
