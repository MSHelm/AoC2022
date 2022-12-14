package main

import (
	"AoC2022/utils"
	"fmt"
)

func main() {
	input := utils.ReadLines()
	part1_score := part1(input)
	fmt.Println("Part1: ", part1_score)
	part2_score := part2(input)
	fmt.Println("Part2: ", part2_score)
}

func part1(input []string) int {
	score := 0
	priorities := generatePriorities()

	for _, rucksack := range input {
		compartment1, compartment2 := splitCompartments(rucksack)
		commonItem := getIntersect(compartment1, compartment2)
		score += priorities[commonItem]
	}

	return score
}

func part2(input []string) int {
	score := 0
	priorities := generatePriorities()
	for i := 0; i < len(input); i += 3 {
		commonItem := getIntersect2(input[i], input[i+1], input[i+2])
		score += priorities[commonItem]
	}
	return score
}

func splitCompartments(s string) (string, string) {
	compartment1 := s[0:(len(s) / 2)]
	compartment2 := s[len(s)/2:]
	return compartment1, compartment2
}

// Makes use of the fact that we know that there is exactly one element that intersects, which allows us to directly return
func getIntersect(compartment1 string, compartment2 string) rune {
	for _, rune1 := range compartment1 {
		for _, rune2 := range compartment2 {
			if rune1 == rune2 {
				return rune1
			}
		}
	}

	return rune(0)
}

// Makes use of ASCII int encoding of runes in Golang
func generatePriorities() map[rune]int {
	priorities := map[rune]int{}

	for i := 0; i < 26; i++ {
		priorities[rune(65+i)] = i + 27 //uppercase characters
		priorities[rune(97+i)] = i + 1  //lowercase characters
	}
	return priorities
}

// Too lazy to make it a generalized function XD
func getIntersect2(elf1 string, elf2 string, elf3 string) rune {
	for _, rune1 := range elf1 {
		for _, rune2 := range elf2 {
			if rune1 == rune2 {
				for _, rune3 := range elf3 {
					if rune1 == rune3 {
						return rune1
					}
				}
			}
		}
	}

	return rune(0)
}
