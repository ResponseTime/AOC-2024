package solutions

import (
	"strconv"
	"strings"
)

func checkValidity(line string) int {
	// var total int = 0
	var checkerStack = []rune{'(', 'n', ',', 'n', ')'}
	var multiplyStack []int
	var i = 0
	var last int = 0
	var v int = 0
	for {
		char := rune(line[v])
		if char == ')' && checkerStack[i] == char {
			break
		} else if checkerStack[i] == 'n' {
			num := ""
			k := v
			for {
				_, err := strconv.Atoi(string(line[k]))
				if err != nil {
					break
				}
				num += string(line[k])
				k += 1
			}
			conv, err := strconv.Atoi(num)
			if err != nil {
				return 0
			}
			multiplyStack = append(multiplyStack, conv)
			last = k
			v = last - 1
		} else if checkerStack[i] != char {
			return 0
		}
		i++
		v++
	}
	return multiplyStack[0] * multiplyStack[1]
}
func DayThreePartOneAndTwo(lines []string) int {
	res := 0
	var do bool = true
	for _, line := range lines {
		for i := range line {
			if strings.HasPrefix(line[i:], "do()") {
				do = true
			}
			if strings.HasPrefix(line[i:], "don't()") {
				do = false
			}
			if strings.HasPrefix(line[i:], "mul") {
				if do {
					res += checkValidity(line[i+3:])
				}
			}
		}
	}
	return res
}
