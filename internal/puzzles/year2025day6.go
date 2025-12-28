package puzzles

import (
	"bufio"
	"fmt"
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
