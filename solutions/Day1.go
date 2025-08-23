package solutions

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

func DayOnePartOne(lines []string) int {
	var left []int
	var right []int
	var res int
	for _, line := range lines {
		ren := strings.Fields(line)
		left_num, err := strconv.Atoi(string(ren[0]))
		right_num, err := strconv.Atoi(string(ren[1]))
		if err != nil {
			fmt.Println(err)
		}
		left = append(left, left_num)
		right = append(right, right_num)
	}
	slices.Sort(left)
	slices.Sort(right)

	for i := range len(left) {
		res += int(math.Abs(float64(right[i] - left[i])))
	}
	return res
}

func DayOnePartTwo(lines []string) int {
	var left []int
	var right []int
	counter := make(map[int]int)
	var res int
	for _, line := range lines {
		ren := strings.Fields(line)
		left_num, err := strconv.Atoi(string(ren[0]))
		right_num, err := strconv.Atoi(string(ren[1]))
		if err != nil {
			fmt.Println(err)
		}
		left = append(left, left_num)
		right = append(right, right_num)
	}
	slices.Sort(left)
	slices.Sort(right)
	for i := range right {
		counter[right[i]] = 0
	}
	for i := range right {
		counter[right[i]] += 1
	}
	for i := range left {
		res += left[i] * counter[left[i]]
	}
	return res
}
