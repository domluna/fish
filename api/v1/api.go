package v1

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

var (
	// Image endpoint
	ImagesEndpoint = createEndpoint("images")

	// Droplet endpoint
	DropletsEndpoint = createEndpoint("droplets")

	// Region endpoint
	RegionsEndpoint = createEndpoint("regions")

	// Sizes endpoint
	SizesEndpoint = createEndpoint("sizes")

	// SSH Key endpoint
	KeysEndpoint = createEndpoint("ssh_keys")
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
