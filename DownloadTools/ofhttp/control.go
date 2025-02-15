package ofhttp

import (
	"encoding/json"
	"io"
	"net/http"
)

func ConvertRespToJson(resp *http.Response) (map[string]interface{}, error) {
	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Close the response body to prevent resource leaks
	defer resp.Body.Close()

	// Unmarshal the JSON into a map
	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
