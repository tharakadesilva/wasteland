package day9

import (
	"fmt"
	"io/ioutil"
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

type Knot struct {
	row             int
	col             int
	next            *Knot
	uniqueLocations map[string]struct{}
}

func NewKnot(row, col int) *Knot {
	k := &Knot{
		row:             row,
		col:             col,
		next:            nil,
		uniqueLocations: map[string]struct{}{},
	}
	k.addUniqueLocation()
	return k
}

func (k *Knot) addTail() {
	if k.next == nil {
		k.next = NewKnot(k.row, k.col)
	} else {
		k.next.addTail()
	}
}

func (k *Knot) isNeighborOfNext() bool {
	isLeffNeighbor := k.row == k.next.row && k.col == k.next.col-1
	isRightNeighbor := k.row == k.next.row && k.col == k.next.col+1
	isTopNeighbor := k.col == k.next.col && k.row == k.next.row-1
	isBottomNeighbor := k.col == k.next.col && k.row == k.next.row+1
	isTopLeftNeighbor := k.col == k.next.col-1 && k.row == k.next.row-1
	isTopRightNeighbor := k.col == k.next.col+1 && k.row == k.next.row-1
	isBottomLeftNeighbor := k.col == k.next.col-1 && k.row == k.next.row+1
	isBottomRightNeighbor := k.col == k.next.col+1 && k.row == k.next.row+1
	isOnSameSpot := k.col == k.next.col && k.row == k.next.row

	return isLeffNeighbor || isRightNeighbor || isTopNeighbor || isBottomNeighbor || isTopLeftNeighbor || isTopRightNeighbor || isBottomLeftNeighbor || isBottomRightNeighbor || isOnSameSpot
}

func (k *Knot) addUniqueLocation() {
	k.uniqueLocations[fmt.Sprintf("%d,%d", k.row, k.col)] = struct{}{}
}

func (k *Knot) setPosition(row int, col int) {
	k.row = row
	k.col = col
	k.addUniqueLocation()

	if k.next == nil || k.isNeighborOfNext() {
		return
	}

	prevRow := k.next.row
	prevCol := k.next.col

	if k.row == prevRow { // current node moved vertically
		if k.col > prevCol {
			k.next.moveRight()
		} else {
			k.next.moveLeft()
		}
	} else if k.col == prevCol { // current node moved horizontally
		if k.row > prevRow {
			k.next.moveDown()
		} else {
			k.next.moveUp()
		}
	} else { // we need to move the next node diagonally
		if k.row > prevRow {
			if k.col > prevCol {
				k.next.moveDownRight()
			} else {
				k.next.moveDownLeft()
			}
		} else {
			if k.col > prevCol {
				k.next.moveUpRight()
			} else {
				k.next.moveUpLeft()
			}
		}
	}
}

func (k *Knot) moveUp() {
	k.setPosition(k.row-1, k.col)
}

func (k *Knot) moveDown() {
	k.setPosition(k.row+1, k.col)
}

func (k *Knot) moveLeft() {
	k.setPosition(k.row, k.col-1)
}

func (k *Knot) moveRight() {
	k.setPosition(k.row, k.col+1)
}

func (k *Knot) moveUpLeft() {
	k.setPosition(k.row-1, k.col-1)
}

func (k *Knot) moveUpRight() {
	k.setPosition(k.row-1, k.col+1)
}

func (k *Knot) moveDownLeft() {
	k.setPosition(k.row+1, k.col-1)
}

func (k *Knot) moveDownRight() {
	k.setPosition(k.row+1, k.col+1)
}

func (k *Knot) move(direction string, distance int) {
	for i := 0; i < distance; i++ {
		switch direction {
		case "U":
			k.moveUp()
		case "D":
			k.moveDown()
		case "L":
			k.moveLeft()
		case "R":
			k.moveRight()
		}
	}
}

func (k *Knot) uniqueLocationsOfLastKnot() map[string]struct{} {
	if k.next == nil {
		return k.uniqueLocations
	}

	return k.next.uniqueLocationsOfLastKnot()
}

func (k *Knot) countOfUniqueLocationsOfLastKnot() int {
	return len(k.uniqueLocationsOfLastKnot())
}

func processRope(lines []string, knots int) (int, error) {
	head := NewKnot(0, 0)
	for i := 1; i < knots; i++ {
		head.addTail()
	}

	for _, line := range lines {
		lineParts := strings.Split(line, " ")
		direction := lineParts[0]
		distance, err := strconv.Atoi(lineParts[1])
		if err != nil {
			return 0, err
		}

		head.move(direction, distance)
	}

	return head.countOfUniqueLocationsOfLastKnot(), nil
}

func part1(input string) (int, error) {
	lines, err := readAllLines(input)
	if err != nil {
		return 0, err
	}

	return processRope(lines, 2)
}

func part2(input string) (int, error) {
	lines, err := readAllLines(input)
	if err != nil {
		return 0, err
	}

	return processRope(lines, 10)
}
