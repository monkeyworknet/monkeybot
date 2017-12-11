package bot

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ChuckJoke() (string, error) {

	fmt.Println("Fetching a CN Joke")
	resp, err := http.Get("http://api.icndb.com/jokes/random")
	if err != nil {
		println(err)
		return "", err
	}

	defer resp.Body.Close()

	var joke struct {
		Value struct {
			Joke string `json:"joke"`
		} `json:"value"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&joke); err != nil {
		return "", err
	}

	return joke.Value.Joke, nil

}
