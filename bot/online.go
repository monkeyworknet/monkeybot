package bot

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func WhoIsOnline() ([]string, int, error) {

	servername := "play.crafttheory.net"
	serverport := "25565"
	empty := []string{"empty result"}
	serverempty := []string{"Nobody is online"}

	// Check to see if anyone is online before we do anything

        resp, err := http.Get("https://api.minetools.eu/query/" + servername + "/" + serverport)

        if err != nil {
                fmt.Println("FATAL Error getting Playerlist (minetools down?)  -  ", err)
                return empty, 0,  err
        }


	var NumberPlayers struct {
		Playercount int `json:"Players"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&NumberPlayers); err != nil {
		fmt.Println("Error Decoding PlayerCountlist  -  ", err)
		return empty, 0,  err
	}

	if NumberPlayers.Playercount == 0 {
		fmt.Println("No One Online")
		return serverempty, 0,  nil
	}

	// Grab the player list

	resp, err = http.Get("https://api.minetools.eu/query/" + servername + "/" + serverport)

	if err != nil {
		fmt.Println("FATAL Error getting Playerlist (minetools down?)  -  ", err)
		return empty, 0, err
	}

	defer resp.Body.Close()

	var ActivePlayers struct {
		Playerlist  []string `json:"Playerlist"`
		Playercount int      `json:"Players"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&ActivePlayers); err != nil {
		fmt.Println("Error Decoding Playerlist  -  ", err)
		return empty, 0, err
	}

	if ActivePlayers.Playercount > 0 {
		return ActivePlayers.Playerlist, ActivePlayers.Playercount, nil
	} else {
		return serverempty, 0, nil
	}

}
