package algorithm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"oftools/oflog"
	"time"

	"gopkg.in/twindagger/httpsig.v1"
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

	// Marshal the data into JSON format
	jsonData, err := json.Marshal(data)
	if err != nil {
		oflog.Print.Errorf("Convert to json Data failed.")
		return err
	}

	// Create a new POST request
	req, err := http.NewRequest("POST", "https://blj.ofilm.com/api/v1/authentication/auth/", bytes.NewBuffer(jsonData))
	if err != nil {
		oflog.Print.Errorf("New Request failed.")
		return err
	}

	// Set the request header
	req.Header.Set("Content-Type", "application/json")

	// Create HTTP client and send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		oflog.Print.Errorf("Get resp failed.")
		return err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response: %v", err)
	}

	// Print the response content
	if resp.StatusCode == http.StatusOK {
		fmt.Println("Response:", string(body))
	} else {
		fmt.Printf("Request failed %s\n", resp.Body)
		fmt.Printf("Request failed with status %s\n", resp.Status)
	}
	return nil
}

func (auth *SigAuth) Sign(r *http.Request) error {
	headers := []string{"(request-target)", "date"}
	signer, err := httpsig.NewRequestSigner(auth.KeyID, auth.SecretID, "hmac-sha256")
	if err != nil {
		return err
	}
	return signer.SignRequest(r, headers, nil)
}

func GetUserInfo(jmsurl string, auth *SigAuth) {
	url := jmsurl + "/api/v1/users/users/"
	gmtFmt := "Mon, 02 Jan 2006 15:04:05 GMT"
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Date", time.Now().Format(gmtFmt))
	req.Header.Add("Accept", "application/json")
	req.Header.Add("X-JMS-ORG", "00000000-0000-0000-0000-000000000002")
	if err != nil {
		log.Fatal(err)
	}
	if err := auth.Sign(req); err != nil {
		log.Fatal(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	json.MarshalIndent(body, "", "    ")
	fmt.Println(string(body))
}

func JumpGetInfo(value string) error {
	auth := SigAuth{
		KeyID:    AccessKeyID,
		SecretID: AccessKeySecret,
	}
	GetUserInfo(JmsServerURL, &auth)
	return nil
}

// func JumpSignIn(value string) error{
// 	auth := SigAuth{
// 		KeyID:    AccessKeyID,
// 		SecretID: AccessKeySecret,
// 	}
// 	_, err := auth.Sign()
// 	return nil
// }
