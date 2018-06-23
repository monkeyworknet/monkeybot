package bot

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func weather(command []string) (string, error) {

	fmt.Println(command[1])

	UrlString := "https://api.openweathermap.org/data/2.5/weather?q=" + command[1] + "&type=like&units=metric&appid=f761c4019670c58208bf0a58c53514da"

	fmt.Println(UrlString)

	fmt.Println("Fetching Weather")
	resp, err := http.Get(UrlString)
	if err != nil {
		println(err)
		return "", err
	}

	defer resp.Body.Close()

	var openweather struct {
		city  string `json:"name"`
		wtype string `json:"weather.main"`
		temp  string `json:"main.temp"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&openweather); err != nil {
		return "", err
	}
	fmt.Println(openweather.city)
	returnstring := "The weather in " + openweather.city + " is " + openweather.wtype + ".   The Temp is " + openweather.temp + " Canadian"

	return returnstring, nil

}
