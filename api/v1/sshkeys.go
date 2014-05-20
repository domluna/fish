package v1

import (
	"encoding/json"
	"fmt"
	"log"
)

// API response for SSHKeys
type SSHKeys struct {
	Status string  `json:"status"`
	SSHKeys []SSHKey `json:"ssh_keys"`
}

// DigitalOcean SSHKey json form
type SSHKey struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
}

func (k SSHKey) String() string {
	return fmt.Sprintf("%s (id: %d)", k.Name, k.ID)
}

func GetSSHKeys() {
	query := fmt.Sprintf("%s?client_id=%s&api_key=%s", KeysEndpoint, config.Conf.ClientID, config.Conf.ApiKey)
	body, err := sendQuery(query)
	if err != nil {
		log.Fatal(err)
	}

	var resp SSHKeys
	err = json.Unmarshal(body, &resp)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("SSHKeys:\n")
	for _, k := range resp.SSHKeys {
		fmt.Printf("%v\n", k)
	}
	fmt.Printf("\n")
}
