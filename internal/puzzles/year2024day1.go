package puzzles

import (
	"bufio"
	"fmt"
	"go-away-2024/internal/utils"
	"slices"
	"strconv"
	"strings"
)

func Year2024Day1Part1(scan *bufio.Scanner) (*string, error) {
	var ans int64 = 0
	a := make([]int64, 0)
	b := make([]int64, 0)
	c := make([]int64, 0)

	for scan.Scan() {
		list := strings.Fields(scan.Text())
		if len(list) != 2 {
			return nil, DataError()
		}

		num, err := strconv.Atoi(list[0])
		if err != nil {
			return nil, DataError()
		}
		a = append(a, int64(num))

		num, err = strconv.Atoi(list[1])
		if err != nil {
			return nil, DataError()
		}
		b = append(b, int64(num))
	}

	slices.Sort(a)
	slices.Sort(b)
	for i := range len(a) {
		c = append(c, utils.Abs64(a[i]-b[i]))
	}
	for i := range len(a) {
		ans += c[i]
	}

	result := fmt.Sprint(ans)
	return &result, nil
}

func Year2024Day1Part2(scan *bufio.Scanner) (*string, error) {
	var ans int64 = 0
	a := make(map[int]int)
	b := make(map[int]int)

	for scan.Scan() {
		list := strings.Fields(scan.Text())
		if len(list) != 2 {
			return nil, DataError()
		}

		num, err := strconv.Atoi(list[0])
		if err != nil {
			return nil, DataError()
		}
		_, exists := a[num]
		if exists {
			a[num] += 1
		} else {
			a[num] = 1
		}

		num, err = strconv.Atoi(list[1])
		if err != nil {
			return nil, DataError()
		}
		_, exists = b[num]
		if exists {
			b[num] += 1
		} else {
			b[num] = 1
		}
	}

	for k, v := range a {
		count, exists := b[k]
		if exists {
			ans += int64(k * v * count)
		}
	}

	result := fmt.Sprint(ans)
	return &result, nil
}
