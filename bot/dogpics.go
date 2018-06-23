package bot

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func DogPics(breed string) (string, error) {


// https://dog.ceo/api/breed/hound/images/random    
// https://dog.ceo/api/breeds/image/random
	siteurl := ""

	if breed == "empty" {
		siteurl = "https://dog.ceo/api/breeds/image/random"
		} else {
		 siteurl = "https://dog.ceo/api/breed/" + breed + "/images/random"
		} 

	fmt.Println(siteurl)

	fmt.Println("Fetching a dog pic",  breed)
	resp, err := http.Get(siteurl)

	if err != nil {
		println(err)
		return "", err
	}

	defer resp.Body.Close()

	var pic struct {
		Url string `json:"message"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&pic); err != nil {

		return "", err
	}

	return pic.Url, nil

}
