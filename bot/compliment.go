package bot

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func Compliment(command []string) (string, error) {

	UrlString := "https://compliment-api.herokuapp.com"

	fmt.Println("Fetching a Generic compliment")
	resp, err := http.Get(UrlString)
	if err != nil {
		println(err)
		return "", err
	}

	defer resp.Body.Close()

	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	responseString := string(responseData)

	if len(command) > 1 {
		fmt.Println("Personalizing Compliment")
		responseString = command[1] + ": " + string(responseData)
	}

	return responseString, nil

}
