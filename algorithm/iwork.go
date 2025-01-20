package algorithm

import (
	"fmt"
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

func IworkGet() error{
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
