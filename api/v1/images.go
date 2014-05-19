package v1

import (
	"encoding/json"
	"fmt"
	"log"
)

type ImagesResp struct {
	Status string  `json:"status"`
	Images []Image `json:"images"`
}

type Image struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Distribution string `json:"distribution"`
	Slug         string `json:"slug"`
	Public       bool   `json:"public"`
}

func (i Image) String() string {
	return fmt.Sprintf("%s (distribution: %s, id: %d)", i.Name, i.Distribution, i.ID)
}

func GetImages() {
	query := fmt.Sprintf("%s?client_id=%s&api_key=%s&filter=global", ImagesEndpoint, config.Conf.ClientID, config.Conf.ApiKey)
	body, err := sendQuery(query)
	if err != nil {
		log.Fatal(err)
	}

	var resp ImagesResp
	err = json.Unmarshal(body, &resp)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Images:\n")
	for _, i := range resp.Images {
		fmt.Printf("%v\n", i)
	}
	fmt.Printf("\n")
}
