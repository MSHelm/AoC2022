package main

import (
	"fmt"
	"strings"

	"AoC2022/utils"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	rounds := utils.ReadLines()
	part1_score := part1(rounds)
	fmt.Println("Part1 final score: ", part1_score)
	part2_score := part2(rounds)
	fmt.Println("Part2 final score: ", part2_score)
}

func part1(rounds []string) int {
	// var handEncodings = map[string]string{
	// 	"A": "rock",
	// 	"B": "paper",
	// 	"C": "scissors",

	// 	"X": "rock",
	// 	"Y": "paper",
	// 	"Z": "scissors",
	// }

	var roundOutcomes = map[string]string{
		"A X": "tie",
		"A Y": "win",
		"A Z": "loose",

		"B X": "loose",
		"B Y": "tie",
		"B Z": "win",

		"C X": "win",
		"C Y": "loose",
		"C Z": "tie",
	}

	var points = map[string]int{
		"loose": 0,
		"tie":   3,
		"win":   6,
		"X":     1, //rock
		"Y":     2, // paper
		"Z":     3, // scissors
	}

	score := 0
	for _, round := range rounds {
		roundOutcome := roundOutcomes[round]
		roundPoint := points[roundOutcome]
		score += roundPoint

		myHand := getOwnHand(round)
		handPoint := points[myHand]
		score += handPoint
	}
	return score
}

func part2(rounds []string) int {
	var playbook = map[string]map[string]string{
		"A": {
			"X": "scissors",
			"Y": "rock",
			"Z": "paper",
		},
		"B": {
			"X": "rock",
			"Y": "paper",
			"Z": "scissors",
		},
		"C": {
			"X": "paper",
			"Y": "scissors",
			"Z": "rock",
		},
	}

	var points = map[string]int{
		"rock":     1,
		"paper":    2,
		"scissors": 3,
		"X":        0,
		"Y":        3,
		"Z":        6,
	}

	score := 0
	for _, round := range rounds {
		opponentHand := getOpponentHand(round)
		desiredOutcome := getDesiredOutcome(round)
		handToPlay := playbook[opponentHand][desiredOutcome]
		score += points[handToPlay]
		score += points[desiredOutcome]
	}

	return score
}

func getOwnHand(s string) string {
	self := strings.Split(s, " ")[1]
	return self
}

func getDesiredOutcome(s string) string {
	return getOwnHand(s)
}

func getOpponentHand(s string) string {
	opponent := strings.Split(s, " ")[0]
	return opponent
}
