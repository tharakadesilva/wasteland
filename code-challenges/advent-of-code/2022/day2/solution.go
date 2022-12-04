package day2

import (
	"io/ioutil"
	"strings"
)

func readAllLines(path string) ([]string, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return strings.Split(string(b), "\n"), nil
}

type Move int

func (me Move) String() string {
	return [...]string{"Rock", "Paper", "String"}[me-1]
}

const (
	Rock Move = iota + 1
	Paper
	Scissors
)

func (myMove Move) netPoints(enemyMove Move) int {
	if myMove == Rock && enemyMove == Scissors ||
		myMove == Paper && enemyMove == Rock ||
		myMove == Scissors && enemyMove == Paper {
		return 6
	}

	if myMove == enemyMove {
		return 3
	}

	return 0
}

func (move Move) resolveWinsAgainst() Move {
	if move == Rock {
		return Scissors
	}

	if move == Paper {
		return Rock
	}

	return Paper
}

func (move Move) resolveLosesTo() Move {
	if move == Rock {
		return Paper
	}

	if move == Paper {
		return Scissors
	}

	return Rock
}

func resolveMove(move string) Move {
	switch move {
	case "A", "X":
		return Rock
	case "B", "Y":
		return Paper
	case "C", "Z":
		return Scissors
	}

	panic("invalid move")
}

func calculatePointsForRound(playerMove, enemyMove Move) int {
	return playerMove.netPoints(enemyMove) + int(playerMove)
}

func part1(input string) (int, error) {
	lines, err := readAllLines(input)
	if err != nil {
		return 0, err
	}

	totalPoints := 0
	for _, line := range lines {
		moves := strings.Split(line, " ")
		enemyMove := resolveMove(moves[0])
		playerMove := resolveMove(moves[1])
		totalPoints += calculatePointsForRound(playerMove, enemyMove)
	}

	return totalPoints, nil
}

func resolvePlayerMove(enemyMove Move, plan string) Move {
	if plan == "Y" {
		return enemyMove
	}

	if plan == "X" { // player should lose
		return enemyMove.resolveWinsAgainst()
	}

	if plan == "Z" { // player should win
		return enemyMove.resolveLosesTo()
	}

	panic("invalid move")
}

func part2(input string) (int, error) {
	lines, err := readAllLines(input)
	if err != nil {
		return 0, err
	}

	totalPoints := 0
	for _, line := range lines {
		moves := strings.Split(line, " ")
		enemyMove := resolveMove(moves[0])
		playerMove := resolvePlayerMove(enemyMove, moves[1])
		totalPoints += calculatePointsForRound(playerMove, enemyMove)
	}

	return totalPoints, nil
}
