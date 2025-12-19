package puzzles

import (
	"bufio"
	"fmt"
	"go-away-2024/internal/minio"
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
			return nil, minio.DataError()
		}

		num, err := strconv.Atoi(list[0])
		if err != nil {
			return nil, minio.DataError()
		}
		a = append(a, int64(num))

		num, err = strconv.Atoi(list[1])
		if err != nil {
			return nil, minio.DataError()
		}
		b = append(b, int64(num))
	}

	slices.Sort(a)
	slices.Sort(b)
	for i := range len(a) {
		c = append(c, utils.Abs(a[i]-b[i]))
	}
	for i := range len(a) {
		ans += c[i]
	}

	result := fmt.Sprint(ans)
	return &result, nil
}
