package v1

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
)

// DigitalOcean ssh key representation.
type SSHKey struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	SSHPublicKey string `json:"ssh_pub_key"`
}

// Get all the users current ssh keys.
func AllSSHKeys() ([]SSHKey, error) {
	query := fmt.Sprintf("%s?client_id=%s&api_key=%s",
		KeysEndpoint,
		config.Conf.ClientID,
		config.Conf.APIKey)

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

	if resp.Status == "ERROR" {
		return nil, errors.New("Error retrieving ssh keys")
	}

	return resp.SSHKeys, nil
}

// Adds an ssh key to the user account
func AddSSHKey(name, publicKey string) (SSHKey, error) {
	query := fmt.Sprintf("%s/new/?name=%s&ssh_pub_key=%s&client_id=%s&api_key=%s",
		KeysEndpoint,
		name,
		url.QueryEscape(publicKey),
		config.Conf.ClientID,
		config.Conf.APIKey)

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

	if resp.Status == "ERROR" {
		return SSHKey{}, errors.New("Error adding key")
	}

	return resp.SSHKey, nil
}

// Show the full public ssh key of the passed id.
func ShowSSHKey(id int) (SSHKey, error) {
	query := fmt.Sprintf("%s/%d/?client_id=%s&api_key=%s",
		KeysEndpoint,
		id,
		config.Conf.ClientID,
		config.Conf.APIKey)

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

	if resp.Status == "ERROR" {
		return SSHKey{}, errors.New("Invalid ssh key id")
	}

	return resp.SSHKey, nil
}

// Destroys the ssh key with passed id from
// user account.
func DestroySSHKey(id int) error {
	query := fmt.Sprintf("%s/%d/destroy/?client_id=%s&api_key=%s",
		KeysEndpoint,
		id,
		config.Conf.ClientID,
		config.Conf.APIKey)

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

	if resp.Status == "ERROR" {
		errors.New("Invalid ssh key id")
	}

	return nil
}
