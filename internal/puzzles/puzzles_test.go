package puzzles

import (
	"bufio"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPuzzles(t *testing.T) {
	a := assert.New(t)

	t.Run("Year 2024 day 1 part 1", func(t *testing.T) {
		file, err := os.Open("year2024day1_test.txt")
		a.Nil(err)
		scan := bufio.NewScanner(file)
		ans, err := Year2024Day1Part1(scan)

		a.Nil(err)
		a.NotNil(ans)
		a.Equal("11", *ans)
	})

	t.Run("Year 2024 day 1 part 2", func(t *testing.T) {
		file, err := os.Open("year2024day1_test.txt")
		a.Nil(err)
		scan := bufio.NewScanner(file)
		ans, err := Year2024Day1Part2(scan)

		a.Nil(err)
		a.NotNil(ans)
		a.Equal("31", *ans)
	})

	t.Run("Year 2025 day 1", func(t *testing.T) {
		file, err := os.Open("year2025day1_test.txt")
		a.Nil(err)
		scan := bufio.NewScanner(file)
		ans, err := Year2025Day1Part1(scan)

		a.Nil(err)
		a.NotNil(ans)
		a.Equal("4", *ans)
	})
}
