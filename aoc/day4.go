package aoc

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Elf struct {
	first int
	last  int
}

func GetElfAssignments(elf string) Elf {
	spl := strings.Split(elf, "-")
	return Elf{ToInt(spl[0]), ToInt(spl[1])}
}

func Day4() {
	file, _ := os.Open("input/aocday4.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var fullyContainingTheOther, overlapping int
	for scanner.Scan() {
		line := scanner.Text()
		elves := strings.Split(line, ",")

		elf1 := GetElfAssignments(elves[0])
		elf2 := GetElfAssignments(elves[1])

		elfPair := [][]Elf{
			{elf1, elf2},
			{elf2, elf1},
		}

		for _, v := range elfPair {
			if v[0].first >= v[1].first && v[0].last <= v[1].last {
				fullyContainingTheOther++
				break
			}
		}

		for _, v := range elfPair {
			if v[0].first >= v[1].first && v[0].first <= v[1].last {
				overlapping++
				break
			}
		}
	}

	fmt.Printf("Number of assignments fully containing the other: %d\n", fullyContainingTheOther)
	fmt.Printf("Number of assignments overlapping: %d", overlapping)
}
