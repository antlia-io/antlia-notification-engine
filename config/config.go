package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Config struct {
	NetworkWebSocket string `json:"networkWebSocket"`
	ContractAddress  string `json:"contractAddress"`
}

func LoadConfiguration(file string) Config {
	var config Config
	bytes, err := ioutil.ReadFile(file)
	log.Println(file)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(bytes, &config)
	if err != nil {
		panic(err)
	}

	return config
}
