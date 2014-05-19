package v1

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
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
	return fmt.Sprintf("%s (Region: %d, image_id: %d, ip: \"%s\", status: %s)", d.Name, d.RegionID, d.ImageID, d.IPAddress, d.Status)
}

// Get all droplets under the given client_id and api_key
func GetDroplets() {
	query := fmt.Sprintf("%s?client_id=%s&api_key=%s", DropletsEndpoint, config.Conf.ClientID, config.Conf.ApiKey)
	body, err := sendQuery(query)
	if err != nil {
		log.Fatal(err)
	}

	var resp DropletsResp
	err = json.Unmarshal(body, &resp)
	if err != nil {
		log.Fatal(err)
	}

	// Output all droplets
	fmt.Printf("Droplets:\n")
	for _, d := range resp.Droplets {
	    fmt.Printf("%v", d)
	}
	fmt.Printf("\n")

}
