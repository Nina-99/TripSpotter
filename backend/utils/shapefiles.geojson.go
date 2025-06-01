package utils

import (
	"encoding/json"
	"fmt"

	"github.com/jonas-p/go-shp"
	"github.com/paulmach/go.geojson"
)

// TODO: ConvertShapefileToGeoJSON Converts the shapefile to a Feature
// TODO: ConvertShapefileToGeoJSON convierte el shapefile a un FeatureCollection GeoJSON
func ConvertShapefileToGeoJSON(shpPath string) (string, error) {
	shpReader, err := shp.Open(shpPath)
	if err != nil {
		return "", err
	}
	defer shpReader.Close()
	fields := shpReader.Fields()

	fc := geojson.NewFeatureCollection()

	for shpReader.Next() {
		_, shape := shpReader.Shape()

		switch s := shape.(type) {

		//TODO: For Point geometry
		//TODO: Para puntos en el plano
		case *shp.Point:
			f := geojson.NewPointFeature([]float64{s.X, s.Y})
			for j, field := range fields {
				attr := shpReader.Attribute(j)
				f.SetProperty(field.String(), attr)
			}
			fc.AddFeature(f)

			//TODO: For PolyLine geometry
			//TODO: Para geometría de polilineas
		case *shp.PolyLine:
			coords := [][]float64{}
			for _, pt := range s.Points {
				coords = append(coords, []float64{pt.X, pt.Y})
			}
			f := geojson.NewLineStringFeature(coords)
			for j, field := range fields {
				attr := shpReader.Attribute(j)
				f.SetProperty(field.String(), attr)
			}
			fc.AddFeature(f)

			//TODO: For Polygon geometry
			//TODO: Para una geometría de Poligono
		case *shp.Polygon:
			coords := [][]float64{}
			for _, pt := range s.Points {
				coords = append(coords, []float64{pt.X, pt.Y})
			}
			//TODO: We ensure close ring
			//TODO: Aseguramos anillo cerrado
			if len(coords) > 0 && (coords[0][0] != coords[len(coords)-1][0] || coords[0][1] != coords[len(coords)-1][1]) {
				coords = append(coords, coords[0])
			}
			f := geojson.NewPolygonFeature([][][]float64{coords})
			for j, field := range fields {
				attr := shpReader.Attribute(j)
				f.SetProperty(field.String(), attr)
			}
			fc.AddFeature(f)

			//TODO: Unsupported geometry type
			//TODO: Tipo de geometría no soportado
		default:
			fmt.Printf("Unsupported geometry type: %T\n", s)
			continue
		}
	}

	// return json.Marshal(fc)
	geojsonBytes, err := json.Marshal(fc)
	if err != nil {
		return "", err
	}

	return string(geojsonBytes), nil
}
