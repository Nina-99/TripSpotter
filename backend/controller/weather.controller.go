package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func GetForecast(ctx *gin.Context) {
	lat := ctx.Query("lat")
	lon := ctx.Query("lon")

	if lat == "" || lon == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "lat and lon is required"})
		return
	}

	url := fmt.Sprintf(
		"https://api.openweathermap.org/data/2.5/forecast?lat=%s&lon=%s&units=metric&lang=es&appid=%s",
		lat, lon, os.Getenv("OPENWEATHER_API_KEY"),
	)

	//TODO: Makes the weather request
	//TODO: Realiza la petición del clima
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != 200 {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Cloud not get weather"})
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var data map[string]interface{}
	json.Unmarshal(body, &data)

	ctx.JSON(http.StatusOK, data)
}
