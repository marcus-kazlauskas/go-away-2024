package puzzles

import (
	"go-away-2024/internal/minio"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPuzzles(t *testing.T) {
	a := assert.New(t)

	t.Run("Year 2024 day 1 part 1", func(t *testing.T) {
		file, err := os.Open("year2024day1part1_test.txt")
		a.Nil(err)
		scan, err := minio.NewScanner(file)
		a.Nil(err)
		ans, err := Year2024Day1Part1(scan)

		a.Nil(err)
		a.Equal("11", *ans)
	})
}
