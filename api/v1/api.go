package v1

import (
	"io/ioutil"
	"net/http"
)

const (
	// Version of the DigitalOcean API
	Version = "v1"

	// Image endpoint
	ImagesEndpoint = "https://api.digitalocean.com/v1/images"

	// Droplet endpoint
	DropletsEndpoint = "https://api.digitalocean.com/v1/droplets"

	// Region endpoint
	RegionsEndpoint = "https://api.digitalocean.com/v1/regions"

	// Sizes endpoint
	SizesEndpoint = "https://api.digitalocean.com/v1/sizes"

	// SSH Key endpoint
	KeysEndpoint = "https://api.digitalocean.com/v1/ssh_keys"
)

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
