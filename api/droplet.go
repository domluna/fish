package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// URL for querying droplets API.
var (
	dropletURL = formatURL("droplets")
	dresp      DropletsResp
)

// JSON response of Droplets.
type DropletsResp struct {
	Status   string    `json:"status"`
	Droplets []Droplet `json:"droplets"`
}

// Represents a DigitalOcean Droplet.
type Droplet struct {
	ID               int       `json:"id"`
	Name             string    `json:"name"`
	ImageID          int       `json:"image_id"`
	SizeID           int       `json:"size_id"`
	RegionID         int       `json:"region_id"`
	BackupsActive    bool      `json:"backups_active"`
	IPAddress        string    `json:"ip_address"`
	PrivateIPAddress bool      `json:"private_ip_address"`
	Locked           bool      `json:"locked"`
	Status           string    `json:"status"`
	CreatedAt        time.Time `json:"created_at"`
}

// String representation of a droplet.
func (d Droplet) String() string {
    return fmt.Sprintf("%s (Region: %d, ip: \"%s\", status: %s)\n", d.Name, d.RegionID, d.IPAddress, d.Status)
}

// Get all droplets under the given client_id and api_key
func GetDroplets() {
	query := fmt.Sprintf("%s?client_id=%s&api_key=%s", dropletURL, config.Conf.ClientID, config.Conf.ApiKey)
	resp, err := http.Get(query)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(body, &dresp)
	if err != nil {
		log.Fatal(err)
	}

	// Output all droplets
	fmt.Printf("Droplets:\n")
	for _, d := range dresp.Droplets {
	    fmt.Printf("%v", d)
	}

}
