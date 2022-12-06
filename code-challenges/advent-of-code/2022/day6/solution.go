package day6

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

func findFirstSubstringOfNUniqueCharacters(s string, n int) int {
	for i := 0; i < len(s)-n; i++ {
		unique := make(map[byte]bool)
		for j := i + 1; j < i+n; j++ {
			if s[i] == s[j] {
				break
			}
			unique[s[j]] = true
		}
		if len(unique) == n-1 {
			return i
		}
	}
	return -1
}

func part1(input string) (int, error) {
	lines, err := readAllLines(input)
	if err != nil {
		return 0, err
	}

	line := lines[0]
	return findFirstSubstringOfNUniqueCharacters(line, 4) + 4, nil
}

func part2(input string) (int, error) {
	lines, err := readAllLines(input)
	if err != nil {
		return 0, err
	}

	line := lines[0]
	return findFirstSubstringOfNUniqueCharacters(line, 14) + 14, nil
}
