package day7

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

type File struct {
	Parent   *File
	Path     string
	isDir    bool
	size     int
	children map[string]*File
}

func NewRootFile() *File {
	return &File{
		Path:     "/",
		isDir:    true,
		children: make(map[string]*File),
	}
}

func NewDirFile(parent *File, path string) *File {
	return &File{
		Parent:   parent,
		Path:     path,
		isDir:    true,
		children: make(map[string]*File),
	}
}

func NewFile(parent *File, path string, size int) *File {
	return &File{
		Parent: parent,
		Path:   path,
		size:   size,
		isDir:  false,
	}
}

func (f *File) IncreaseSize(size int) {
	f.size += size
	if f.Parent != nil {
		f.Parent.IncreaseSize(size)
	}
}

func (f *File) AddChild(child *File) *File {
	f.children[child.Path] = child
	f.IncreaseSize(child.size)
	return f
}

func (f *File) GetChild(path string) *File {
	return f.children[path]
}

func isCdCommand(line string) bool {
	return strings.HasPrefix(line, "$ cd ")
}

func isCdBackCommand(line string) bool {
	return line == "$ cd .."
}

func isCdDirCommand(line string) bool {
	return isCdCommand(line) && !isCdBackCommand(line)
}

func parseCdCommand(line string) string {
	return line[5:]
}

func isLsCommand(line string) bool {
	return strings.HasPrefix(line, "$ ls")
}

func isLsOutput(line string) bool {
	return !isCdCommand(line) && !isLsCommand(line)
}

func isLsOutputForFile(line string) bool {
	return isLsOutput(line) && !strings.HasPrefix(line, "dir ")
}

func stringToInt(s string) (int, error) {
	return strconv.Atoi(s)
}

func parseLsOutputForFile(line string) (string, int, error) {
	parts := strings.Split(line, " ")
	size, err := stringToInt(parts[0])
	if err != nil {
		return "", 0, err
	}

	return parts[1], size, nil
}

func buildFileTree(lines []string, root *File, dirs []*File) (*File, []*File, error) {
	current := root

	for _, line := range lines {
		if isCdDirCommand(line) {
			dirName := parseCdCommand(line)
			dir := NewDirFile(current, dirName)
			dirs = append(dirs, dir)
			current.AddChild(dir)
			current = dir
		} else if isCdBackCommand(line) {
			current = current.Parent
		} else if isLsOutputForFile(line) {
			fileName, size, err := parseLsOutputForFile(line)
			if err != nil {
				return nil, nil, err
			}

			current.AddChild(NewFile(current, fileName, size))
		}
	}

	return root, dirs, nil
}

func part1(input string) (int, error) {
	lines, err := readAllLines(input)
	if err != nil {
		return 0, err
	}

	root := NewRootFile()
	dirs := []*File{root}

	root, dirs, err = buildFileTree(lines[2:], root, dirs)
	if err != nil {
		return 0, err
	}

	sum := 0
	for _, dir := range dirs {
		if dir.size < 100000 {
			sum += dir.size
		}
	}
	return sum, nil
}

func part2(input string) (int, error) {
	lines, err := readAllLines(input)
	if err != nil {
		return 0, err
	}

	root := NewRootFile()
	dirs := []*File{root}

	root, dirs, err = buildFileTree(lines[2:], root, dirs)
	if err != nil {
		return 0, err
	}

	currentFree := 70000000 - root.size

	min := math.MaxInt
	for _, dir := range dirs {
		if (currentFree+dir.size) > 30000000 && dir.size < min {
			min = dir.size
		}
	}

	return min, nil
}
