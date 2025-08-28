package solutions

import (
	"strconv"
	"strings"
)

const (
	ADD = iota
	MULTIPLY
)

func combs(target, res int, args []int) bool {
	if len(args) == 0 {
		return res == target
	}
	if res > target {
		return false
	}

	if combs(target, res+args[0], args[1:]) {
		return true
	}
	mul, q := 10, 10
	for q != 0 {
		q = args[0] / mul
		if q > 0 {
			mul *= 10
		}
	}
	calculation := (res * mul) + args[0]
	if combs(target, calculation, args[1:]) {
		return true
	}
	if combs(target, res*args[0], args[1:]) {
		return true
	}
	return false
}

func DaySevenPartOne(lines []string) int {
	var res int
	for _, line := range lines {
		num_left := strings.Split(line, ":")[0]
		numleft, _ := strconv.Atoi(num_left)
		num_args := strings.TrimSpace(strings.Split(line, ":")[1])
		fields := strings.Fields(num_args)
		var nums []int
		for _, i := range fields {
			conv, _ := strconv.Atoi(i)
			nums = append(nums, conv)
		}
		if combs(numleft, nums[0], nums[1:]) {
			res += numleft
		}
	}
	return res
}

func DaySevenPartTwo(lines []string) int {
	var res int
	for _, line := range lines {
		num_left := strings.Split(line, ":")[0]
		numleft, _ := strconv.Atoi(num_left)
		num_args := strings.TrimSpace(strings.Split(line, ":")[1])
		fields := strings.Fields(num_args)
		var nums []int
		for _, i := range fields {
			conv, _ := strconv.Atoi(i)
			nums = append(nums, conv)
		}
		if combs(numleft, nums[0], nums[1:]) {
			res += numleft
		}
		// for i := range len(nums) - 1 {
		// 	var newArr []int
		// 	newArr = append(newArr, nums[:i]...)
		// 	newArr = append(newArr, nums[i]*10+nums[i+1])
		// 	newArr = append(newArr, nums[i+2:]...)
		// 	fmt.Println(newArr)
		// 	if combs(numleft, newArr[0], newArr[0], 1, newArr) {
		// 		// fmt.Println(line)
		// 		res += numleft
		// 	}
		// }
	}
	return res
}
