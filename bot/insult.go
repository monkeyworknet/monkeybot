package bot

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func PersonalAttack(command []string) (string, error) {

	UrlString := "https://insult.mattbas.org/api/en/insult.json"

	fmt.Println("Fetching a Generic Insult")
	resp, err := http.Get(UrlString)
	if err != nil {
		println(err)
		return "", err
	}

	defer resp.Body.Close()

	var insult struct {
		Insult string `json:"insult"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&insult); err != nil {
		return "", err
	}

	responseString := insult.Insult

	if len(command) > 1 {
		fmt.Println("Personalizing Insult")
		responseString = command[1] + ": " + string(responseString)
	}

	return responseString, nil

}
