package solutions

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func makeMapOfReq(lines []string) map[string][]string {
	reqMap := make(map[string][]string)
	for _, line := range lines {
		splitkv := strings.Split(line, "|")
		reqMap[splitkv[1]] = append(reqMap[splitkv[1]], splitkv[0])
	}
	return reqMap
}

func pageOrderChecker(lines []string, mapOfReq map[string][]string) (map[int][]int, map[int][]int) {
	var nums []int
	safe := make(map[int][]int)
	unsafe := make(map[int][]int, 0)
	for ln, line := range lines {
		arr := strings.Split(line, ",")
		nums = make([]int, len(arr))
		for i, n := range arr {
			conv, err := strconv.Atoi(n)
			if err != nil {
				fmt.Println(err)
			}
			nums[i] = conv
		}
		var safe_row bool = true
		var i int
		for i = len(nums) - 1; i >= 0; i-- {
			neededOrder := mapOfReq[strconv.Itoa(nums[i])]
			if len(neededOrder) < 1 {
				continue
			}
			order := make([]int, len(neededOrder))
			for k, n := range neededOrder {
				conv, err := strconv.Atoi(n)
				if err != nil {
					fmt.Println(err)
				}
				order[k] = conv
			}

			for _, check := range order {
				if slices.Contains(nums, check) && !slices.Contains(nums[:i], check) {
					safe_row = false
					break
				}
			}
		}
		if safe_row {
			safe[ln] = append(safe[ln], nums...)
		} else {
			unsafe[ln] = append(unsafe[ln], nums...)
		}
	}
	return safe, unsafe
}

func countMiddles(safeRows map[int][]int) int {
	var total int
	for _, v := range safeRows {
		total += v[len(v)/2]
	}
	return total
}

func fixPageOrder(unsafeRows map[int][]int, mapOfReq map[string][]string) map[int][]int {
	mapOfCorrected := make(map[int][]int)

	// Create a dependency map for faster lookups
	precedenceMap := make(map[int][]int)
	for kStr, vStrs := range mapOfReq {
		k, err := strconv.Atoi(kStr)
		if err != nil {
			fmt.Println("Error converting key:", err)
			continue
		}
		var precedes []int
		for _, vStr := range vStrs {
			v, err := strconv.Atoi(vStr)
			if err != nil {
				fmt.Println("Error converting value:", err)
				continue
			}
			precedes = append(precedes, v)
		}
		precedenceMap[k] = precedes
	}

	for k, v := range unsafeRows {
		slices.SortFunc(v, func(a, b int) int {
			for _, p := range precedenceMap[b] {
				if p == a {
					return -1
				}
			}
			for _, p := range precedenceMap[a] {
				if p == b {
					return 1
				}
			}
			return 0
		})
		mapOfCorrected[k] = v
	}

	return mapOfCorrected
}

func DayFivePartOne(lines []string) int {
	var res int
	indOfBreak := slices.Index(lines, "")
	reqMap := makeMapOfReq(lines[:indOfBreak])
	fmt.Println(reqMap)
	mapOfSafeRows, _ := pageOrderChecker(lines[indOfBreak+1:], reqMap)
	res = countMiddles(mapOfSafeRows)
	return res
}

func DayFivePartTwo(lines []string) int {
	var res int
	indOfBreak := slices.Index(lines, "")
	reqMap := makeMapOfReq(lines[:indOfBreak])
	fmt.Println(reqMap)
	_, mapOfUnsafeRows := pageOrderChecker(lines[indOfBreak+1:], reqMap)
	correctedPageOrder := fixPageOrder(mapOfUnsafeRows, reqMap)
	res = countMiddles(correctedPageOrder)
	return res
}
