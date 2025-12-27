package puzzles

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2/log"
)

func Year2025Day2Part1(scan *bufio.Scanner) (*string, error) {
	var ans int = 0

	scan.Scan()
	line := scan.Text()
	for l := range strings.SplitSeq(line, ",") {
		ids := strings.Split(l, "-")
		if len(ids) != 2 {
			return nil, DataError()
		}
		firstId, lastId, err := getIdRange(ids)
		if err != nil {
			return nil, DataError()
		}

		for id := firstId; id <= lastId; id++ {
			digits := getDigits(id)
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

func getIdRange(ids []string) (firstId int, lastId int, err error) {
	firstId, err = strconv.Atoi(ids[0])
	if err != nil {
		return
	}
	lastId, err = strconv.Atoi(ids[1])
	if err != nil {
		return
	}
	return
}

func getDigits(id int) (digits []int) {
	digit := id % 10
	number := id / 10
	digits = append(digits, digit)
	for number > 0 {
		digit = number % 10
		number = number / 10
		digits = append(digits, digit)
	}
	return
}

func checkId(digits []int) bool {
	for i := 0; i < len(digits)/2; i++ {
		if digits[i] != digits[len(digits)/2+i] {
			return false
		}
	}
	return true
}

func Year2025Day2Part2(scan *bufio.Scanner) (*string, error) {
	var ans int = 0

	scan.Scan()
	line := scan.Text()
	for l := range strings.SplitSeq(line, ",") {
		ids := strings.Split(l, "-")
		if len(ids) != 2 {
			return nil, DataError()
		}
		firstId, lastId, err := getIdRange(ids)
		if err != nil {
			return nil, DataError()
		}

		for id := firstId; id <= lastId; id++ {
			digits := getDigits(id)
			if checkIdV2(digits) {
				log.Debugf("Invalid id found: id=%d len()=%d digits=%v", id, len(digits), digits)
				ans += id
			}
		}
	}

	result := fmt.Sprint(ans)
	return &result, nil
}

func checkIdV2(digits []int) bool {
	for r := 1; r <= len(digits)/2; r++ {
		n := len(digits) / r
		if len(digits)%r != 0 {
			continue
		}
		check := true
		for i := 0; i < r; i++ {
			for j := 1; j < n; j++ {
				check = check && digits[i] == digits[r*j+i]
			}
		}
		if check {
			return true
		}
	}
	return false
}
