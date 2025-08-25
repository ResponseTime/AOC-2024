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

func goPath(graph [][]string, Direction, i, j, distinct int, seen map[point]bool) int {
	if i < 0 || i >= len(graph) || j < 0 || j >= len(graph[0]) {
		return distinct
	}
	p := point{x: i, y: j}
	// for _, pointCheck := range obstaclePath {
	// 	if p.x == pointCheck.x || p.y == pointCheck.y {
	// 		objsPlaced++
	// 		break
	// 	}
	// }
	_, ok := seen[p]
	if !ok {
		graph[i][j] = "X"
		distinct++
	}
	seen[p] = true
	switch Direction {
	case UP:
		if i-1 < 0 {
			return distinct
		}
		if graph[i-1][j] == "#" {
			// obstaclePath = append(obstaclePath, point{x: i - 1, y: j})
			return goPath(graph, RIGHT, i, j+1, distinct, seen)
		}
		return goPath(graph, UP, i-1, j, distinct, seen)
	case DOWN:
		if i+1 >= len(graph) {
			return distinct
		}
		if graph[i+1][j] == "#" {
			// obstaclePath = append(obstaclePath, point{x: i + 1, y: j})
			return goPath(graph, LEFT, i, j-1, distinct, seen)
		}
		return goPath(graph, DOWN, i+1, j, distinct, seen)
	case RIGHT:
		if j+1 >= len(graph[0]) {
			return distinct
		}
		if graph[i][j+1] == "#" {
			// obstaclePath = append(obstaclePath, point{x: i, y: j + 1})
			return goPath(graph, DOWN, i+1, j, distinct, seen)
		}
		return goPath(graph, RIGHT, i, j+1, distinct, seen)
	case LEFT:
		if j-1 < 0 {
			return distinct
		}
		if graph[i][j-1] == "#" {
			// obstaclePath = append(obstaclePath, point{x: i, y: j - 1})
			return goPath(graph, UP, i-1, j, distinct, seen)
		}
		return goPath(graph, LEFT, i, j-1, distinct, seen)
	}
	return 0
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
	res = goPath(graph, startPoint.Direction, startPoint.x, startPoint.y, 0, map[point]bool{})
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
	fmt.Println(startPoint)
	goPath(graph, startPoint.Direction, startPoint.x, startPoint.y, 0, map[point]bool{})
	// res = objsPlaced
	for _, row := range graph {
		fmt.Println(row)
	}
	return res
}
