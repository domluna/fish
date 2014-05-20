package v1

import (
	"encoding/json"
	"fmt"
	"log"
)

type Sizes struct {
	Status string  `json:"status"`
	Sizes []Size `json:"sizes"`
}

type Size struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Slug         string `json:"slug"`
}

func (s Size) String() string {
	return fmt.Sprintf("%s (id: %d)", s.Name, s.ID)
}

func GetSizes() {
	query := fmt.Sprintf("%s?client_id=%s&api_key=%s", SizesEndpoint, config.Conf.ClientID, config.Conf.ApiKey)
	body, err := sendQuery(query)
	if err != nil {
		log.Fatal(err)
	}

	var resp Sizes
	err = json.Unmarshal(body, &resp)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Sizes:\n")
	for _, s := range resp.Sizes {
		fmt.Printf("%v\n", s)
	}
	fmt.Printf("\n")
}
