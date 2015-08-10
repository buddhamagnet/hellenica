package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var words = map[string]string{
	"and":                    "keh",
	"goodbye":                "andio",
	"greek":                  "ellenica",
	"hello informal":         "yassu",
	"hello formal":           "yessas",
	"I do not understand":    "ven kalataveno",
	"I speak a little Greek": "meelow liga ellenica",
	"no":        "ochi",
	"please":    "parakalo",
	"thank you": "efkaristo",
	"yes":       "ne",
	"I'm sorry": "signomi",
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		question, answer := getQuestion()
		fmt.Printf("what is the phonetic for %s?\n", question)
		text, _ := reader.ReadString('\n')
		data := strings.TrimSpace(text)
		switch data {
		case "Q", "q":
			fmt.Println("αντίο!")
			os.Exit(0)
		case answer:
			fmt.Println("\acorrect!")
		default:
			fmt.Printf("wrong it is %s\n", words[question])
		}
	}
}

func getQuestion() (question, answer string) {
	for q, a := range words {
		question, answer = q, a
		break
	}
	return question, answer
}
