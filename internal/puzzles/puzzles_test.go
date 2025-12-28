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

	t.Run("Year 2025 day 1 part 1", func(t *testing.T) {
		file, err := os.Open("year2025day1_test.txt")
		a.Nil(err)
		scan := bufio.NewScanner(file)
		ans, err := Year2025Day1Part1(scan)

		a.Nil(err)
		a.NotNil(ans)
		a.Equal("7", *ans)
	})

	t.Run("Year 2025 day 1 part 2", func(t *testing.T) {
		file, err := os.Open("year2025day1_test.txt")
		a.Nil(err)
		scan := bufio.NewScanner(file)
		ans, err := Year2025Day1Part2(scan)

		a.Nil(err)
		a.NotNil(ans)
		a.Equal("16", *ans)
	})

	t.Run("Year 2025 day 2 part 1", func(t *testing.T) {
		file, err := os.Open("year2025day2_test.txt")
		a.Nil(err)
		scan := bufio.NewScanner(file)
		ans, err := Year2025Day2Part1(scan)

		a.Nil(err)
		a.NotNil(ans)
		a.Equal("1227775554", *ans)
	})

	t.Run("Year 2025 day 2 part 2", func(t *testing.T) {
		file, err := os.Open("year2025day2_test.txt")
		a.Nil(err)
		scan := bufio.NewScanner(file)
		ans, err := Year2025Day2Part2(scan)

		a.Nil(err)
		a.NotNil(ans)
		a.Equal("4174379265", *ans)
	})

	t.Run("Year 2025 day 3 part 1", func(t *testing.T) {
		file, err := os.Open("year2025day3_test.txt")
		a.Nil(err)
		scan := bufio.NewScanner(file)
		ans, err := Year2025Day3Part1(scan)

		a.Nil(err)
		a.NotNil(ans)
		a.Equal("357", *ans)
	})

	t.Run("Year 2025 day 3 part 2", func(t *testing.T) {
		file, err := os.Open("year2025day3_test.txt")
		a.Nil(err)
		scan := bufio.NewScanner(file)
		ans, err := Year2025Day3Part2(scan)

		a.Nil(err)
		a.NotNil(ans)
		a.Equal("3121910778619", *ans)
	})

	t.Run("Year 2025 day 4 part 1", func(t *testing.T) {
		file, err := os.Open("year2025day4_test.txt")
		a.Nil(err)
		scan := bufio.NewScanner(file)
		ans, err := Year2025Day4Part1(scan)

		a.Nil(err)
		a.NotNil(ans)
		a.Equal("13", *ans)
	})

	t.Run("Year 2025 day 4 part 2", func(t *testing.T) {
		file, err := os.Open("year2025day4_test.txt")
		a.Nil(err)
		scan := bufio.NewScanner(file)
		ans, err := Year2025Day4Part2(scan)

		a.Nil(err)
		a.NotNil(ans)
		a.Equal("43", *ans)
	})

	t.Run("Year 2025 day 5 part 1", func(t *testing.T) {
		file, err := os.Open("year2025day5_test.txt")
		a.Nil(err)
		scan := bufio.NewScanner(file)
		ans, err := Year2025Day5Part1(scan)

		a.Nil(err)
		a.NotNil(ans)
		a.Equal("3", *ans)
	})

	t.Run("Year 2025 day 5 part 2", func(t *testing.T) {
		file, err := os.Open("year2025day5_test.txt")
		a.Nil(err)
		scan := bufio.NewScanner(file)
		ans, err := Year2025Day5Part2(scan)

		a.Nil(err)
		a.NotNil(ans)
		a.Equal("18", *ans)
	})

	t.Run("Year 2025 day 6 part 1", func(t *testing.T) {
		file, err := os.Open("year2025day6_test.txt")
		a.Nil(err)
		scan := bufio.NewScanner(file)
		ans, err := Year2025Day6Part1(scan)

		a.Nil(err)
		a.NotNil(ans)
		a.Equal("4277556", *ans)
	})
}
