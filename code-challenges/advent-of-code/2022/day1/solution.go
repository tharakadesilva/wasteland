package day1

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

func toInt(s string) (int, error) {
	return strconv.Atoi(s)
}

func part1(input string) (int, error) {
	lines, err := readAllLines(input)
	if err != nil {
		return 0, err
	}

	max := math.MinInt
	groupSum := 0

	for _, line := range lines {
		if line == "" {
			if groupSum > max {
				max = groupSum
			}
			groupSum = 0
		} else {
			num, err := toInt(line)
			if err != nil {
				return 0, err
			}

			groupSum += num
		}
	}

	return max, err
}

func part2(input string) (int, error) {
	lines, err := readAllLines(input)
	if err != nil {
		return 0, err
	}

	max, secondMax, thirdMax := math.MinInt, math.MinInt, math.MinInt
	groupSum := 0

	for _, line := range lines {
		if line == "" {
			if groupSum > thirdMax {
				thirdMax = groupSum
			}

			if thirdMax > secondMax {
				secondMax, thirdMax = thirdMax, secondMax
			}

			if secondMax > max {
				max, secondMax = secondMax, max
			}

			groupSum = 0
		} else {
			num, err := toInt(line)
			if err != nil {
				return 0, err
			}

			groupSum += num
		}
	}

	return max + secondMax + thirdMax, nil
}
