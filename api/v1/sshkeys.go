package v1

import (
	"encoding/json"
	"fmt"
)

// DigitalOcean ssh key representation.
type SSHKey struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	SSHPublicKey string `json:"ssh_pub_key"`
}

// Default string output representation of ssh key.
func (k SSHKey) String() string {
	return fmt.Sprintf("%s (id: %d)", k.Name, k.ID)
}

// Get all users ssh keys. The name and id of each
// key will be returned.
func AllSSHKeys() ([]SSHKey, error) {
	query := fmt.Sprintf("%s?client_id=%s&api_key=%s",
		KeysEndpoint,
		config.Conf.ClientID,
		config.Conf.ApiKey)

	body, err := sendQuery(query)
	if err != nil {
		return nil, err
	}

	resp := struct {
		Status  string   `json:"status"`
		SSHKeys []SSHKey `json:"ssh_keys"`
	}{}

	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}

	return resp.SSHKeys, nil
}

// Adds an ssh key to the user account
func AddSSHKey(name string, publicKey string) (SSHKey, error) {
	query := fmt.Sprintf("%s/new/?name=%s&ssh_pub_key=%s&client_id=%s&api_key=%s",
		KeysEndpoint,
		name,
		publicKey,
		config.Conf.ClientID,
		config.Conf.ApiKey)

	body, err := sendQuery(query)
	if err != nil {
		return SSHKey{}, err
	}

	resp := struct {
		Status string `json:"status"`
		SSHKey SSHKey `json:"ssh_key"`
	}{}

	err = json.Unmarshal(body, &resp)
	if err != nil {
		return SSHKey{}, err
	}

	return resp.SSHKey, nil
}

// Show the full public ssh key of the passed id.
func ShowSSHKey(id int) (SSHKey, error) {
	query := fmt.Sprintf("/%d/%s?client_id=%s&api_key=%s",
		id,
		KeysEndpoint,
		config.Conf.ClientID,
		config.Conf.ApiKey)

	body, err := sendQuery(query)
	if err != nil {
		return SSHKey{}, err
	}

	resp := struct {
		Status string `json:"status"`
		SSHKey SSHKey `json:"ssh_key"`
	}{}

	err = json.Unmarshal(body, &resp)
	if err != nil {
		return SSHKey{}, err
	}

	return resp.SSHKey, nil
}

// Destroys the ssh key with passed id from
// user account.
func DestroySSHKey(id int) error {
	query := fmt.Sprintf("/%d/destroy/%s?client_id=%s&api_key=%s",
		id,
		KeysEndpoint,
		config.Conf.ClientID,
		config.Conf.ApiKey)

	body, err := sendQuery(query)
	if err != nil {
		return err
	}

	resp := struct {
		Status string `json:"status"`
	}{}

	err = json.Unmarshal(body, &resp)
	if err != nil {
		return err
	}

	return nil
}
