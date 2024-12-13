package mapping

import (
	"encoding/xml"
	"os"
)

// This structure parses the Overpass Data json response

// OSM stands for OpenStreet Maps- This is geolocation maps
// that am using get the coordinates of a subregion
type OSM struct {
	Version   string     `xml:"version,attr"`
	Generator string     `xml:"generator,attr"`
	Nodes     []Node     `xml:"node"`
	Ways      []Way      `xml:"way"`
	Relations []Relation `xml:"relation"`
}

// Node represents an OSM node with latitude and longitude.
type Node struct {
	ID  int64   `xml:"id,attr"`
	Lat float64 `xml:"lat,attr"`
	Lon float64 `xml:"lon,attr"`
}

// Way represents an OSM way with references to nodes.
type Way struct {
	ID    int64     `xml:"id,attr"`
	Nodes []WayNode `xml:"nd"`
	Tags  []Tag     `xml:"tag"`
}

// WayNode represents a node reference in a way.
type WayNode struct {
	Ref int64 `xml:"ref,attr"`
}

// Relation represents an OSM relation.
type Relation struct {
	ID      int64    `xml:"id,attr"`
	Members []Member `xml:"member"`
	Tags    []Tag    `xml:"tag"`
}

// Member represents a member of a relation.
type Member struct {
	Type string `xml:"type,attr"`
	Ref  int64  `xml:"ref,attr"`
	Role string `xml:"role,attr"`
}

// Tag represents a key-value pair in OSM data.
type Tag struct {
	Key   string `xml:"k,attr"`
	Value string `xml:"v,attr"`
}

// This functions returns the raw data from the xml file which
// has data of a particular region.
func ParseOSMFile(filename string) (*OSM, error) {
	// opening the osm file
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var osmData OSM
	decoder := xml.NewDecoder(file)

	err = decoder.Decode(&osmData)
	if err != nil {
		return nil, err
	}
	return &osmData, nil
}
