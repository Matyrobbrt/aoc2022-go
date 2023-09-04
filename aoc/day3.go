package aoc

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

var priorities = func() map[rune]int {
	priorities := make(map[rune]int, 52)
	prio := 0

	for i := 97; i <= 122; i++ {
		prio += 1
		priorities[rune(i)] = prio
	}

	for i := 65; i <= 90; i++ {
		prio += 1
		priorities[rune(i)] = prio
	}
	return priorities
}()

func Day3() {
	file, _ := os.Open("input/aocday3.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	groupsOf3 := NewCollector[string](3)

	var everyPrioSum int
	for scanner.Scan() {
		line := scanner.Text()
		length := len(line)
		groupsOf3.Accept(line)

		first := []rune(line[0 : length/2])
		second := []rune(line[length/2:])

		var inBoth rune
		for _, char := range first {
			if slices.Contains(second, char) {
				inBoth = char
			}
		}
		everyPrioSum += priorities[inBoth]
	}

	var badgePrioSum int
	for _, group := range groupsOf3.Values {
		chars1 := []rune(group[0])
		chars2 := []rune(group[1])
		chars3 := []rune(group[2])

		var inAll rune
		for _, char := range chars1 {
			if slices.Contains(chars2, char) && slices.Contains(chars3, char) {
				inAll = char
			}
		}

		badgePrioSum += priorities[inAll]
	}

	fmt.Printf("The sum of the priorities of items found in each compartment is %d\n", everyPrioSum)
	fmt.Printf("The sum of the priorities of badges is %d", badgePrioSum)
}
