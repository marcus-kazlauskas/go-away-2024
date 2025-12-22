package puzzles

import (
	"bufio"
	"fmt"
	"go-away-2024/internal/minio"
	"strconv"
	"strings"
	"unicode/utf8"
)

func Year2025Day1Part1(scan *bufio.Scanner) (*string, error) {
	var ans int = 0
	var pos int = 50

	for scan.Scan() {
		line := scan.Text()

		firstRune, _ := utf8.DecodeRuneInString(line)
		rotationStr := strings.Trim(line, "LR")
		rotation, err := strconv.Atoi(rotationStr)
		if err != nil {
			return nil, minio.DataError()
		}

		switch firstRune {
		case 'L':
			pos = rotate(pos, -rotation)
		case 'R':
			pos = rotate(pos, rotation)
		default:
			return nil, minio.DataError()
		}

		if pos == 0 {
			ans++
		}
	}

	result := fmt.Sprint(ans)
	return &result, nil
}

func rotate(position int, rotation int) int {
	delta := position + rotation
	if delta < 0 {
		return (100 + delta) % 100
	} else if delta > 99 {
		return (delta - 100) % 100
	} else {
		return delta
	}
}
