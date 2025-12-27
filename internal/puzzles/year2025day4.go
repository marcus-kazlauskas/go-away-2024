package puzzles

import (
	"bufio"
	"fmt"

	"github.com/gofiber/fiber/v2/log"
)

func Year2025Day4Part1(scan *bufio.Scanner) (*string, error) {
	var ans int = 0

	grid, err := getGrid(scan)
	if err != nil {
		return nil, err
	}

	ans = countRolls(grid)

	result := fmt.Sprint(ans)
	return &result, nil
}

func getGrid(scan *bufio.Scanner) ([][]bool, error) {
	grid := make([][]bool, 0)

	scan.Scan()
	line := scan.Text()
	n := len(line)
	grid = append(grid, appendStubLine(n))

	l, err := appendLine(line)
	if err != nil {
		return nil, err
	}
	grid = append(grid, l)

	for scan.Scan() {
		line = scan.Text()
		l, err = appendLine(line)
		if err != nil {
			return nil, err
		}
		grid = append(grid, l)
	}

	grid = append(grid, appendStubLine(n))
	return grid, nil
}

func appendStubLine(n int) []bool {
	l := make([]bool, 0)
	for range n + 2 {
		l = append(l, false)
	}
	return l
}

func appendLine(line string) ([]bool, error) {
	l := make([]bool, 0)
	l = append(l, false)
	for _, r := range line {
		if r != '.' && r != '@' {
			return nil, DataError()
		}
		if r == '@' {
			l = append(l, true)
		} else {
			l = append(l, false)
		}
	}
	l = append(l, false)
	return l, nil
}

func countRolls(grid [][]bool) int {
	count := 0
	for i := 1; i <= len(grid)-2; i++ {
		for j := 1; j <= len(grid[0])-2; j++ {
			if checkRoll(grid, i, j) {
				count++
			}
		}
	}
	return count
}

func checkRoll(grid [][]bool, i int, j int) bool {
	if !grid[i][j] {
		return false
	}
	rollsCount := 0

	if grid[i-1][j-1] {
		rollsCount++
	}
	if grid[i-1][j] {
		rollsCount++
	}
	if grid[i-1][j+1] {
		rollsCount++
	}
	if grid[i][j+1] {
		rollsCount++
	}
	if grid[i+1][j+1] {
		rollsCount++
	}
	if grid[i+1][j] {
		rollsCount++
	}
	if grid[i+1][j-1] {
		rollsCount++
	}
	if grid[i][j-1] {
		rollsCount++
	}

	if rollsCount < 4 {
		log.Debugf("Roll accessible: i=%d j=%d", i-1, j-1)
		return true
	} else {
		return false
	}
}

func Year2025Day4Part2(scan *bufio.Scanner) (*string, error) {
	var ans int = 0

	grid, err := getGrid(scan)
	if err != nil {
		return nil, err
	}

	count, rolls := countRollsV2(grid)
	ans += count
	for count > 0 {
		log.Debugf("Removed %d rolls", count)
		removeRolls(grid, rolls)
		count, rolls = countRollsV2(grid)
		ans += count
	}

	result := fmt.Sprint(ans)
	return &result, nil
}

func countRollsV2(grid [][]bool) (int, [][]int) {
	count := 0
	rolls := make([][]int, 0)
	for i := 1; i <= len(grid)-2; i++ {
		for j := 1; j <= len(grid[0])-2; j++ {
			if checkRoll(grid, i, j) {
				count++
				roll := make([]int, 0)
				roll = append(roll, i)
				roll = append(roll, j)
				rolls = append(rolls, roll)
			}
		}
	}
	return count, rolls
}

func removeRolls(grid [][]bool, rolls [][]int) {
	for _, p := range rolls {
		i := p[0]
		j := p[1]
		grid[i][j] = false
	}
}
