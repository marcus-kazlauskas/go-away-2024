package puzzles

import (
	"bufio"
	"fmt"
	"go-away-2024/internal/minio"
	"strconv"
	"strings"
)

func Year2025Day2Part1(scan *bufio.Scanner) (*string, error) {
	var ans int = 0

	scan.Scan()
	line := scan.Text()
	for l := range strings.SplitSeq(line, ",") {
		ids := strings.Split(l, "-")
		if len(ids) != 2 {
			return nil, minio.DataError()
		}
		firstId, err := strconv.Atoi(ids[0])
		if err != nil {
			return nil, minio.DataError()
		}
		lastId, err := strconv.Atoi(ids[1])
		if err != nil {
			return nil, minio.DataError()
		}

		for id := firstId; id <= lastId; id++ {
			digit := id % 10
			number := id / 10
			digits := make([]int, 0)
			digits = append(digits, digit)
			for number > 0 {
				digit = number % 10
				number = number / 10
				digits = append(digits, digit)
			}
			if len(digits)%2 != 0 {
				continue
			}

			if checkId(digits) {
				ans += id
			}
		}
	}

	result := fmt.Sprint(ans)
	return &result, nil
}

func checkId(digits []int) bool {
	for i := 0; i < len(digits)/2; i++ {
		if digits[i] != digits[len(digits)/2+i] {
			return false
		}
	}
	return true
}
