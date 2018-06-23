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

        var joke []string

        _ = json.NewDecoder(resp.Body).Decode(&joke)

	return joke[0], nil

}
