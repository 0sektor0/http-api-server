package main

import (
	"encoding/json"
	"io/ioutil"
)

type Configs struct {
	Connector  string `json: "connector"`
	Connection string `json: "connection"`
	Port       string `json: "port"`
}

func LoadConfigs(path string) (*Configs, error) {
	config := &Configs{}

	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bytes, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
