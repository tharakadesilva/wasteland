package day10

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

func isNoop(line string) bool {
	return line == "noop"
}

func resolveXStr(line string) string {
	return strings.Split(line, " ")[1]
}

func resolveX(line string) (int, error) {
	return strconv.Atoi(resolveXStr(line))
}

func shouldCalculateSignalStrength(cycle int) bool {
	return (cycle+20)%40 == 0
}

type Machine struct {
	x              int
	cycle          int
	signalStrength int
	screen         []rune
}

func (m *Machine) GetSpriteLowerBound() int {
	return m.x - 1
}

func (m *Machine) GetSpriteUpperBound() int {
	return m.x + 1
}

func (m *Machine) IsWithinSpriteRange() bool {
	return m.cycle%40 >= m.GetSpriteLowerBound() && m.cycle%40 <= m.GetSpriteUpperBound()
}

func (m *Machine) IncreaseCycle() {
	if m.IsWithinSpriteRange() {
		m.screen[m.cycle] = '#'
	}

	m.cycle += 1

	if shouldCalculateSignalStrength(m.cycle) {
		m.signalStrength += m.x * m.cycle
	}
}

func (m *Machine) SprintScreen() string {
	var sb strings.Builder
	for i, r := range m.screen {
		if i%40 == 0 {
			sb.WriteString("\n")
		}
		sb.WriteRune(r)
	}
	return sb.String()
}

func initScreen() []rune {
	screen := make([]rune, 240)
	for i := range screen {
		screen[i] = '.'
	}
	return screen
}

func NewMachine() *Machine {
	return &Machine{
		x:              0,
		cycle:          0,
		signalStrength: 0,
		screen:         initScreen(),
	}
}

func runMachine(input string) (*Machine, error) {
	lines, err := readAllLines(input)
	if err != nil {
		return nil, err
	}

	m := NewMachine()
	nextX := 1

	for _, line := range lines {
		m.x = nextX
		m.IncreaseCycle()

		if !isNoop(line) {
			m.IncreaseCycle()

			dx, err := resolveX(line)
			if err != nil {
				return nil, err
			}

			nextX = m.x + dx
		}
	}

	return m, nil
}

func part1(input string) (int, error) {
	m, err := runMachine(input)
	if err != nil {
		return 0, err
	}

	return m.signalStrength, nil
}

func part2(input string) (string, error) {
	m, err := runMachine(input)
	if err != nil {
		return "", err
	}

	return m.SprintScreen(), nil
}
