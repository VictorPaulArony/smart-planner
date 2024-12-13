package mapping

import (
	"encoding/xml"
	"os"
)

// This structure parses the Overpass Data json response

// OSM stands for OpenStreet Maps- This is geolocation maps
// that am using get the coordinates of a subregion
type OSM struct {
	XMLName   xml.Name   `xml:"osm"`
	Version   string     `xml:"version,attr"`
	Generator string     `xml:"generator,attr"`
	Relations []Relation `xml:"relation"`
}

type Relation struct {
	ID        int64    `xml:"id,attr"`
	Visible   bool     `xml:"visible,attr"`
	Version   int      `xml:"version,attr"`
	Changeset int64    `xml:"changeset,attr"`
	Timestamp string   `xml:"timestamp,attr"`
	User      string   `xml:"user,attr"`
	UID       int64    `xml:"uid,attr"`
	Members   []Member `xml:"member"`
	Tags      []Tag    `xml:"tag"`
}

type Member struct {
	Type string `xml:"type,attr"`
	Ref  int64  `xml:"ref,attr"`
	Role string `xml:"role,attr"`
}

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

	var osm OSM
	decoder := xml.NewDecoder(file)
	err = decoder.Decode(&osm)
	if err != nil {
		return nil, err
	}
	return &osm, nil
}
