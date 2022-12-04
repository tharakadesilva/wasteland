package day3

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

func resolveItems(compartment string) map[rune]struct{} {
	items := map[rune]struct{}{}
	for _, item := range compartment {
		items[item] = struct{}{}
	}
	return items
}

func resolveRuneValue(r rune) int {
	if r >= 'a' && r <= 'z' {
		return int(r - 'a' + 1)
	}
	return int(r - 'A' + 27)
}

func part1(input string) (int, error) {
	lines, err := readAllLines(input)
	if err != nil {
		return 0, err
	}

	totalPoints := 0

	for _, line := range lines {
		length := len(line)
		compartment1 := line[:length/2]
		compartment2 := line[length/2:]

		compartment1Items := resolveItems(compartment1)
		compartment2Items := resolveItems(compartment2)

		for item := range compartment1Items {
			if _, ok := compartment2Items[item]; ok {
				totalPoints += resolveRuneValue(item)
			}
		}
	}

	return totalPoints, nil
}

func part2(input string) (int, error) {
	lines, err := readAllLines(input)
	if err != nil {
		return 0, err
	}

	totalPoints := 0

	for index := 0; index < len(lines)-1; index += 3 {
		backpack1Items := resolveItems(lines[index])
		backpack2Items := resolveItems(lines[index+1])
		backpack3Items := resolveItems(lines[index+2])

		for item := range backpack1Items {
			_, backpack2HasItem := backpack2Items[item]
			_, backpack3HasItem := backpack3Items[item]
			if backpack2HasItem && backpack3HasItem {
				totalPoints += resolveRuneValue(item)
			}
		}
	}

	return totalPoints, nil
}
