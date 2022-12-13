package day11

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func readFile(path string) (string, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func splitStringByEmptyLine(s string) []string {
	return strings.Split(s, "\n\n")
}

type Monkey struct {
	items          []int
	operation      string
	testDivisor    int
	successPasser  int
	failPasser     int
	inspectedCount int
}

func EvaluateMathOp(op string) (int, error) {
	parts := strings.Split(op, " ")

	a, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, err
	}

	b, err := strconv.Atoi(parts[2])
	if err != nil {
		return 0, err
	}

	if parts[1] == "*" {
		return a * b, nil
	}

	return a + b, nil
}

func (m *Monkey) EvaluateOperation(item int) (int, error) {
	operation := strings.Replace(m.operation, "old", fmt.Sprint(item), -1)
	return EvaluateMathOp(operation)
}

func parseitems(s string) ([]int, error) {
	parts := strings.Split(s, ": ")
	parts = strings.Split(parts[1], ", ")
	items := make([]int, len(parts))
	for i, part := range parts {
		item, err := strconv.Atoi(part)
		if err != nil {
			return nil, err
		}

		items[i] = item
	}
	return items, nil
}

func parseOperation(s string) string {
	parts := strings.Split(s, "new = ")
	return parts[1]
}

func parseTestDivisor(s string) (int, error) {
	parts := strings.Split(s, "Test: divisible by ")
	divisor, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, err
	}

	return divisor, nil
}

func parsePasser(s string) (int, error) {
	parts := strings.Split(s, "throw to monkey ")
	return strconv.Atoi(parts[1])
}

func NewMonkey(s string) (*Monkey, error) {
	lines := strings.Split(s, "\n")

	items, err := parseitems(lines[1])
	if err != nil {
		return nil, err
	}

	operation := parseOperation(lines[2])

	testDivisor, err := parseTestDivisor(lines[3])
	if err != nil {
		return nil, err
	}

	successPasser, err := parsePasser(lines[4])
	if err != nil {
		return nil, err
	}

	failPasser, err := parsePasser(lines[5])
	if err != nil {
		return nil, err
	}

	return &Monkey{
		items:          items,
		operation:      operation,
		testDivisor:    testDivisor,
		successPasser:  successPasser,
		failPasser:     failPasser,
		inspectedCount: 0,
	}, nil
}

func productOfAllDivisors(monkeys []*Monkey) int {
	product := 1
	for _, monkey := range monkeys {
		product *= monkey.testDivisor
	}
	return product
}

func findTwoHighestInspectedCounts(monkeys []*Monkey) (int, int) {
	highest, secondHighest := 0, 0
	for _, monkey := range monkeys {
		if monkey.inspectedCount > highest {
			secondHighest = highest
			highest = monkey.inspectedCount
		} else if monkey.inspectedCount > secondHighest {
			secondHighest = monkey.inspectedCount
		}
	}
	return highest, secondHighest
}

func monkeyBusiness(input string, rounds int, worryLevelManageable bool) (int, error) {
	s, err := readFile(input)
	if err != nil {
		return 0, err
	}

	monkeyStrings := splitStringByEmptyLine(s)
	monkeys := make([]*Monkey, len(monkeyStrings))
	for i, monkeyString := range monkeyStrings {
		monkeys[i], err = NewMonkey(monkeyString)
		if err != nil {
			return 0, err
		}
	}

	productOfAllDivisors := productOfAllDivisors(monkeys)

	for i := 0; i < rounds; i++ {
		for _, monkey := range monkeys {
			for len(monkey.items) > 0 {
				monkey.inspectedCount += 1

				item := monkey.items[0]
				monkey.items = monkey.items[1:]

				result, err := monkey.EvaluateOperation(item)
				if err != nil {
					return 0, err
				}

				if worryLevelManageable {
					result /= 3
				}

				result %= productOfAllDivisors

				if result%monkey.testDivisor == 0 {
					monkeys[monkey.successPasser].items = append(monkeys[monkey.successPasser].items, result)
				} else {
					monkeys[monkey.failPasser].items = append(monkeys[monkey.failPasser].items, result)
				}
			}
		}
	}

	highest, secondHighest := findTwoHighestInspectedCounts(monkeys)
	return highest * secondHighest, nil
}

func part1(input string) (int, error) {
	return monkeyBusiness(input, 20, true)
}

func part2(input string) (int, error) {
	return monkeyBusiness(input, 10000, false)
}
