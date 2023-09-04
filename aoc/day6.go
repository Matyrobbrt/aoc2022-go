package aoc

import (
	"bufio"
	"fmt"
	"os"
)

func Day6() {
	file, _ := os.Open("input/aocday6.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	line := []rune(scanner.Text())

	for i := 0; i < len(line); i++ {
		last4 := line[i : i+4]
		if AllValuesUnique(&last4) {
			fmt.Printf("%d chars need to be processed before the first start-of-packet marker", i+4)
			break
		}
	}

	for i := 0; i < len(line); i++ {
		last14 := line[i : i+14]
		if AllValuesUnique(&last14) {
			fmt.Printf("%d chars need to be processed before the first start-of-message marker\n", i+14)
			break
		}
	}
}
