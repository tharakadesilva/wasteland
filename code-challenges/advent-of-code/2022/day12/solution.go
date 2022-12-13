package day12

import (
	"io/ioutil"
	"strings"
)

func readFile(path string) (string, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func buildGrid(input string) [][]rune {
	grid := [][]rune{}

	for _, line := range strings.Split(input, "\n") {
		grid = append(grid, []rune(line))
	}

	return grid
}

func getStartingCoords(grid [][]rune) (int, int) {
	for i, row := range grid {
		for j, r := range row {
			if r == 'E' {
				return i, j
			}
		}
	}

	return -1, -1
}

func isMoveValid(grid [][]rune, prevX, prevY, nextX, nextY int) bool {
	previous := grid[prevX][prevY]
	if previous == 'E' {
		previous = 'z'
	}

	next := grid[nextX][nextY]
	if next == 'S' {
		next = 'a'
	}

	return next+1 >= previous
}

func possibleMoves(grid [][]rune, x, y int) [][]int {
	possibleMoves := [][]int{}

	if x > 0 && isMoveValid(grid, x, y, x-1, y) {
		possibleMoves = append(possibleMoves, []int{x - 1, y})
	}
	if x < len(grid)-1 && isMoveValid(grid, x, y, x+1, y) {
		possibleMoves = append(possibleMoves, []int{x + 1, y})
	}
	if y > 0 && isMoveValid(grid, x, y, x, y-1) {
		possibleMoves = append(possibleMoves, []int{x, y - 1})
	}
	if y < len(grid[0])-1 && isMoveValid(grid, x, y, x, y+1) {
		possibleMoves = append(possibleMoves, []int{x, y + 1})
	}

	return possibleMoves
}

func djiksras(grid [][]rune, x, y int, end rune) int {
	visited := make(map[int]map[int]bool)
	for i := 0; i < len(grid); i++ {
		visited[i] = make(map[int]bool)
	}
	visited[x][y] = true

	queue := [][]int{{x, y, 0}}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if grid[current[0]][current[1]] == end {
			return current[2]
		}

		for _, move := range possibleMoves(grid, current[0], current[1]) {
			if visited[move[0]][move[1]] {
				continue
			}

			visited[move[0]][move[1]] = true
			queue = append(queue, []int{move[0], move[1], current[2] + 1})
		}
	}

	return -1
}

func solve(input string, end rune) (int, error) {
	contents, err := readFile(input)
	if err != nil {
		return -1, err
	}

	grid := buildGrid(contents)
	x, y := getStartingCoords(grid)
	return djiksras(grid, x, y, end), nil
}

func part1(input string) (int, error) {
	return solve(input, 'S')
}

func part2(input string) (int, error) {
	return solve(input, 'a')
}
