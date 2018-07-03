package bot

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"strings"
	"time"

	"github.com/mndrix/rand"
	"github.com/monkeyworknet/monkeybot/config"
	"github.com/nanobox-io/golang-scribble"
)

type PlayersDB struct {
	Playerid      string `json:"playerid"`
	Playername    string `json:"playername"`
	Totalguessed  int    `json:"totalguessed"`
	Currentpoints int    `json:"currentpoints"`
}

func updatescore(score int, senderid string, sendername string) bool {

	db, err := scribble.New(config.DatabasePath, nil)
	if err != nil {
		fmt.Println("FATAL Error creating db", err)
		return false
	}

	playerread := PlayersDB{}

	// End of DB Trivia Setup

	// Read in current values from DB

	if err := db.Read(config.DatabaseName, senderid, &playerread); err != nil {
		fmt.Printf("%v not found, creating new entry for them", sendername)
		// Default Values if user doesn't exist
		playerID := senderid
		playerName := sendername
		totalGuessed := 1
		currentPoints := score

		//		playerwrite := PlayersDB{playerID: playerID, playerName: playerName, totalGuessed: totalGuessed, currentPoints: currentPoints}
		playerwrite := PlayersDB{Playerid: playerID, Playername: playerName, Totalguessed: totalGuessed, Currentpoints: currentPoints}

		if err := db.Write(config.DatabaseName, senderid, playerwrite); err != nil {
			fmt.Printf("Error - Couldn't create db entry for %v - %v", senderid, err)
			return false
		}
		fmt.Println(playerwrite)
		return true

	}

	// Update values

	playerread.Playerid = senderid
	playerread.Playername = sendername
	playerread.Totalguessed = playerread.Totalguessed + 1
	playerread.Currentpoints = playerread.Currentpoints + score
	playerwrite := PlayersDB{Playerid: playerread.Playerid, Playername: playerread.Playername, Totalguessed: playerread.Totalguessed, Currentpoints: playerread.Currentpoints}

	//playerwrite := PlayersDB{playerID: playerread.playerID, playerName: playerread.playerName, totalGuessed: playerread.totalGuessed, currentPoints: playerread.currentPoints}
	if err := db.Write(config.DatabaseName, senderid, playerwrite); err != nil {
		fmt.Printf("Error - Couldn't create db entry for %v - %v", senderid, err)
		return false
	}
	fmt.Println(playerwrite)
	return true

}

func readscore(senderid string) int {

	db, err := scribble.New(config.DatabasePath, nil)
	if err != nil {
		fmt.Println("FATAL Error creating db", err)
		return 0
	}

	playerread := PlayersDB{}

	if err := db.Read(config.DatabaseName, senderid, &playerread); err != nil {
		fmt.Printf("%v not found in DB", senderid)
		return 0
	}

	return playerread.Currentpoints

}

func answer(question question, answer []string, senderid string, sendername string) (question, string) {

	score := 0
	formatedanswer := append(answer[:0], answer[1:]...)
	givenanswer := strings.Join(formatedanswer, " ")
	givenanswer = strings.TrimSpace(givenanswer)
	currentq = question
	correctanswer := strings.ToLower(currentq.correct)
	givenanswer = strings.ToLower(givenanswer)
	timeout := float64(35)
	timesince := time.Since(currentq.time)

	newquestionplease := "This Question has already been answered, please ask a new one"
	answeredcorrect := fmt.Sprintf("Congrats!  %v was the correct answer", correctanswer)
	answeredwrong := fmt.Sprintf("Sorry %v is incorrect.  The right answer was %v", givenanswer, correctanswer)
	timedout := fmt.Sprintf("Sorry the time has expired for this question.  Please request a new question.  The Correct answer was %v", correctanswer)

	if timesince.Seconds() > timeout {
		currentq.answered = true
		return currentq, timedout
	}

	if currentq.answered {
		return currentq, newquestionplease
	}

	if givenanswer == correctanswer {
		currentq.answered = true

		if currentq.difficulty == "hard" {
			score = 10
		}

		if currentq.difficulty == "medium" {
			score = 5
		}
		if currentq.difficulty == "easy" {
			score = 3
		}

		res := updatescore(score, senderid, sendername)
		fmt.Println(res)
		score = readscore(senderid)
		answeredcorrect = fmt.Sprintf("%v \nYou're Current Score is %v", answeredcorrect, score)

		return currentq, answeredcorrect

	}

	if currentq.difficulty == "hard" {
		score = -1
	}

	if currentq.difficulty == "medium" {
		score = -2
	}
	if currentq.difficulty == "easy" {
		score = -3
	}

	res := updatescore(score, senderid, sendername)
	fmt.Println(res)
	score = readscore(senderid)
	answeredwrong = fmt.Sprintf("%v \nYou're Current Score is %v", answeredwrong, score)

	currentq.answered = true
	return currentq, answeredwrong

}

func ask() (question, error) {

	var questionlist opentdb

	fmt.Println("Fetching 10 random questions")
	fmt.Println(questionurl)
	rannum := rand.Intn(10)
	resp, err := http.Get(questionurl)
	if err != nil {
		fmt.Println("Unable to reach trivia services", err)
		return currentq, err
	}

	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&questionlist); err != nil {
		fmt.Println("Error in decoding trivia info", err)
	}
	choices := append(questionlist.Results[rannum].IncorrectAnswers, questionlist.Results[rannum].CorrectAnswer)

	sort.Strings(choices)
	currentq = question{questionlist.Results[rannum].Category, questionlist.Results[rannum].Difficulty, questionlist.Results[rannum].Question, questionlist.Results[rannum].CorrectAnswer, choices, false, time.Now()}

	fmt.Println(currentq.answered)
	return currentq, nil
}

func highscore() (response string) {

	db, err := scribble.New(config.DatabasePath, nil)
	if err != nil {
		fmt.Println("FATAL Error creating db", err)
	}

	type kv struct {
		Key   string
		Value int
	}

	records, err := db.ReadAll(config.DatabaseName)
	if err != nil {
		fmt.Println(err)
	}

	var ss []kv
	// broken out all the records into both a slice set and a map for testing
	for _, p := range records {
		player := PlayersDB{}
		if err := json.Unmarshal([]byte(p), &player); err != nil {
			fmt.Println(err)
		}
		ss = append(ss, kv{player.Playerid, player.Currentpoints})
	}

	fmt.Println(ss)

	// sorting out the slice set

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})

	for _, kv := range ss {
		fmt.Printf("%s, %d\n", kv.Key, kv.Value)
	}

	// High Score Board will be Top 5, or fewer if there are less players
	top := 5
	if len(ss) < top {
		top = len(ss)
	}
	fmt.Println(top)

	highscorelist := "***Trivia Leader Board.***\n(#) = Times played. \n\n"

	for i := 0; i < (top); i++ {
		fmt.Println(i)
		playerread := PlayersDB{}
		if err := db.Read(config.DatabaseName, ss[(i)].Key, &playerread); err != nil {
			fmt.Printf("%v not found", ss[(i)].Key)
		}

		lineitem := fmt.Sprintf("***%v*** (%v) \t %v\n", playerread.Currentpoints, playerread.Totalguessed, playerread.Playername)
		highscorelist = highscorelist + lineitem

	}

	fmt.Println(highscorelist)
	return highscorelist

}
