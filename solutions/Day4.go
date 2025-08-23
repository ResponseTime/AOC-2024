package solutions

const (
	DOWN = iota
	RIGHT
	LEFT
	UP
	DIAGDOWNPOS
	DIAGUPPOS
	DIAGDOWNNEG
	DIAGUPNEG
)

func buildGraph(lines []string) [][]string {
	var matrix = make([][]string, len(lines))
	for i, line := range lines {
		var row []string
		for _, r := range line {
			row = append(row, string(r))
		}
		matrix[i] = row
	}
	return matrix
}

var needed = "XMAS"

func DFS(graph [][]string, i, j int, Direction int, visited []int, needInd int) bool {
	if needInd >= len(needed) {
		return true
	}
	if i < 0 || i >= len(graph) || j < 0 || j >= len(graph[0]) {
		return false
	}
	if graph[i][j] != string(needed[needInd]) {
		return false
	}

	switch Direction {
	case DOWN:
		return DFS(graph, i+1, j, DOWN, visited, needInd+1)
	case RIGHT:
		return DFS(graph, i, j+1, RIGHT, visited, needInd+1)
	case LEFT:
		return DFS(graph, i, j-1, LEFT, visited, needInd+1)
	case UP:
		return DFS(graph, i-1, j, UP, visited, needInd+1)
	case DIAGDOWNPOS:
		return DFS(graph, i+1, j+1, DIAGDOWNPOS, visited, needInd+1)
	case DIAGUPPOS:
		return DFS(graph, i-1, j+1, DIAGUPPOS, visited, needInd+1)
	case DIAGDOWNNEG:
		return DFS(graph, i+1, j-1, DIAGDOWNNEG, visited, needInd+1)
	case DIAGUPNEG:
		return DFS(graph, i-1, j-1, DIAGUPNEG, visited, needInd+1)
	}
	return true
}

func DayFourPartOne(lines []string) int {
	var res int
	graph := buildGraph(lines)
	for i := 0; i < len(graph); i++ {
		for j := 0; j < len(graph[0]); j++ {
			if graph[i][j] == "X" {
				for _, k := range []int{RIGHT, DOWN, LEFT, UP, DIAGDOWNPOS, DIAGUPPOS, DIAGDOWNNEG, DIAGUPNEG} {
					if DFS(graph, i, j, k, []int{}, 0) {
						res++
					}
				}
			}
		}
	}
	return res
}

func DayFourPartTwo(lines []string) int {
	var res int
	graph := buildGraph(lines)
	for i := 1; i < len(graph)-1; i++ {
		for j := 1; j < len(graph[0])-1; j++ {
			if graph[i][j] == "A" {
				diag1 := graph[i-1][j-1] + graph[i][j] + graph[i+1][j+1]
				diag2 := graph[i+1][j-1] + graph[i][j] + graph[i-1][j+1]
				if (diag1 == "MAS" || diag1 == "SAM") &&
					(diag2 == "MAS" || diag2 == "SAM") {
					res++
				}
			}
		}
	}
	return res
}
