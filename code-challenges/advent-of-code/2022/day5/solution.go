package day5

import (
	"io/ioutil"
	"log"
	"math"
	"regexp"
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

func addToBottomOfStack(stack []string, s string) []string {
	return append(stack, s)
}

func addToTopOfStack(stack []string, s string) []string {
	return append([]string{s}, stack...)
}

func removeFromTopOfStack(stack []string) ([]string, string) {
	return stack[1:], stack[0]
}

func peekTopOfStack(stack []string) string {
	return stack[0]
}

func joinSliceToString(slice []string) string {
	return strings.Join(slice, "")
}

func peekTopOfAllStacks(stacks [][]string) string {
	topOfStacks := make([]string, len(stacks))
	for i, stack := range stacks {
		topOfStacks[i] = peekTopOfStack(stack)
	}
	return joinSliceToString(topOfStacks)
}

func parseLine(line string, stacks [][]string) [][]string {
	for i := 1; i < len(line)-1; i += 4 {
		container := line[i]
		if container != ' ' {
			stackIndex := (i - 1) / 4
			stacks[stackIndex] = addToBottomOfStack(stacks[stackIndex], string(container))
		}
	}

	return stacks
}

func stringToInt(s string) (int, error) {
	return strconv.Atoi(s)
}

func parseInstruction(instruction string) (int, int, int, error) {
	instructionR, err := regexp.Compile(`move (\d+) from (\d+) to (\d+)`)
	if err != nil {
		return 0, 0, 0, err
	}

	matches := instructionR.FindStringSubmatch(instruction)
	noOfContainers, err := stringToInt(matches[1])
	if err != nil {
		return 0, 0, 0, err
	}

	fromStack, err := stringToInt(matches[2])
	if err != nil {
		return 0, 0, 0, err
	}

	toStack, err := stringToInt(matches[3])
	if err != nil {
		return 0, 0, 0, err
	}

	return noOfContainers, fromStack - 1, toStack - 1, nil
}

func executeInstruction9000(instruction string, stacks [][]string) ([][]string, error) {
	noOfContainers, fromStack, toStack, err := parseInstruction(instruction)
	if err != nil {
		return nil, err
	}

	for i := 0; i < noOfContainers; i++ {
		var value string
		stacks[fromStack], value = removeFromTopOfStack(stacks[fromStack])
		stacks[toStack] = addToTopOfStack(stacks[toStack], value)
	}

	return stacks, nil
}

func part1(input string) (string, error) {
	lines, err := readAllLines(input)
	if err != nil {
		return "", err
	}

	noOfStacks := (len(lines[0]) / 4) + 1
	stacks := make([][]string, noOfStacks)

	instructionStartFrom := math.MinInt

	for i, line := range lines {
		if line[1] == '1' {
			instructionStartFrom = i + 2
			break
		}

		stacks = parseLine(line, stacks)
	}

	for i := instructionStartFrom; i < len(lines); i++ {
		stacks, err = executeInstruction9000(lines[i], stacks)
		if err != nil {
			return "", err
		}
	}

	return peekTopOfAllStacks(stacks), nil
}

func removeNElemementsFromTopOfStack(stack []string, n int) ([]string, []string) {
	return stack[n:], stack[:n]
}

func addSliceToTopOfStack(stack []string, slice []string) []string {
	s := []string{}
	s = append(s, slice...)
	s = append(s, stack...)
	return s
}

func executeInstruction9001(instruction string, stacks [][]string) ([][]string, error) {
	noOfContainers, fromStack, toStack, err := parseInstruction(instruction)
	if err != nil {
		return nil, err
	}

	var values []string
	stacks[fromStack], values = removeNElemementsFromTopOfStack(stacks[fromStack], noOfContainers)
	stacks[toStack] = addSliceToTopOfStack(stacks[toStack], values)

	return stacks, nil
}

func part2(input string) (string, error) {
	lines, err := readAllLines(input)
	if err != nil {
		return "", err
	}

	noOfStacks := (len(lines[0]) / 4) + 1
	stacks := make([][]string, noOfStacks)

	instructionStartFrom := math.MinInt

	for i, line := range lines {
		if line[1] == '1' {
			instructionStartFrom = i + 2
			break
		}

		stacks = parseLine(line, stacks)
	}

	for i := instructionStartFrom; i < len(lines); i++ {
		log.Println(stacks)
		stacks, err = executeInstruction9001(lines[i], stacks)
		if err != nil {
			return "", err
		}
	}
	log.Println(stacks)

	return peekTopOfAllStacks(stacks), nil
}
