package main

import (
	"AoC2022/utils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := utils.ReadLines()
	part1_score := part1(input)
	fmt.Println("Score1: ", part1_score)
	part2_score := part2(input)
	fmt.Println("Score2: ", part2_score)
}

func part1(input []string) int {
	score := 0
	for _, pair := range input {
		if isFullyContained(pair) {
			score += 1
		}
	}
	return score
}

func part2(input []string) int {
	score := 0
	for _, pair := range input {
		if isOverlapping(pair) {
			score += 1
		}
	}
	return score
}

func isFullyContained(pair string) bool {
	elf1, elf2 := splitElves(pair)
	elf1_lower, elf1_upper := getBoundaries(elf1)
	elf2_lower, elf2_upper := getBoundaries(elf2)
	contained := false
	if (elf1_lower <= elf2_lower && elf1_upper >= elf2_upper) || (elf2_lower <= elf1_lower && elf2_upper >= elf1_upper) {
		contained = true
	}

	return contained
}

func splitElves(s string) (string, string) {
	split := strings.Split(s, ",")
	elf1 := split[0]
	elf2 := split[1]

	return elf1, elf2
}

func getBoundaries(s string) (int, int) {
	split := strings.Split(s, "-")
	lower, _ := strconv.Atoi(split[0])
	upper, _ := strconv.Atoi(split[1])

	return lower, upper
}

func isOverlapping(pair string) bool {
	elf1, elf2 := splitElves(pair)
	elf1_lower, elf1_upper := getBoundaries(elf1)
	elf2_lower, elf2_upper := getBoundaries(elf2)

	// sort!
	if elf1_lower > elf2_upper {
		elf1_lower, elf1_upper, elf2_lower, elf2_upper = elf2_lower, elf2_upper, elf1_lower, elf1_upper
	}

	overlapping := false

	if elf1_upper >= elf2_lower {
		overlapping = true
	}

	return overlapping
}
