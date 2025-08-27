package solutions

import (
	"fmt"
	"slices"
)

type traveller struct {
	x         int
	y         int
	Direction int
}
type point struct {
	x int
	y int
}

// var obstaclePath []point
// var objsPlaced int

func goPath(graph [][]string, Direction, i, j, distinct int, seen map[traveller]bool) (int, bool) {
	if i < 0 || i >= len(graph) || j < 0 || j >= len(graph[0]) {
		return distinct, false
	}
	p := traveller{x: i, y: j, Direction: Direction}
	// for _, pointCheck := range obstaclePath {
	// 	if p.x == pointCheck.x || p.y == pointCheck.y {
	// 		objsPlaced++
	// 		break
	// 	}
	// }
	val, ok := seen[p]
	if val {
		return 0, true
	}
	if !ok {
		// graph[i][j] = "X"
		distinct++
	}
	seen[p] = true
	switch Direction {
	case UP:
		if i-1 < 0 {
			return distinct, false
		}
		if graph[i-1][j] == "#" {
			return goPath(graph, RIGHT, i, j, distinct, seen)
		}
		return goPath(graph, UP, i-1, j, distinct+1, seen)

	case RIGHT:
		if j+1 >= len(graph[0]) {
			return distinct, false
		}
		if graph[i][j+1] == "#" {
			return goPath(graph, DOWN, i, j, distinct, seen)
		}
		return goPath(graph, RIGHT, i, j+1, distinct+1, seen)

	case DOWN:
		if i+1 >= len(graph) {
			return distinct, false
		}
		if graph[i+1][j] == "#" {
			return goPath(graph, LEFT, i, j, distinct, seen)
		}
		return goPath(graph, DOWN, i+1, j, distinct+1, seen)

	case LEFT:
		if j-1 < 0 {
			return distinct, false
		}
		if graph[i][j-1] == "#" {
			return goPath(graph, UP, i, j, distinct, seen)
		}
		return goPath(graph, LEFT, i, j-1, distinct+1, seen)
	}
	return 0, false
}

func DaySixPartOne(lines []string) int {
	var res int
	var startPoint traveller
	graph := buildGraph(lines)
	for i := 0; i < len(graph); i++ {
		if y := slices.Index(graph[i], "^"); y > -1 {
			startPoint.x = i
			startPoint.y = y
			startPoint.Direction = UP
		}
	}
	fmt.Println(startPoint)
	res, _ = goPath(graph, startPoint.Direction, startPoint.x, startPoint.y, 0, map[traveller]bool{})
	for _, row := range graph {
		fmt.Println(row)
	}
	return res
}

func DaySixPartTwo(lines []string) int {
	var res int
	var startPoint traveller
	graph := buildGraph(lines)
	for i := 0; i < len(graph); i++ {
		if y := slices.Index(graph[i], "^"); y > -1 {
			startPoint.x = i
			startPoint.y = y
			startPoint.Direction = UP
		}
	}
	for i := 0; i < len(graph); i++ {
		for j := 0; j < len(graph[0]); j++ {
			if graph[i][j] == "." {
				graph[i][j] = "#"
				if _, cycle := goPath(graph, startPoint.Direction, startPoint.x, startPoint.y, 0, map[traveller]bool{}); cycle {
					res++
				}
				graph[i][j] = "."
			}
		}
	}
	fmt.Println(startPoint)
	// goPath(graph, startPoint.Direction, startPoint.x, startPoint.y, 0, map[point]bool{})
	// res = objsPlaced
	// for _, row := range graph {
	// 	fmt.Println(row)
	// }
	return res
}
