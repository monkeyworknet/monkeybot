package bot

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func MommaJokes() (string, error) {

	fmt.Println("Fetching a Momma Joke")
	resp, err := http.Get("http://api.yomomma.info/")
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
