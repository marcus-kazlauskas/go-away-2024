package puzzles

import (
	"bufio"
	"os"
)

func Year2024Day1Part1(file *os.File) (*string, error) {
	file.Seek(0, 0)
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()
	// TODO: puzzle solution
	ans := line
	return &ans, nil
}
