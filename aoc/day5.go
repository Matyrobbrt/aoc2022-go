package aoc

import (
	"bufio"
	"os"
	"regexp"
	"strings"
)

type Step struct {
	amount      int
	source      int
	destination int
}

type Stack []rune
type Scheme []Stack

func readScheme(lines *[]string) *Scheme {
	numbers := (*lines)[len(*lines)-1]     // Last line contains the numbers of the stacks
	split := strings.Split(numbers, "   ") // We make the assumption that we encounter max 9 stacks (so 1 digit numbers)
	length := ToInt(strings.Trim(split[len(split)-1], " "))

	scheme := make([]Stack, length)

	// Fill with empty arrays
	amount := len(*lines) - 1
	for i := 0; i < length; i++ {
		scheme[i] = make([]rune, 0)
	}

	// Now to fill the stacks
	for li := amount - 1; li >= 0; li-- {
		line := []rune((*lines)[li])
		linelen := len(line)

		for i := 0; i < length; i++ {
			// The crate will find itself encased in [ ] and we also have a space between crates
			// So the second crate is located at 5 chars from the start
			idx := (4 * i) + 1
			if idx < linelen { // If the index is bigger than chars we have, it means that that row of creates finished early and we can break
				crate := line[idx]
				if crate != 32 { // The empty char means there's no crate
					scheme[i] = append(scheme[i], crate)
				}
			} else {
				break
			}
		}
	}

	return (*Scheme)(&scheme)
}

func getSteps(lines *[]string) *[]Step {
	steps := make([]Step, len(*lines))

	regex := regexp.MustCompile("move (?P<tomove>\\d+) from (?P<from>\\d+) to (?P<to>\\d+)")
	for i := 0; i < len(*lines); i++ {
		match := regex.FindStringSubmatch((*lines)[i])
		steps[i] = Step{
			amount:      ToInt(match[1]),
			source:      ToInt(match[2]),
			destination: ToInt(match[3]),
		}
	}

	return &steps
}

func Day5() {
	file, _ := os.Open("input/aocday5.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	schemeDone := false
	schemeLines := make([]string, 0)
	moveLines := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" { // Empty line means scheme9000 is done
			schemeDone = true
		} else {
			if schemeDone {
				moveLines = append(moveLines, line)
			} else {
				schemeLines = append(schemeLines, line)
			}
		}
	}

	moves := *getSteps(&moveLines)

	scheme9000 := *readScheme(&schemeLines)
	scheme9001 := *readScheme(&schemeLines)

	for _, move := range moves {
		for i := 0; i < move.amount; i++ {
			from := &scheme9000[move.source-1]
			toMove := (*from)[len(*from)-1]
			scheme9000[move.source-1] = (*from)[:len(*from)-1]

			scheme9000[move.destination-1] = append(scheme9000[move.destination-1], toMove)
		}
	}

	for _, move := range moves {
		from := &scheme9001[move.source-1]
		toMove := (*from)[len(*from)-move.amount : len(*from)]
		scheme9001[move.source-1] = (*from)[:len(*from)-move.amount]

		scheme9001[move.destination-1] = append(scheme9001[move.destination-1], toMove...)
	}

	var lastCrates string
	for _, stack := range scheme9000 {
		lastCrates += string(stack[len(stack)-1])
	}
	println("Crate sum 9000: " + lastCrates)

	lastCrates = ""
	for _, stack := range scheme9001 {
		lastCrates += string(stack[len(stack)-1])
	}
	println("Crate sum 9001: " + lastCrates)
}
