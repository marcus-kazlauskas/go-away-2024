package puzzles

import (
	"errors"
	"fmt"
	"go-away-2024/internal/minio"
	"go-away-2024/internal/utils"
	"io"
	"slices"
	"strconv"
	"strings"
)

func Year2024Day1Part1(scan minio.MinioScanner) (*string, error) {
	var ans int64 = 0
	a := make([]int64, 0)
	b := make([]int64, 0)
	c := make([]int64, 0)

	lineLink, errLink := scan.LineLink()

	for errors.Is(*errLink, nil) {
		list := strings.Fields(*lineLink)
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

		lineLink, errLink = scan.LineLink()
	}
	if !errors.Is(*errLink, io.EOF) {
		return nil, *errLink
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
