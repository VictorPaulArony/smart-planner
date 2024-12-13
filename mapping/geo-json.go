package mapping

type GeoJSONFeature struct {
	Type       string                 `json:"type"`
	Geometry   map[string]interface{} `json:"geometry"`
	Properties map[string]interface{} `json:"properties"`
}

type GeoJSON struct {
	Type     string           `json:"type"`
	Features []GeoJSONFeature `json:"features"`
}

// ConvertOSMToGeoJSON converts OSM data to GeoJSON format.
func ConvertOSMToGeoJSON(osmData *OSM) *GeoJSON {
	geoJSON := &GeoJSON{
		Type:     "FeatureCollection",
		Features: []GeoJSONFeature{},
	}

	// Add nodes as points
	for _, node := range osmData.Nodes {
		feature := GeoJSONFeature{
			Type: "Feature",
			Geometry: map[string]interface{}{
				"type":        "Point",
				"coordinates": []float64{node.Lon, node.Lat},
			},
			Properties: map[string]interface{}{
				"id": node.ID,
			},
		}
		geoJSON.Features = append(geoJSON.Features, feature)
	}

	// Add ways as line strings
	for _, way := range osmData.Ways {
		coordinates := []interface{}{}
		for _, nd := range way.Nodes {
			// Match node references to actual nodes
			for _, node := range osmData.Nodes {
				if node.ID == nd.Ref {
					coordinates = append(coordinates, []float64{node.Lon, node.Lat})
				}
			}
		}
		feature := GeoJSONFeature{
			Type: "Feature",
			Geometry: map[string]interface{}{
				"type":        "LineString",
				"coordinates": coordinates,
			},
			Properties: map[string]interface{}{
				"id": way.ID,
			},
		}
		geoJSON.Features = append(geoJSON.Features, feature)
	}

	// Add relations as multipolygons (if applicable)
	// Customize this part for your specific data needs.

	return geoJSON
}
