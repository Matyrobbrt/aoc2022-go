package aoc

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func Day1() {
	file, _ := os.Open("input/aocday1.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	currentElf := 0
	elves := make([]int, 20)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			elves = append(elves, currentElf)
			currentElf = 0
		} else {
			currentElf += ToInt(line)
		}
	}

	elves = append(elves, currentElf)

	sort.Slice(elves, func(i, j int) bool {
		return elves[i] > elves[j]
	})

	fmt.Printf("The most Calories carried by an elf are %d\n", elves[0])
	fmt.Printf("The Calories carried by the 3 best elves are %d", elves[0]+elves[1]+elves[2])
}
