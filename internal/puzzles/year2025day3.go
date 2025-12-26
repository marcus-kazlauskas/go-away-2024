package puzzles

import (
	"bufio"
	"fmt"
	"unicode"
)

func Year2025Day3Part1(scan *bufio.Scanner) (*string, error) {
	var ans int = 0

	for scan.Scan() {
		line := scan.Text()
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
