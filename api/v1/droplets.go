package v1

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

// Representation of a DigitalOcean Droplet.
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

// Get all droplets under the given client_id and api_key
func GetDroplets() ([]Droplet, error) {
	query := fmt.Sprintf("%s?client_id=%s&api_key=%s",
		DropletsEndpoint,
		config.Conf.ClientID,
		config.Conf.APIKey)

	body, err := sendQuery(query)
	if err != nil {
		return nil, err
	}

	resp := struct {
		Status   string    `json:"status"`
		Droplets []Droplet `json:"droplets"`
	}{}

	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}

	if resp.Status == "ERROR" {
		return nil, errors.New("Error retrieving droplets")
	}

	return resp.Droplets, nil

}

// Gets the Droplet with the given id
func GetDroplet(id int) (Droplet, error) {
	query := fmt.Sprintf("%s/%d/?client_id=%s&api_key=%s",
		DropletsEndpoint,
		id,
		config.Conf.ClientID,
		config.Conf.APIKey)

	body, err := sendQuery(query)
	if err != nil {
		return Droplet{}, err
	}

	resp := struct {
		Status  string  `json:"status"`
		Droplet Droplet `json:"droplet"`
	}{}

	err = json.Unmarshal(body, &resp)
	if err != nil {
		return Droplet{}, err
	}

	if resp.Status == "ERROR" {
		return Droplet{}, errors.New("Error retrieving droplet")
	}

	return resp.Droplet, nil
}
