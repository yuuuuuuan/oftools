package algorithm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"oftools/oflog"
	"github.com/google/uuid"
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

// Define the structure to match your JSON data
type ConnectionData struct {
	ID             string `json:"id"`
	Asset          string `json:"asset"`
	Account        string `json:"account"`
	InputUsername  string `json:"input_username"`
	InputSecret    string `json:"input_secret"`
	ConnectMethod  string `json:"connect_method"`
	ConnectOptions struct {
		Charset                    string `json:"charset"`
		Reusable                   bool   `json:"reusable"`
		Resolution                 string `json:"resolution"`
		DisableAutoHash            bool   `json:"disableautohash"`
		BackspaceAsCtrlH           bool   `json:"backspaceAsCtrlH"`
		AppletConnectMethod        string `json:"appletConnectMethod"`
		TerminalThemeName          string `json:"terminal_theme_name"`
		FileNameConflictResolution string `json:"file_name_conflict_resolution"`
	} `json:"connect_options"`
	Protocol   string `json:"protocol"`
	CreatedBy  string `json:"created_by"`
	UpdatedBy  string `json:"updated_by"`
	OrgID      string `json:"org_id"`
	IsActive   bool   `json:"is_active"`
	IsReusable bool   `json:"is_reusable"`
}

func JumpServer() error {
	newuuid := uuid.New().String()
	// Create an instance of the struct and populate it with values
	data := ConnectionData{
		ID:            newuuid,
		Asset:         "f071d78c-df9a-4b8e-aaac-13fcbfa1426e",
		Account:       "@INPUT",
		InputUsername: "NF3266",
		InputSecret:   "",
		ConnectMethod: "mstsc",
		ConnectOptions: struct {
			Charset                    string `json:"charset"`
			Reusable                   bool   `json:"reusable"`
			Resolution                 string `json:"resolution"`
			DisableAutoHash            bool   `json:"disableautohash"`
			BackspaceAsCtrlH           bool   `json:"backspaceAsCtrlH"`
			AppletConnectMethod        string `json:"appletConnectMethod"`
			TerminalThemeName          string `json:"terminal_theme_name"`
			FileNameConflictResolution string `json:"file_name_conflict_resolution"`
		}{
			Charset:                    "default",
			Reusable:                   false,
			Resolution:                 "auto",
			DisableAutoHash:            false,
			BackspaceAsCtrlH:           false,
			AppletConnectMethod:        "client",
			TerminalThemeName:          "Default",
			FileNameConflictResolution: "replace",
		},
		Protocol:   "rdp",
		CreatedBy:  "邹航远",
		UpdatedBy:  "邹航远",
		OrgID:      "7d81c9a7-90a0-4ee0-9d60-ded84e971314",
		IsActive:   true,
		IsReusable: false,
	}

	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
		return err
	}
	fmt.Println(string(jsonData))
	return nil
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
	fmt.Printf("%+v\n", resp)
	fmt.Printf("%#v\n", resp)

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
