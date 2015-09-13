package shared

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
Question is a struct for variable user input.
*/
type Question struct {
	question      string
	answers       []Answer
	caseSensitive bool
}

/*
Answer is a possible requested answer to a question.
*/
type Answer struct {
	option int
	valid  []string
}

/*
CreateQuestion creates a new question object with the given question. Per default
the answers are not case sensitive.
*/
func CreateQuestion(questionText string) *Question {
	return &Question{question: questionText, caseSensitive: false}
}

/*
CreateYesNo creates a question with predeclared Yes and No answer options. The
Option for No is negative, for Yes positive.
*/
func CreateYesNo(questionText string) *Question {
	question := &Question{question: "(Y/N) " + questionText, caseSensitive: false}
	question.CreateAnswer(-1, "n", "no")
	question.CreateAnswer(1, "y", "yes")
	return question
}

/*
Ask the question and returns the option value of the chosen answer.
*/
func (q *Question) Ask() int {
	// prepare console reader
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(q.question)
	// keep asking the question until we get an answer
	for {
		// read input
		input, _ := reader.ReadString('\n')
		input = strings.Trim(input, "\n")
		// check if one of the answers has been chosen
		for _, a := range q.answers {
			for _, value := range a.valid {
				if q.caseSensitive && value == input {
					// if case sensitive it must be an exact match
					return a.option
				} else if !q.caseSensitive && strings.EqualFold(value, input) {
					// if not case sensitive use EqualFold
					return a.option
				}
			} // end answer check
		} // end all answers check: if we reach this no legal answer was given
		fmt.Println("Invalid reply!\n", q.question)
	}
}

/*
CreateAnswer creates an answer for the given question.
*/
func (q *Question) CreateAnswer(option int, valid ...string) {
	q.answers = append(q.answers, Answer{option: option, valid: valid})
}
