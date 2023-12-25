// SPDX-License-Identifier: Apache-2.0
package config

import (
	"encoding/json"
	"io"
	"os"
)

// Config represents the configuration for avail DA integration
type Config struct {
	Seed               string `json:"seed"`
	ApiURL             string `json:"api_url"`
	AppID              int    `json:"app_id"`
	DestinationDomain  int    `json:"destination_domain"`
	DestinationAddress string `json:"destination_address"`
	Timeout            int    `json:"timeout"`
}

// GetConfig reads the configuration from a file
func (c *Config) GetConfig(configFileName string) error {
	jsonFile, err := os.Open(configFileName)
	if err != nil {
		return err
	}
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return err
	}

	err = json.Unmarshal(byteValue, c)
	if err != nil {
		return err
	}

	return nil
}
