package algorithm

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
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
