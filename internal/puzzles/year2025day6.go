package puzzles

import (
	"bufio"
	"fmt"
	"go-away-2024/internal/utils"
	"regexp"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2/log"
)

func Year2025Day6Part1(scan *bufio.Scanner) (*string, error) {
	var ans int64 = 0
	input := make([][]string, 0)
	re := regexp.MustCompile(`\s+`)

	for scan.Scan() {
		line := scan.Text()
		values := re.Split(strings.Trim(line, " "), -1)
		input = append(input, values)
	}
	log.Debugf("%v", input)

	for i := range len(input[len(input)-1]) {
		nums := make([]int64, 0)
		for j := range len(input) - 1 {
			num, err := strconv.ParseInt(input[j][i], 10, 64)
			if err != nil {
				return nil, DataError()
			}
			nums = append(nums, num)
		}

		switch input[len(input)-1][i] {
		case "+":
			ans += sumArray(nums)
		case "*":
			ans += multiplyArray(nums)
		default:
			return nil, DataError()
		}
	}

	result := fmt.Sprint(ans)
	return &result, nil
}

func sumArray(nums []int64) int64 {
	var ans int64 = 0
	for _, num := range nums {
		ans += num
	}
	return ans
}

func multiplyArray(nums []int64) int64 {
	var ans int64 = 1
	for _, num := range nums {
		ans *= num
	}
	return ans
}

func Year2025Day6Part2(scan *bufio.Scanner) (*string, error) {
	var ans int64 = 0
	input := make([][]rune, 0)

	for scan.Scan() {
		line := []rune(scan.Text())
		input = append(input, line)
	}

	for i := 1; i < len(input); i++ {
		if len(input[i]) != len(input[0]) {
			return nil, DataError()
		}
	}

	pos := len(input[0]) - 1
	var problem [][]rune
	for pos >= 0 {
		problem, pos = readProblem(input, pos)
		log.Debugf("Read problem: %v", problem)
		switch getOperationFromProblem(problem) {
		case '+':
			delta := sumRuneArrays(reverseProblem(problem))
			log.Debugf("Sum of numbers is %d", delta)
			ans += delta
		case '*':
			delta := multiplyRuneArrays(reverseProblem(problem))
			log.Debugf("Multiplication of numbers is %d", delta)
			ans += delta
		default:
			return nil, DataError()
		}
	}

	result := fmt.Sprint(ans)
	return &result, nil
}

func readProblem(input [][]rune, pos int) ([][]rune, int) {
	problem := make([][]rune, 0)
	for range len(input) {
		line := make([]rune, 0)
		problem = append(problem, line)
	}

	for notEmptyColumn(input, pos) {
		for i := range len(input) {
			problem[i] = append(problem[i], input[i][pos])
		}
		pos--
	}
	pos--
	return problem, pos
}

func reverseProblem(problem [][]rune) (ans [][]rune) {
	buf := problem[0 : len(problem)-1]
	iLen := len(buf[0])
	jLen := len(buf)
	for i := range iLen {
		line := make([]rune, 0)
		for j := range jLen {
			line = append(line, buf[j][i])
		}
		ans = append(ans, line)
	}
	return
}

func getOperationFromProblem(problem [][]rune) rune {
	return problem[len(problem)-1][len(problem[0])-1]
}

func notEmptyColumn(input [][]rune, pos int) bool {
	if pos < 0 {
		return false
	}

	ans := false
	for i := range len(input) {
		ans = ans || (input[i][pos] != ' ')
	}
	return ans
}

func sumRuneArrays(nums [][]rune) int64 {
	var ans int64 = 0
	for _, r := range nums {
		if len(r) > 0 {
			ans += runeArrayToInt64(r)
		}
	}
	return ans
}

func multiplyRuneArrays(nums [][]rune) int64 {
	var ans int64 = 1
	for _, r := range nums {
		if len(r) > 0 {
			ans *= runeArrayToInt64(r)
		}
	}
	return ans
}

func runeArrayToInt64(num []rune) int64 {
	var ans int64 = 0
	for i, r := range num {
		var n int64 = 0
		if r != ' ' {
			n = int64(r - '0')
		}
		ans += int64(n) * utils.Pow64(10, len(num)-i-1)
	}
	for i := len(num) - 1; i >= 0; i-- {
		if num[i] != ' ' {
			break
		} else {
			ans /= 10
		}
	}
	log.Debugf("Runes converted to number %d", ans)
	return ans
}
