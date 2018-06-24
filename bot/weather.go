/*

{"coord":{"lon":-79.39,"lat":43.65},"weather":[{"id":701,"main":"Mist","description":"mist","icon":"50n"}],"base":"stations","main":{"temp":16.99,"pressure":1008,"humidity":88,"temp_min":16,"temp_max":18},"visibility":9656,"wind":{"speed":2.6,"deg":330},"clouds":{"all":90},"dt":1529816460,"sys":{"type":1,"id":2117,"message":0.0054,"country":"CA","sunrise":1529833017,"sunset":1529888589},"id":6167865,"name":"Toronto","cod":200}

*/

package bot

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/monkeyworknet/monkeybot/config"
)

//https://mholt.github.io/json-to-go/

func Weather(command []string) (string, error) {

	var WeatherResponse struct {
		Coord struct {
			Lon float64 `json:"lon"`
			Lat float64 `json:"lat"`
		} `json:"coord"`
		Weather []struct {
			ID          int    `json:"id"`
			Main        string `json:"main"`
			Description string `json:"description"`
			Icon        string `json:"icon"`
		} `json:"weather"`
		Base string `json:"base"`
		Main struct {
			Temp     float64 `json:"temp"`
			Pressure int     `json:"pressure"`
			Humidity int     `json:"humidity"`
			TempMin  int     `json:"temp_min"`
			TempMax  int     `json:"temp_max"`
		} `json:"main"`
		Visibility int `json:"visibility"`
		Wind       struct {
			Speed float64 `json:"speed"`
			Deg   int     `json:"deg"`
		} `json:"wind"`
		Clouds struct {
			All int `json:"all"`
		} `json:"clouds"`
		Dt  int `json:"dt"`
		Sys struct {
			Type    int     `json:"type"`
			ID      int     `json:"id"`
			Message float64 `json:"message"`
			Country string  `json:"country"`
			Sunrise int     `json:"sunrise"`
			Sunset  int     `json:"sunset"`
		} `json:"sys"`
		ID   int    `json:"id"`
		Name string `json:"name"`
		Cod  int    `json:"cod"`
	}

	fmt.Println(command[1])
	if command[1] == "" {
		return "Sorry Please Specify a City", nil
	}

	UrlString := "https://api.openweathermap.org/data/2.5/weather?q=" + command[1] + "&type=like&units=metric&appid=" + config.WeatherAPI

	fmt.Println("Fetching Weather")
	fmt.Println(UrlString)

	resp, err := http.Get(UrlString)
	if err != nil {
		println(err)
		return "", err
	}

	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&WeatherResponse); err != nil {
		return "", err
	}
	fmt.Println(WeatherResponse.Name, WeatherResponse.Sys.Country, WeatherResponse.Main.Temp, WeatherResponse.Weather[0].Description)

	reportedtemp := strconv.FormatFloat(WeatherResponse.Main.Temp, 'E', -1, 64) + " C"
	if WeatherResponse.Sys.Country == "US" {
		ftemp := (WeatherResponse.Main.Temp * 9 / 5) + 32
		reportedtemp = strconv.FormatFloat(ftemp, 'E', -1, 64) + " F"
	}

	returnstring := "The weather in " + WeatherResponse.Name + " is " + WeatherResponse.Weather[0].Description + ".   The Temp is " + reportedtemp

	return returnstring, nil

}
