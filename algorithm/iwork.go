package algorithm

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"

	"github.com/PuerkitoBio/goquery"
)

// Fetch all /NF links from a given URL
func getNFLinks(url string) ([]string, error) {
	// Send HTTP GET request
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Check for successful status code
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("unable to access %s, status code: %d", url, resp.StatusCode)
	}

	// Parse the HTML using goquery
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	// Use a regular expression to extract all links starting with /NF
	var nfLinks []string
	re := regexp.MustCompile(`/NF\d+`)
	doc.Find("a[href]").Each(func(i int, s *goquery.Selection) {
		link, exists := s.Attr("href")
		if exists && re.MatchString(link) {
			nfLinks = append(nfLinks, link)
		}
	})

	return nfLinks, nil
}

func IworkGet() error {
	// Base URL
	baseURL := "https://it.ofilm.com/hr/hr-ks//rest/kskinsfolk/kskinsfolk/findUserNoNcHrEK/"

	// Starting URL
	startURL := baseURL + "NF3266"

	// Get all /NF links
	nfLinks, err := getNFLinks(startURL)
	if err != nil {
		log.Fatalf("Error fetching links: %v", err)
		return err
	}

	// Print all found links
	if len(nfLinks) > 0 {
		fmt.Println("Found the following /NF related links:")
		for _, link := range nfLinks {
			fmt.Println(baseURL + link)
		}
	} else {
		fmt.Println("No /NF related links found.")
	}
	return nil
}

func IworkSent(user string) error {

	// Define the request URL
	url := "https://it.ofilm.com/hr/hr-ks//rest/kskinsfolk/kskinsfolk/findUserNoNcHrEK/" + user

	// Define the request body (empty JSON data)
	jsonData := []byte(`{}`)

	// Create a new HTTP request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil
	}

	// Set headers
	req.Header.Set("Host", "it.ofilm.com")
	req.Header.Set("Content-Type", "application/json")

	// Skip HTTPS certificate verification (insecure)
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return nil
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return nil
	}

	// Print the response status and body
	fmt.Printf("Status Code: %d\n", resp.StatusCode)
	fmt.Printf("Response Body: %s\n", body)
	return nil
}
