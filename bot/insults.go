package bot

import (
	"encoding/xml"
	"fmt"
	"net/http"
)

func Insults() (string, error) {

	fmt.Println("Fetching an Insult")

	resp, err := http.Get("http://www.dickless.org/api/insult.xml")

	if err != nil {
		println(err)
		return "", err
	}
	defer resp.Body.Close()

	fmt.Println(resp.Body)

	var joke struct {
		Joke string `xml:"insult"`
	}

	// 5 minute cache on all responses here ...
	if err := xml.NewDecoder(resp.Body).Decode(&joke); err != nil {
		return "", err
	}

	return joke.Joke, nil

}
