package algorithm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"oftools/oflog"
)

const (
	JmsServerURL    = "http://192.168.120.11"
	AccessKeyID     = "28f18d9c-3eab-4506-a2ae-11886e9db7c9"
	AccessKeySecret = "SOKNq0vCHXl1emQVwcsOP5gQ64FOsrPHopl3"
)

type SigAuth struct {
	KeyID    string
	SecretID string
}

func JumpUpdateToken() error {
	// Define the request body data
	data := map[string]string{
		"username": "NF3266",
		"password": "Zhy1395131175",
	}

	// Define the request URL
	url := "https://blj.ofilm.com/api/v1/authentication/auth/"

	// Marshal the data into JSON format
	jsonData, err := json.Marshal(data)
	if err != nil {
		oflog.Print.Errorf("Convert to json Data failed.")
		return err
	}
	oflog.Print.Infof("%s", jsonData)
	// Create a new POST request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		oflog.Print.Errorf("New Request failed.")
		return err
	}

	// Set the request header
	req.Header.Set("Content-Type", "application/json")
	fmt.Printf("Request %s\n", req.Header)
	// Create HTTP client and send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		oflog.Print.Errorf("Get resp failed.")
		return err
	}
	defer resp.Body.Close()
	oflog.Print.Infof("%s.", resp.Body)
	// // Read the response body
	// body, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Fatalf("Error reading response: %v", err)
	// }

	// // Print the response content
	// if resp.StatusCode == http.StatusOK {
	// 	fmt.Println("Response:", string(body))
	// } else {
	// 	fmt.Printf("Request failed %s\n", resp.Body)
	// 	fmt.Printf("Request failed with status %s\n", resp.Status)
	// }
	return nil
}
