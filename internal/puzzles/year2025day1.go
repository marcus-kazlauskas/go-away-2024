package puzzles

import (
	"bufio"
	"fmt"
	"go-away-2024/internal/utils"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/gofiber/fiber/v2/log"
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
			return nil, DataError()
		}

		switch firstRune {
		case 'L':
			pos = rotate(pos, -rotation)
		case 'R':
			pos = rotate(pos, rotation)
		default:
			return nil, DataError()
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
		return (100 + delta%100) % 100
	} else {
		return delta % 100
	}
}

func Year2025Day1Part2(scan *bufio.Scanner) (*string, error) {
	var ans int = 0
	var pos int = 50
	var clicks int = 0

	for scan.Scan() {
		line := scan.Text()

		firstRune, _ := utf8.DecodeRuneInString(line)
		rotationStr := strings.Trim(line, "LR")
		rotation, err := strconv.Atoi(rotationStr)
		if err != nil {
			return nil, DataError()
		}

		switch firstRune {
		case 'L':
			pos, clicks = rotateV2(pos, -rotation)
		case 'R':
			pos, clicks = rotateV2(pos, rotation)
		default:
			return nil, DataError()
		}

		ans += clicks
	}

	result := fmt.Sprint(ans)
	return &result, nil
}

func rotateV2(position int, rotation int) (pos int, clicks int) {
	delta := position + rotation
	clicks = utils.Abs(delta) / 100
	if delta <= 0 {
		pos = (100 + delta%100) % 100
		if position != 0 {
			clicks++
		}
	} else {
		pos = delta % 100
	}
	log.Debugf(
		"Rotate: position=%d rotation=%d -> pos=%d clicks=%d",
		position, rotation, pos, clicks,
	)
	return
}
