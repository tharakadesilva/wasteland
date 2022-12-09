package day8

import (
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func readAllLines(path string) ([]string, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return strings.Split(string(b), "\n"), nil
}

func runeToInt(r rune) (int, error) {
	return strconv.Atoi(string(r))
}

func readDataTo2DArray(lines []string) ([][]int, error) {
	data := make([][]int, len(lines))
	for i, line := range lines {
		for _, num := range line {
			n, err := runeToInt(num)
			if err != nil {
				return nil, err
			}

			data[i] = append(data[i], n)
		}
	}

	return data, nil
}

func initializeVisibilityGrid(data [][]int) [][]bool {
	visibilityGrid := make([][]bool, len(data))
	for i := range visibilityGrid {
		visibilityGrid[i] = make([]bool, len(data[i]))
	}
	return visibilityGrid
}

func markEdgesOfGridAsVisible(data [][]int, visibilityGrid [][]bool) {
	for i := range data {
		visibilityGrid[i][0] = true
		visibilityGrid[i][len(data[i])-1] = true
	}

	for i := range data[0] {
		visibilityGrid[0][i] = true
		visibilityGrid[len(data)-1][i] = true
	}
}

func processLeftToRight(data [][]int, visibilityGrid [][]bool) {
	for i := 1; i < len(data)-1; i++ {
		row := data[i]
		currentMax := row[0]
		for j := 1; j < len(row)-1; j++ {
			if data[i][j] > currentMax {
				currentMax = data[i][j]
				visibilityGrid[i][j] = true
			}
		}
	}
}

func processRightToLeft(data [][]int, visibilityGrid [][]bool) {
	for i := 1; i < len(data)-1; i++ {
		row := data[i]
		currentMax := row[len(row)-1]
		for j := len(row) - 2; j > 0; j-- {
			if data[i][j] > currentMax {
				currentMax = data[i][j]
				visibilityGrid[i][j] = true
			}
		}
	}
}

func processTopToBottom(data [][]int, visibilityGrid [][]bool) {
	for j := 1; j < len(data[0])-1; j++ {
		currentMax := data[0][j]
		for i := 1; i < len(data)-1; i++ {
			if data[i][j] > currentMax {
				currentMax = data[i][j]
				visibilityGrid[i][j] = true
			}
		}
	}
}

func processBottomToTop(data [][]int, visibilityGrid [][]bool) {
	for j := 1; j < len(data[0])-1; j++ {
		currentMax := data[len(data)-1][j]
		for i := len(data) - 2; i > 0; i-- {
			if data[i][j] > currentMax {
				currentMax = data[i][j]
				visibilityGrid[i][j] = true
			}
		}
	}
}

func countVisibleNodes(visibilityGrid [][]bool) int {
	count := 0
	for _, row := range visibilityGrid {
		for _, visible := range row {
			if visible {
				count++
			}
		}
	}
	return count
}

func part1(input string) (int, error) {
	lines, err := readAllLines(input)
	if err != nil {
		return 0, err
	}

	data, err := readDataTo2DArray(lines)
	if err != nil {
		return 0, err
	}

	visibilityGrid := initializeVisibilityGrid(data)
	markEdgesOfGridAsVisible(data, visibilityGrid)
	processLeftToRight(data, visibilityGrid)
	processRightToLeft(data, visibilityGrid)
	processTopToBottom(data, visibilityGrid)
	processBottomToTop(data, visibilityGrid)

	return countVisibleNodes(visibilityGrid), nil
}

func countTreesVisibleOnLeft(data [][]int, row, col int) int {
	count := 0

	value := data[row][col]
	for j := col - 1; j >= 0; j-- {
		count++

		if data[row][j] >= value {
			break
		}
	}

	return count
}

func countTreesVisibleOnRight(data [][]int, row, col int) int {
	count := 0

	value := data[row][col]
	for j := col + 1; j < len(data[row]); j++ {
		count++

		if data[row][j] >= value {
			break
		}
	}

	return count
}

func countTreesVisibleOnTop(data [][]int, row, col int) int {
	count := 0

	value := data[row][col]
	for i := row - 1; i >= 0; i-- {
		count++

		if data[i][col] >= value {
			break
		}
	}

	return count
}

func countTreesVisibleOnBottom(data [][]int, row, col int) int {
	count := 0

	value := data[row][col]
	for i := row + 1; i < len(data); i++ {
		count++

		if data[i][col] >= value {
			break
		}
	}

	return count
}

func calculateScenicScore(data [][]int, row, col int) int {
	leftCount := countTreesVisibleOnLeft(data, row, col)
	rightCount := countTreesVisibleOnRight(data, row, col)
	topCount := countTreesVisibleOnTop(data, row, col)
	bottomCount := countTreesVisibleOnBottom(data, row, col)

	return leftCount * rightCount * topCount * bottomCount
}

func calculateMaxScenicScore(data [][]int, visibilityGrid [][]bool) int {
	maxScore := math.MinInt

	for i, row := range data {
		for j := range row {
			if visibilityGrid[i][j] {
				score := calculateScenicScore(data, i, j)
				if score > maxScore {
					maxScore = score
				}
			}
		}
	}

	return maxScore
}

func part2(input string) (int, error) {
	lines, err := readAllLines(input)
	if err != nil {
		return 0, err
	}

	data, err := readDataTo2DArray(lines)
	if err != nil {
		return 0, err
	}

	visibilityGrid := initializeVisibilityGrid(data)
	markEdgesOfGridAsVisible(data, visibilityGrid)
	processLeftToRight(data, visibilityGrid)
	processRightToLeft(data, visibilityGrid)
	processTopToBottom(data, visibilityGrid)
	processBottomToTop(data, visibilityGrid)

	return calculateMaxScenicScore(data, visibilityGrid), nil
}
