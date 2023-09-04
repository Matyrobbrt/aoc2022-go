package main

import (
	"github.com/AlecAivazis/survey/v2"
	"golearning/aoc"
	"strconv"
)

var days = []func(){
	aoc.Day1, aoc.Day2, aoc.Day3, aoc.Day4, aoc.Day5,
}

func main() {
	var selection string
	survey.AskOne(&survey.Select{
		Message: "Which challenge would you like to run?",
		Options: func() []string {
			strings := make([]string, len(days))
			for i := 0; i < len(days); i++ {
				strings[i] = strconv.Itoa(i + 1)
			}
			return strings
		}(),
	}, &selection, survey.WithValidator(survey.Required))
	days[aoc.ToInt(selection)-1]()
}
