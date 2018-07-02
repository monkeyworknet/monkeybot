package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var (
	Token         string
	BotPrefix     string
	MCServer      string
	MCPort        string
	GoogleAPIkey  string
	DarkSkyAPIkey string
	DatabasePath  string
	DatabaseName  string

	config *configStruct
)

type configStruct struct {
	Token         string `json:"Token"`
	BotPrefix     string `json:"BotPrefix"`
	MCServer      string `json:"MCServer"`
	MCPort        string `json:"MCPort"`
	GoogleAPIkey  string `json:"GoogleAPIkey"`
	DarkSkyAPIkey string `json:"DarkSkyAPIkey"`
	DatabasePath  string `json:"DatabasePath"`
	DatabaseName  string `json:"DatabaseName`
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

	GoogleAPIkey = config.GoogleAPIkey
	DarkSkyAPIkey = config.DarkSkyAPIkey

	DatabaseName = config.DatabaseName
	DatabasePath = config.DatabasePath

	return nil

}
