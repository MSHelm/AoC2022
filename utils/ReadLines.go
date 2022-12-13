package utils

import (
	"bufio"
	"os"
)

func ReadLines() []string {
	lines := []string{}

	f, err := os.Open("input.txt")
	check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}
