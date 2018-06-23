package bot

import (
	"fmt"
	"net/http"
)

func CatPics() (string, error) {

	fmt.Println("Fetching a cat pic")

	UrlString := "http://thecatapi.com/api/images/get?format=src&results_per_page=1"

	resp, err := http.Get(UrlString)
	if err != nil {
		println(err)
		return "error grabbing cat pic", err
	}
	defer resp.Body.Close()

	finalURL := resp.Request.URL.String()
	fmt.Println(finalURL)

	return finalURL, nil

}
