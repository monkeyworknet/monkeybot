package bot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func DadJokes() (string, error) {

	fmt.Println("Fetching a Dad Joke")

	client := &http.Client{}
	postData := make([]byte, 100)
	req, err := http.NewRequest("GET", "https://icanhazdadjoke.com/", bytes.NewReader(postData))

	if err != nil {
		println(err)
		return "", err
	}
	req.Header.Add("Accept", "application/json")
	resp, err := client.Do(req)
	defer resp.Body.Close()

	fmt.Println(resp.Body)

	var joke struct {
		Joke string `json:"joke"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&joke); err != nil {
		return "", err
	}

	return joke.Joke, nil

}
