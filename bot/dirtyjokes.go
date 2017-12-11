package bot

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func DirtyJokes() (string, error) {

	fmt.Println("Fetching a Dirty Joke")
	resp, err := http.Get("https://crackmeup-api.herokuapp.com/dirty")
	if err != nil {
		println(err)
		return "", err
	}

	defer resp.Body.Close()

	var joke struct {
		Joke string `json:"joke"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&joke); err != nil {
		return "", err
	}

	return joke.Joke, nil

}
