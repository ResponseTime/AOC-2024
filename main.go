package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/responsetime/aoc/solutions"
)

func main() {
	var lines []string
	reader := bufio.NewScanner(os.Stdin)
	for reader.Scan() {
		lines = append(lines, reader.Text())
	}
	// Day 1 P1 P2
	// sol := solutions.DayOnePartOne(lines)
	// sol := solutions.DayOnePartTwo(lines)
	// fmt.Println(sol)

	// Day 2 P1 P2
	// sol := solutions.DayTwoPartOne(lines)
	// sol := solutions.DayTwoPartTwo(lines)
	// fmt.Println(sol)

	// Day 3 P1 P2
	// sol := solutions.DayThreePartOneAndTwo(lines)
	// fmt.Println(sol)

	//Day 4 P1 P2
	// sol := solutions.DayFourPartOne(lines)
	// sol := solutions.DayFourPartTwo(lines)
	// fmt.Println(sol)

	//Day 5 P1 P2
	// sol := solutions.DayFivePartOne(lines)
	// sol := solutions.DayFivePartTwo(lines)
	// fmt.Println(sol)

	//Day 6 P1 P2
	// sol := solutions.DaySixPartOne(lines)
	sol := solutions.DaySixPartTwo(lines)
	fmt.Println(sol)
}
