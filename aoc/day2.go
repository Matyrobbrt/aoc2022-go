package aoc

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Outcome int8

const (
	win  Outcome = 6
	draw Outcome = 3
	lose Outcome = 0
)

type RoundOutcome struct {
	Rock     Outcome
	Paper    Outcome
	Scissors Outcome
}

type Move struct {
	Score    int
	Outcomes *RoundOutcome
}

var rock = Move{Score: 1, Outcomes: &RoundOutcome{Rock: draw, Paper: lose, Scissors: win}}
var paper = Move{Score: 2, Outcomes: &RoundOutcome{Rock: win, Paper: draw, Scissors: lose}}
var scissors = Move{Score: 3, Outcomes: &RoundOutcome{Rock: lose, Paper: win, Scissors: draw}}

// The outcome of the match for the given play of the opponent
func (receiver Move) getOutcome(opponent Move) Outcome {
	switch opponent {
	case paper:
		return receiver.Outcomes.Paper
	case rock:
		return receiver.Outcomes.Rock
	case scissors:
		return receiver.Outcomes.Scissors
	}
	return 0
}

func getMove(value, rockMove, paperMove, scissorsMove string) Move {
	switch value {
	case rockMove:
		return rock
	case paperMove:
		return paper
	case scissorsMove:
		return scissors
	}
	return Move{}
}

func (receiver Move) getMoveNeededForOutcome(outcome Outcome) Move {
	for _, v := range []struct {
		outcome Outcome
		move    Move
	}{{receiver.Outcomes.Rock, rock}, {receiver.Outcomes.Paper, paper}, {receiver.Outcomes.Scissors, scissors}} {
		if v.outcome == outcome {
			return v.move
		}
	}
	return Move{}
}

func (receiver Outcome) reverse() Outcome {
	switch receiver {
	case draw:
		return draw
	case lose:
		return win
	default:
		return lose
	}
}

func Day2() {
	file, _ := os.Open("input/aocday2.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	myScore := 0
	myScoreWithKnownOutcome := 0
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " ")

		them := getMove(split[0], "A", "B", "C")
		me := getMove(split[1], "X", "Y", "Z")
		myScore += me.Score + int(me.getOutcome(them))

		var knownOutcome Outcome
		switch split[1] {
		case "X":
			knownOutcome = lose
		case "Y":
			knownOutcome = draw
		case "Z":
			knownOutcome = win
		}
		myMoveKnown := them.getMoveNeededForOutcome(knownOutcome.reverse())
		myScoreWithKnownOutcome += myMoveKnown.Score + int(knownOutcome)
	}

	fmt.Printf("My score is %d\n", myScore)
	fmt.Printf("My known score is %d", myScoreWithKnownOutcome)
}
