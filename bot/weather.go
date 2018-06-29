package bot

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"strings"

	"github.com/monkeyworknet/monkeybot/config"
)

/*

Project Notes

	Dark Sky:       https://api.darksky.net/forecast/[key]/[latitude],[longitude]
	Google Maps API:   https://maps.googleapis.com/maps/api/geocode/json?parameters
	JSON Converter:   https://mholt.github.io/json-to-go/


	Testing for St. Michaels:    https://maps.googleapis.com/maps/api/geocode/json?address=st.+michaels+md&key=xxxxxxxxxxx
	               "lat" : 38.785393,
					"lng" : -76.22332020000002
					"formatted_address" : "St Michaels, MD 21663, USA",

					https://api.darksky.net/forecast/xxxxxxxxxxxxxx/38.78,-76.22?exclude=minutely,hourly,daily,flags
					"time":1530246455,
					"summary":"Clear"
					"icon":"clear-night"
					"temperature":76.79,
					"apparentTemperature":77.84

					Icon optional
							A machine-readable text summary of this data point, suitable for selecting an icon for display.
							If defined, this property will have one of the following values:
							clear-day,   :sunny:
							clear-night,  :full_moon:
							rain, :droplet:
							snow,:snowman:
							sleet, :cloud_snow:
							wind, :wind_blowing_face:
							fog,  :fog:
							cloudy, :cloud:
							partly-cloudy-day, :white_sun_cloud:
							partly-cloudy-night.
							(Developers should ensure that a sensible default is defined, as additional values, such as hail, thunderstorm, or tornado, may be defined in the future.)
*/

var darksky struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Timezone  string  `json:"timezone"`
	Currently struct {
		Time                 int     `json:"time"`
		Summary              string  `json:"summary"`
		Icon                 string  `json:"icon"`
		NearestStormDistance int     `json:"nearestStormDistance"`
		NearestStormBearing  int     `json:"nearestStormBearing"`
		PrecipIntensity      int     `json:"precipIntensity"`
		PrecipProbability    int     `json:"precipProbability"`
		Temperature          float64 `json:"temperature"`
		ApparentTemperature  float64 `json:"apparentTemperature"`
		DewPoint             float64 `json:"dewPoint"`
		Humidity             float64 `json:"humidity"`
		Pressure             float64 `json:"pressure"`
		WindSpeed            float64 `json:"windSpeed"`
		WindGust             float64 `json:"windGust"`
		WindBearing          int     `json:"windBearing"`
		CloudCover           int     `json:"cloudCover"`
		UvIndex              int     `json:"uvIndex"`
		Visibility           float64 `json:"visibility"`
		Ozone                float64 `json:"ozone"`
	} `json:"currently"`
	Offset int `json:"offset"`
}

var googlemaps struct {
	Results []struct {
		AddressComponents []struct {
			LongName  string   `json:"long_name"`
			ShortName string   `json:"short_name"`
			Types     []string `json:"types"`
		} `json:"address_components"`
		FormattedAddress string `json:"formatted_address"`
		Geometry         struct {
			Bounds struct {
				Northeast struct {
					Lat float64 `json:"lat"`
					Lng float64 `json:"lng"`
				} `json:"northeast"`
				Southwest struct {
					Lat float64 `json:"lat"`
					Lng float64 `json:"lng"`
				} `json:"southwest"`
			} `json:"bounds"`
			Location struct {
				Lat float64 `json:"lat"`
				Lng float64 `json:"lng"`
			} `json:"location"`
			LocationType string `json:"location_type"`
			Viewport     struct {
				Northeast struct {
					Lat float64 `json:"lat"`
					Lng float64 `json:"lng"`
				} `json:"northeast"`
				Southwest struct {
					Lat float64 `json:"lat"`
					Lng float64 `json:"lng"`
				} `json:"southwest"`
			} `json:"viewport"`
		} `json:"geometry"`
		PlaceID string   `json:"place_id"`
		Types   []string `json:"types"`
	} `json:"results"`
	Status string `json:"status"`
}

func Weather(command []string) (string, error) {

	GoogleAPIkey := config.GoogleAPIkey
	DarkSkyAPIkey := config.DarkSkyAPIkey

	// take input from user, drop the weather command and format the rest as a search query for google.
	//command := []string{"!weather", "oshawa", "on"}
	command = append(command[:0], command[1:]...)
	city := strings.Join(command, "+")

	if city == "" {
		fmt.Println("Please try again:  !weather city, state")
	}

	googleurl := fmt.Sprintf("https://maps.googleapis.com/maps/api/geocode/json?address=%v&key=%v", city, GoogleAPIkey)

	fmt.Println("Fetching Lat/Long From Google")
	fmt.Println(googleurl)

	resp, err := http.Get(googleurl)
	if err != nil {
		fmt.Println("Unable to reach google location services", err)
		return "Unable to reach google location services", err
	}

	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&googlemaps); err != nil {
		fmt.Println("Error in decoding google location", err)
	}

	fmt.Println("Returns: "+googlemaps.Results[0].FormattedAddress, googlemaps.Results[0].Geometry.Location.Lat, googlemaps.Results[0].Geometry.Location.Lng)
	address := googlemaps.Results[0].FormattedAddress
	lat := googlemaps.Results[0].Geometry.Location.Lat
	lng := googlemaps.Results[0].Geometry.Location.Lng

	// take these results and format a darksy query

	darkskyurl := fmt.Sprintf("https://api.darksky.net/forecast/%v/%v,%v?exclude=minutely,hourly,daily,flags", DarkSkyAPIkey, lat, lng)

	fmt.Println("Fetching Weather from Darksky")
	fmt.Println(darkskyurl)

	resp1, err := http.Get(darkskyurl)
	if err != nil {
		fmt.Println("Unable to reach Darksky Weather services", err)
		return "Unable to reach Darksky Weather services", err
	}

	defer resp1.Body.Close()

	if err := json.NewDecoder(resp1.Body).Decode(&darksky); err != nil {
		fmt.Println("Error in decoding Darksky Weather", err)
	}

	fmt.Println("Returns: "+darksky.Currently.Summary, darksky.Currently.Temperature, darksky.Currently.ApparentTemperature, address)

	currenttempc := math.Round((darksky.Currently.Temperature - 32) * 0.5556)
	feeltempc := math.Round((darksky.Currently.ApparentTemperature - 32) * 0.5556)

	// attempt icon using emoji's in discord

	conditions := []string{"clear-day", "clear-night", "rain", "snow", "sleet", "wind", "fog", "cloudy", "partly-cloudy-day"}
	icons := []string{":sunny:", ":full_moon:", ":droplet:", ":snowman:", ":cloud_snow:", ":wind_blowing_face:", ":fog:", ":cloud:", ":white_sun_cloud:"}
	myicon := ""

	for i, v := range conditions {
		if v == darksky.Currently.Icon {
			myicon = icons[i]
		}
	}

	// final output

	output := fmt.Sprintf("Current weather for %v is described as %v %v\nIt is currently %v C / %v F and feels like %v C / %v F", address, myicon, darksky.Currently.Summary, currenttempc, math.Round(darksky.Currently.Temperature), feeltempc, math.Round(darksky.Currently.ApparentTemperature))

	fmt.Println(output)
	return output, nil

}
