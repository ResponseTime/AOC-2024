package solutions

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func DayTwoPartOne(lines []string) int {
	var safe int
	var res int
	for _, line := range lines {
		ren := strings.Fields(line)
		asc := slices.IsSortedFunc(ren, func(a, b string) int {
			aN, _ := strconv.Atoi(a)
			bN, _ := strconv.Atoi(b)
			return aN - bN
		})

		desc := slices.IsSortedFunc(ren, func(a, b string) int {
			aN, _ := strconv.Atoi(a)
			bN, _ := strconv.Atoi(b)
			return bN - aN
		})
		safe = 0
		prev := 0
		for _, num := range ren {

			n, err := strconv.Atoi(num)
			if err != nil {
				fmt.Println(err)
			}
			if prev == 0 {
				prev = n
				continue
			}
			if asc && prev < n && n-prev <= 3 {
				safe++
			} else if desc && n < prev && prev-n <= 3 {
				safe++
			}
			prev = n
		}
		if safe == len(ren)-1 {
			res++
		}
	}
	return res

}

// func DayTwoPartTwo(lines []string) int {
// 	var safe int
// 	var res int
// 	for _, line := range lines {
// 		ren := strings.Fields(line)
// 		asc := slices.IsSortedFunc(ren, func(a, b string) int {
// 			aN, _ := strconv.Atoi(a)
// 			bN, _ := strconv.Atoi(b)
// 			if aN > bN {
// 				return 1
// 			}
// 			return -1
// 		})

// 		desc := slices.IsSortedFunc(ren, func(a, b string) int {
// 			aN, _ := strconv.Atoi(a)
// 			bN, _ := strconv.Atoi(b)
// 			if aN < bN {
// 				return 1
// 			}
// 			return -1
// 		})

// 		if !asc && !desc {
// 			for i := range ren {
// 				// asc = slices.IsSortedFunc(slices.Concat(ren[0:i], ren[i+1:]), func(a, b string) int {
// 				// 	aN, _ := strconv.Atoi(a)
// 				// 	bN, _ := strconv.Atoi(b)
// 				// 	if aN > bN {
// 				// 		return 1
// 				// 	}
// 				// 	return -1
// 				// })
// 				// desc = slices.IsSortedFunc(slices.Concat(ren[0:i], ren[i+1:]), func(a, b string) int {
// 				// 	aN, _ := strconv.Atoi(a)
// 				// 	bN, _ := strconv.Atoi(b)
// 				// 	if aN < bN {
// 				// 		return 1
// 				// 	}
// 				// 	return -1
// 				// })

// 				{

// 					prev := 0
// 					safe = 0
// 					for _, num := range slices.Concat(ren[0:i], ren[i+1:]) {
// 						asc = slices.IsSortedFunc(slices.Concat(ren[0:i], ren[i+1:]), func(a, b string) int {
// 							aN, _ := strconv.Atoi(a)
// 							bN, _ := strconv.Atoi(b)
// 							if aN > bN {
// 								return 1
// 							}
// 							return -1
// 						})
// 						desc = slices.IsSortedFunc(slices.Concat(ren[0:i], ren[i+1:]), func(a, b string) int {
// 							aN, _ := strconv.Atoi(a)
// 							bN, _ := strconv.Atoi(b)
// 							if aN < bN {
// 								return 1
// 							}
// 							return -1
// 						})
// 						if asc || desc {
// 							fmt.Println(slices.Concat(ren[0:i], ren[i+1:]))
// 							n, err := strconv.Atoi(num)
// 							if err != nil {
// 								fmt.Println(err)
// 							}
// 							if prev == 0 {
// 								prev = n
// 								continue
// 							}

// 							if asc && prev < n && (n-prev) <= 3 {
// 								safe++
// 							} else if desc && n < prev && (prev-n) <= 3 {
// 								safe++
// 							}
// 							prev = n
// 						} else {
// 							continue
// 						}
// 						if safe == len(slices.Concat(ren[0:i], ren[i+1:]))-1 {
// 							res++
// 							break
// 						}
// 					}

// 				}
// 			}
// 		}
// 		safe = 0
// 		prev := 0
// 		for _, num := range ren {
// 			n, err := strconv.Atoi(num)
// 			if err != nil {
// 				fmt.Println(err)
// 			}
// 			if prev == 0 {
// 				prev = n
// 				continue
// 			}

//				if asc && prev < n && (n-prev) <= 3 {
//					safe++
//				} else if desc && n < prev && (prev-n) <= 3 {
//					safe++
//				} else {
//					safe = -1
//					break
//				}
//				prev = n
//			}
//			if safe == len(ren)-1 {
//				res++
//			}
//		}
//		return res
//	}
func isValid(nums []int) bool {
	asc := slices.IsSortedFunc(nums, func(a, b int) int {
		if a > b {
			return 1
		}
		return -1
	})
	desc := slices.IsSortedFunc(nums, func(a, b int) int {
		if a < b {
			return 1
		}
		return -1
	})

	if !asc && !desc {
		return false
	}

	for i := 1; i < len(nums); i++ {
		diff := nums[i] - nums[i-1]
		if asc {
			if diff <= 0 || diff > 3 {
				return false
			}
		}
		if desc {
			diff = nums[i-1] - nums[i]
			if diff <= 0 || diff > 3 {
				return false
			}
		}
	}

	return true
}

func isSafe(fields []string) bool {
	nums := make([]int, len(fields))
	for i, f := range fields {
		nums[i], _ = strconv.Atoi(f)
	}

	if isValid(nums) {
		return true
	}

	for i := range nums {
		conc := slices.Concat(nums[:i], nums[i+1:])
		if isValid(conc) {
			return true
		}
	}

	return false
}

func DayTwoPartTwo(lines []string) int {
	res := 0
	for _, line := range lines {
		fields := strings.Fields(line)
		if isSafe(fields) {
			res++
		}
	}
	return res
}
