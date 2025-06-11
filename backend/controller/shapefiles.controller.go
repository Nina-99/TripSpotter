package controller

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/Nina-99/TripSpotter/backend/config"
	"github.com/Nina-99/TripSpotter/backend/models"
	"github.com/Nina-99/TripSpotter/backend/utils"
	"github.com/gin-gonic/gin"
)

// TODO: Function for Upload shapefile to Postgis
// TODO: Función para cargar un shapefile a Postgis
func UploadShapefile(ctx *gin.Context) {
	form, _ := ctx.MultipartForm()
	files := form.File["files"]
	if len(files) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "No file uploaded"})
		return
	}

	tempDir := "./uploads"
	os.MkdirAll(tempDir, os.ModePerm)

	file := files[0]
	fileExt := strings.ToLower(filepath.Ext(file.Filename))
	savedFilePath := filepath.Join(tempDir, filepath.Base(file.Filename))

	if err := ctx.SaveUploadedFile(file, savedFilePath); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot save file"})
		return
	}

	var geojsonStr string
	switch fileExt {
	case ".zip":

		//TODO: Unzip the .zip
		//TODO: Extraer el .zip
		shpExtractedDir := filepath.Join(tempDir, "extracted")
		os.MkdirAll(shpExtractedDir, os.ModePerm)
		err := utils.Unzip(savedFilePath, shpExtractedDir)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot unzip file"})
			return
		}

		//TODO: Find the .shp file
		//TODO: Encontrar el archivo .shp
		subDir := filepath.Join(shpExtractedDir, strings.TrimSuffix(file.Filename, ".zip"))
		var shpPath string
		err = filepath.Walk(subDir, func(path string, info os.FileInfo, err error) error {
			if filepath.Ext(path) == ".shp" {
				shpPath = path
			}
			return nil
		})
		if err != nil || shpPath == "" {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Shapefile (.shp) not found in archive"})
			return
		}

		//TODO: Convert to GeoJSON
		//TODO: Convertir a GeoJSON
		jsonStr, err := utils.ConvertShapefileToGeoJSON(shpPath)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not convert to GeoJSON"})
			return
		}
		geojsonStr = jsonStr
	case ".geojson":
		// Leer el archivo .geojson
		content, err := os.ReadFile(savedFilePath)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot read GeoJSON file"})
			return
		}
		geojsonStr = string(content)
	default:
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Unsupported file format. Only .zip or .geojson allowed"})
		return
	}

	//TODO: Save in DataBase
	//TODO: Guardar en la BD
	shapefile := models.Shapefile{Name: file.Filename, GeoJSON: geojsonStr}
	if err := config.DB.Create(&shapefile).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not insert into database"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Shapefile uploaded and stored successfully"})
}
