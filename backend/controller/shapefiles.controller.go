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

	zipFile := files[0]
	zipPath := filepath.Join(tempDir, filepath.Base(zipFile.Filename))
	if err := ctx.SaveUploadedFile(zipFile, zipPath); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot save zip"})
		return
	}

	//TODO: Unzip the .zip
	//TODO: Extraer el .zip
	shpExtractedDir := filepath.Join(tempDir, "extracted")
	os.MkdirAll(shpExtractedDir, os.ModePerm)
	err := utils.Unzip(zipPath, shpExtractedDir)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot unzip file"})
		return
	}

	//TODO: Find the .shp file
	//TODO: Encontrar el archivo .shp
	var shpPath string
	shpExtractedDir = shpExtractedDir + "/" + strings.SplitN(zipFile.Filename, ".", 2)[0]
	err = filepath.Walk(shpExtractedDir, func(path string, info os.FileInfo, err error) error {
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
	geojsonStr, err := utils.ConvertShapefileToGeoJSON(shpPath)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not convert to GeoJSON"})
		return
	}

	//TODO: Save in DataBase
	//TODO: Guardar en la BD
	shapefile := models.Shapefile{Name: zipFile.Filename, GeoJSON: geojsonStr}
	if err := config.DB.Create(&shapefile).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not insert into database"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Shapefile uploaded and stored successfully"})
}
