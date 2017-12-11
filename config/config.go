package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var (
	Token     string
	BotPrefix string

	config *configStruct
)

type configStruct struct {
	Token     string `json:"Token"`
	BotPrefix string `json:"BotPrefix"`
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

	return nil

}
