/*

{"coord":{"lon":-79.39,"lat":43.65},"weather":[{"id":701,"main":"Mist","description":"mist","icon":"50n"}],"base":"stations","main":{"temp":16.99,"pressure":1008,"humidity":88,"temp_min":16,"temp_max":18},"visibility":9656,"wind":{"speed":2.6,"deg":330},"clouds":{"all":90},"dt":1529816460,"sys":{"type":1,"id":2117,"message":0.0054,"country":"CA","sunrise":1529833017,"sunset":1529888589},"id":6167865,"name":"Toronto","cod":200}

*/

package bot

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

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
	var UrlString string

	command = append(command[:0], command[1:]...)
	city := strings.Join(command, "%20")
	UrlString = "https://api.openweathermap.org/data/2.5/weather?q=" + city + "&type=accurate&units=metric&appid=" + config.WeatherAPI

	fmt.Println("Fetching Weather")
	fmt.Println(UrlString)

	resp, err := http.Get(UrlString)
	if err != nil {
		fmt.Println("Unable to reach weather services", err)
		return "Unable to reach weather services", err
	}

	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&WeatherResponse); err != nil {
		fmt.Println("Error in decoding weather... maybe not found?", err)
	}

	if WeatherResponse.Cod == 0 {
		fmt.Println("City Not Found")
		return "City Not Found, try !weatherzip <zipcode>", nil
	}

	fmt.Println(WeatherResponse.Name, WeatherResponse.Sys.Country, WeatherResponse.Main.Temp, WeatherResponse.Weather[0].Description)

	ftemp := (WeatherResponse.Main.Temp * 9 / 5) + 32
	reportedtemp := strconv.FormatFloat(WeatherResponse.Main.Temp, 'f', 1, 64) + " C / " + strconv.FormatFloat(ftemp, 'f', 1, 64) + " F"

	returnstring := WeatherResponse.Name + "'s weather is described as " + WeatherResponse.Weather[0].Description + ".   The Temp is currently " + reportedtemp

	return returnstring, nil

}

func WeatherZip(command []string) (string, error) {

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
			TempMin  float64 `json:"temp_min"`
			TempMax  float64 `json:"temp_max"`
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

	var UrlString string
	command = append(command[:0], command[1:]...)
	city := strings.Join(command, "%20")
	UrlString = "https://api.openweathermap.org/data/2.5/weather?zip=" + city + "&units=metric&appid=" + config.WeatherAPI

	fmt.Println("Fetching Weather via Zip")
	fmt.Println(UrlString)

	resp, err := http.Get(UrlString)
	if err != nil {
		fmt.Println("Unable to reach weather services", err)
		return "Unable to reach weather services", err
	}

	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&WeatherResponse); err != nil {
		fmt.Println("Error in decoding weather... maybe not found?", err)
	}

	if WeatherResponse.Cod == 0 {
		fmt.Println("Zip Not Found")
		return "Zip Not Found, currently only supporting USA", nil

	}

	fmt.Println(WeatherResponse.Name, WeatherResponse.Sys.Country, WeatherResponse.Main.Temp, WeatherResponse.Weather[0].Description)

	ftemp := (WeatherResponse.Main.Temp * 9 / 5) + 32
	reportedtemp := strconv.FormatFloat(WeatherResponse.Main.Temp, 'f', 1, 64) + " C / " + strconv.FormatFloat(ftemp, 'f', 1, 64) + " F"

	returnstring := WeatherResponse.Name + "'s weather is described as " + WeatherResponse.Weather[0].Description + ".   The Temp is currently " + reportedtemp

	return returnstring, nil

}
