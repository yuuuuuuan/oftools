package ofhttp

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func ConvertRespToJson(resp *http.Response) (map[string]interface{}, error) {
	if resp == nil {
		return nil, fmt.Errorf("received nil response")
	}
	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("received nil response")
	}

	// Close the response body to prevent resource leaks
	defer resp.Body.Close()

	// Unmarshal the JSON into a map
	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, fmt.Errorf("received nil response")
	}
	log.Println("ConvertRespToJson")
	log.Println(result)
	return result, nil
}

// 函数从 map[string]interface{} 中提取 "data" 键的值，并返回 map[string]string
func ExtractDataAsStringMap(input map[string]interface{}) (map[string]string, error) {
	// 检查 map 中是否有 "data" 键
	if data, exists := input["data"]; exists {
		// 尝试将 "data" 转换为 map[string]string 类型
		if dataMap, ok := data.(map[string]string); ok {
			return dataMap, nil
		}
		return nil, fmt.Errorf("data is not of type map[string]string")
	}
	return nil, fmt.Errorf("data key does not exist in map")
}
