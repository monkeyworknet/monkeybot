package bot

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"strings"

	"github.com/mndrix/rand"
)

func answer(question question, answer []string) (question, string) {

	formatedanswer := append(answer[:0], answer[1:]...)
	givenanswer := strings.Join(formatedanswer, " ")
	currentq = question
	correctanswer := strings.ToLower(currentq.correct)
	givenanswer = strings.ToLower(givenanswer)

	newquestionplease := "This Question has already been answered, please ask a new one"
	answeredcorrect := "That's correct, the right answer was " + correctanswer
	answeredwrong := fmt.Sprintf("Sorry %v is Wrong,  the right answer was %v", givenanswer, correctanswer)

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
