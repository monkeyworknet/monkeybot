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
	playerID      string `json:"playerID"`
	playerName    string `json:"playerName"`
	totalGuessed  int    `json:"totalGuessed"`
	currentPoints int    `json:"currentPoints"`
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

		playerwrite := PlayersDB{playerID: playerID, playerName: playerName, totalGuessed: totalGuessed, currentPoints: currentPoints}
		if err := db.Write(config.DatabaseName, senderid, playerwrite); err != nil {
			fmt.Printf("Error - Couldn't create db entry for %v - %v", senderid, err)
			return false
		}
		fmt.Println(playerwrite)
		return true

	}

	// Update values

	playerread.playerID = senderid
	playerread.playerName = sendername
	playerread.totalGuessed = playerread.totalGuessed + 1
	playerread.currentPoints = playerread.currentPoints + score
	playerwrite := PlayersDB{playerID: playerread.playerID, playerName: playerread.playerName, totalGuessed: playerread.totalGuessed, currentPoints: playerread.currentPoints}
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

	return playerread.currentPoints

}

func answer(question question, answer []string, senderid string, sendername string) (question, string) {

	score := 0
	formatedanswer := append(answer[:0], answer[1:]...)
	givenanswer := strings.Join(formatedanswer, " ")
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
			score = 3
		}

		if currentq.difficulty == "medium" {
			score = 2
		}
		if currentq.difficulty == "easy" {
			score = 1
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
		score = -1
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
