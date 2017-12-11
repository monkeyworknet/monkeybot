package bot

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Trbmb() (string, error) {

	fmt.Println("Fetching a TRBMB Joke")
	resp, err := http.Get("http://api.chew.pro/trbmb")
	if err != nil {
		println(err)
		return "", err
	}

	defer resp.Body.Close()

	// returning a json Array ... how to deal???

	var joke struct {
		Joke string `""`
	}

	if err := json.NewDecoder(resp.Body).Decode(&joke); err != nil {
		return "", err
	}

	return joke.Joke, nil

}
