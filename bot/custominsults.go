package bot

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func PersonalAttack() (string, error) {

	fmt.Println("Fetching a Personalized Insult")
	resp, err := http.Get("https://insult.mattbas.org/api/en/insult.json")
	if err != nil {
		println(err)
		return "", err
	}

	defer resp.Body.Close()

	var joke struct {
		Joke string `json:"insult"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&joke); err != nil {
		return "", err
	}

	return joke.Joke, nil

}
