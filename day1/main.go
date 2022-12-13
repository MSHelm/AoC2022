package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	f, err := os.ReadFile("input.txt")
	check(err)

	s := string(f)
	s = strings.TrimRight(s, "\n")
	elves := strings.Split(s, "\n\n")

	total_calories := calcTotalCalories(elves)
	sort.Sort(sort.Reverse(sort.IntSlice(total_calories)))

	fmt.Println("Part1: ", total_calories[0])
	fmt.Println("Part2: ", calcSumIntSlice(total_calories[0:3]))
}

func calcTotalCalories(elves []string) []int {
	var calories []int
	for _, elf := range elves {
		elf_string := strings.Split(elf, "\n")
		elf_num, err := sliceAtoi(elf_string)
		check(err)
		elf_calories := calcSumIntSlice(elf_num)
		calories = append(calories, elf_calories)
	}

	return calories
}

func sliceAtoi(strings []string) ([]int, error) {
	var ints []int
	for _, s := range strings {
		i, err := strconv.Atoi(s)
		check(err)
		ints = append(ints, i)
	}

	return ints, nil
}

func calcSumIntSlice(nums []int) int {
	var sum int
	for _, n := range nums {
		sum += n
	}

	return sum
}
