package bot

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"strings"
	"time"

	"github.com/mndrix/rand"
)

func answer(question question, answer []string) (question, string) {

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
	currentq = question{questionlist.Results[rannum].Category, questionlist.Results[rannum].Difficulty, questionlist.Results[rannum].Question, questionlist.Results[rannum].CorrectAnswer, choices, false, time.Now()}

	fmt.Println(currentq.answered)
	return currentq, nil
}
