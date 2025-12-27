package puzzles

import (
	"bufio"
	"fmt"
	"go-away-2024/internal/utils"
	"unicode"
	"unicode/utf8"

	"github.com/gofiber/fiber/v2/log"
)

const NUMBER_OF_BATTERIES = 12

func Year2025Day3Part1(scan *bufio.Scanner) (*string, error) {
	var ans int = 0

	for scan.Scan() {
		line := scan.Text()
		if utf8.RuneCountInString(line) < 2 {
			return nil, DataError()
		}

		bank := make([]int, 0)
		for _, r := range line {
			var joltage int
			if unicode.IsDigit(r) {
				joltage = int(r - '0')
			} else {
				return nil, DataError()
			}
			bank = append(bank, joltage)
		}

		lMax := 0
		for l := 1; l <= len(bank)-2; l++ {
			if bank[l] > bank[lMax] {
				lMax = l
			}
		}
		rMax := len(bank) - 1
		for r := len(bank) - 2; r > lMax; r-- {
			if bank[r] > bank[rMax] {
				rMax = r
			}
		}

		ans += bank[lMax]*10 + bank[rMax]
	}

	result := fmt.Sprint(ans)
	return &result, nil
}

func Year2025Day3Part2(scan *bufio.Scanner) (*string, error) {
	var ans int64 = 0

	for scan.Scan() {
		line := scan.Text()
		if utf8.RuneCountInString(line) < NUMBER_OF_BATTERIES {
			return nil, DataError()
		}

		bank := make([]int, 0)
		for _, r := range line {
			var joltage int
			if unicode.IsDigit(r) {
				joltage = int(r - '0')
			} else {
				return nil, DataError()
			}
			bank = append(bank, joltage)
		}

		maxJoltage := getMaxJoltage(bank)
		ans += maxJoltage
	}

	result := fmt.Sprint(ans)
	return &result, nil
}

func getMaxJoltage(bank []int) int64 {
	newBank := make([]int, 0)
	maxPos := -1
	for i := range NUMBER_OF_BATTERIES {
		maxPos++
		maxJoltage := bank[maxPos]
		for j := maxPos + 1; j <= len(bank)-NUMBER_OF_BATTERIES+i; j++ {
			if bank[j] > maxJoltage {
				maxPos = j
				maxJoltage = bank[j]
			}
		}
		newBank = append(newBank, maxJoltage)
	}
	log.Debugf("Max joltage: %v -> %v", bank, newBank)
	return getJoltage(newBank)
}

func getJoltage(bank []int) (res int64) {
	i := NUMBER_OF_BATTERIES - 1
	for _, j := range bank {
		res += int64(j) * utils.Pow64(10, i)
		i--
	}
	return
}
