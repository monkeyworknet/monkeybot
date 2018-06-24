package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var (
	Token      string
	BotPrefix  string
	MCServer   string
	MCPort     string
	WeatherAPI string

	config *configStruct
)

type configStruct struct {
	Token      string `json:"Token"`
	BotPrefix  string `json:"BotPrefix"`
	MCServer   string `json:"MCServer"`
	MCPort     string `json:"MCPort"`
	WeatherAPI string `json:"WeatherAPI"`
}

func ReadConfig() error {
	fmt.Println("Reading Config File")

	file, err := ioutil.ReadFile("./config.json")
	if err != nil {
		fmt.Println("Error Reading Config")
		return err
	}
	fmt.Println(string(file))

	err = json.Unmarshal(file, &config)
	if err != nil {
		fmt.Println("Error converting config to vars")
		fmt.Println(err)
		return err
	}

	Token = config.Token
	BotPrefix = config.BotPrefix
	MCServer = config.MCServer
	MCPort = config.MCPort
	WeatherAPI = config.WeatherAPI

	return nil

}
