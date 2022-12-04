package day4

import (
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

func toInt(s string) (int, error) {
	return strconv.Atoi(s)
}

func resolveBounds(assignments string) (int, int, error) {
	parts := strings.Split(assignments, "-")
	left, err := toInt(parts[0])
	if err != nil {
		return 0, 0, err
	}

	right, err := toInt(parts[1])
	if err != nil {
		return 0, 0, err
	}

	return left, right, nil
}

func part1(input string) (int, error) {
	lines, err := readAllLines(input)
	if err != nil {
		return 0, err
	}

	count := 0

	for _, line := range lines {
		assignments := strings.Split(line, ",")

		firstAssignmentLeft, firstAssignmentRight, err := resolveBounds(assignments[0])
		if err != nil {
			return 0, err
		}

		secondAssignmentLeft, secondAssignmentRight, err := resolveBounds(assignments[1])
		if err != nil {
			return 0, err
		}

		if (firstAssignmentLeft <= secondAssignmentLeft &&
			firstAssignmentRight >= secondAssignmentRight) ||
			(secondAssignmentLeft <= firstAssignmentLeft &&
				secondAssignmentRight >= firstAssignmentRight) {
			count++
		}
	}

	return count, nil
}

func part2(input string) (int, error) {
	lines, err := readAllLines(input)
	if err != nil {
		return 0, err
	}

	count := 0

	for _, line := range lines {
		assignments := strings.Split(line, ",")

		firstAssignmentLeft, firstAssignmentRight, err := resolveBounds(assignments[0])
		if err != nil {
			return 0, err
		}

		secondAssignmentLeft, secondAssignmentRight, err := resolveBounds(assignments[1])
		if err != nil {
			return 0, err
		}

		if (firstAssignmentLeft <= secondAssignmentLeft &&
			firstAssignmentRight >= secondAssignmentLeft) ||
			(secondAssignmentLeft <= firstAssignmentLeft &&
				secondAssignmentRight >= firstAssignmentLeft) {
			count++
		}
	}

	return count, nil
}
