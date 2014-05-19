package api

import (
	"fmt"
)

var (
	regionURL = formatURL("regions")
)

type RegionsResp struct {
	Status  string   `json:"status"`
	Regions []Region `json:"regions"`
}

type Region struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

func (r Region) String() string {
	return fmt.Sprintf("%s (id: %s)\n", r.Name, r.ID)
}

func GetRegions() {
	query := fmt.Sprintf("%s?client_id=%s&api_key=%s", regionURL, config.Conf.ClientID, config.Conf.ApiKey)
}
