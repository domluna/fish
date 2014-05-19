package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	// Version of the DigitalOcean API
	Version = "v1"
	// DigitalOcean API endpoint	
	APIEndpoint = "https://api.digitalocean.com"
)

// Endpoints for DigitalOcean resources.
var (
	ImagesEndpoint = createEndpoint("images")
	DropletsEndpoint = createEndpoint("droplets")
	RegionsEndpoint = createEndpoint("regions")
)



// Forms the api url to for the DigitalOcean resource.
func createEndpoint(resource string) string {
	return fmt.Sprintf("%s/%s/%s", APIEndpoint, Version, resource)
}

// Sends a GET request to the query url and returns
// the response or an error.
func sendQuery(query string) ([]byte, error) {
	resp, err := http.Get(query)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
