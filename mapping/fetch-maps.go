package mapping

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func KisumuMap(url, query string) (interface{}, error) {
	// Send request
	resp, err := http.Post(url, "text/plain", bytes.NewBufferString(query))
	if err != nil {
		return nil, errors.New("Error making request to Overpass API")
	}
	defer resp.Body.Close()

	// Read response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("Error reading response")
	}

	// Parse JSON response
	var data interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, errors.New("Error parsing JSON response")
	}
	return data, nil
}
