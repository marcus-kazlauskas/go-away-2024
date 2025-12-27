package puzzles

import (
	"bufio"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2/log"
)

func Year2025Day5Part1(scan *bufio.Scanner) (*string, error) {
	var ans int = 0
	fresh := make([][]int64, 0)
	write := true

	for scan.Scan() {
		line := scan.Text()
		if line == "" {
			write = false
			log.Debug("All ids written")
			continue
		}

		if write {
			p, err := getFreshIds(line)
			if err != nil {
				return nil, err
			}
			fresh = append(fresh, p)
		} else {
			value, err := strconv.ParseInt(line, 10, 64)
			if err != nil {
				return nil, err
			}

			if checkFresh(value, fresh) {
				log.Debugf("Fresh id=%d", value)
				ans++
			} else {
				log.Debugf("Spoiled id=%d", value)
			}
		}
	}

	result := fmt.Sprint(ans)
	return &result, nil
}

func getFreshIds(line string) ([]int64, error) {
	values := strings.Split(line, "-")
	if len(values) != 2 {
		return nil, DataError()
	}
	l, err := strconv.ParseInt(values[0], 10, 64)
	if err != nil {
		return nil, DataError()
	}
	r, err := strconv.ParseInt(values[1], 10, 64)
	if err != nil {
		return nil, DataError()
	}

	p := make([]int64, 0)
	p = append(p, l)
	p = append(p, r)
	return p, nil
}

func checkFresh(id int64, fresh [][]int64) bool {
	for _, p := range fresh {
		if id >= p[0] && id <= p[1] {
			return true
		}
	}
	return false
}

func Year2025Day5Part2(scan *bufio.Scanner) (*string, error) {
	var ans int64 = 0
	freshIds := make([][]int64, 0)

	for scan.Scan() {
		line := scan.Text()
		if line == "" {
			break
		}

		p, err := getFreshIds(line)
		if err != nil {
			return nil, err
		}
		freshIds = append(freshIds, p)
	}

	slices.SortFunc(freshIds, func(a []int64, b []int64) int {
		delta := a[0] - b[0]
		if delta == 0 {
			return int(a[1] - b[1])
		} else {
			return int(delta)
		}
	})
	for _, ids := range freshIds {
		log.Debugf("%v", ids)
	}

	if len(freshIds) < 1 {
		return nil, DataError()
	}
	fresh := mergeFreshIds(freshIds)

	ans = countFreshIds(fresh) // too low

	result := fmt.Sprint(ans)
	return &result, nil
}

func mergeFreshIds(freshIds [][]int64) [][]int64 {
	fresh := make([][]int64, 0)
	currentFreshIds := freshIds[0]
	for _, ids := range freshIds {
		if ids[0] <= currentFreshIds[1]+1 && ids[1] > currentFreshIds[1] {
			currentFreshIds[1] = ids[1]
		} else if ids[0] > currentFreshIds[1]+1 {
			fresh = append(fresh, currentFreshIds)
			log.Debugf("Added to fresh: %v", currentFreshIds)
			currentFreshIds = ids
		}
	}
	fresh = append(fresh, currentFreshIds)
	log.Debugf("Added to fresh: %v", currentFreshIds)
	return fresh
}

func countFreshIds(freshIds [][]int64) int64 {
	var ans int64 = 0
	for _, ids := range freshIds {
		ans += ids[1] - ids[0] + 1
	}
	return ans
}
