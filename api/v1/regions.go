package v1

import (
	"fmt"
	"encoding/json"
	"log"
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
	return fmt.Sprintf("%s (id: %d)", r.Name, r.ID)
}

func GetRegions() {
	query := fmt.Sprintf("%s?client_id=%s&api_key=%s", RegionsEndpoint, config.Conf.ClientID, config.Conf.ApiKey)
	body, err := sendQuery(query)
	if err != nil {
		log.Fatal(err)
	}

	var resp RegionsResp
	err = json.Unmarshal(body, &resp)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Regions:\n")
	for _, r := range resp.Regions {
		fmt.Printf("%v\n", r)
	}
	fmt.Printf("\n")
}
