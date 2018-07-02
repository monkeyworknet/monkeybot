package bot

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"

	"github.com/mndrix/rand"
)

func answer(question question, answer []string) (question, string) {

	currentq = question
	correctanswer := currentq.correct
	givenanswer := answer[1]

	newquestionplease := "This Question has already been answered, please ask a new one"
	answeredcorrect := "That's correct, the right answer was " + correctanswer
	answeredwrong := "Sorry " + givenanswer + "is Wrong,  the right answer was " + correctanswer

	if currentq.answered {
		return currentq, newquestionplease
	}

	if givenanswer == correctanswer {
		currentq.answered = true
		return currentq, answeredcorrect
	}

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
	currentq = question{questionlist.Results[rannum].Category, questionlist.Results[rannum].Difficulty, questionlist.Results[rannum].Question, questionlist.Results[rannum].CorrectAnswer, choices, false}

	fmt.Println(currentq.answered)
	return currentq, nil
}
