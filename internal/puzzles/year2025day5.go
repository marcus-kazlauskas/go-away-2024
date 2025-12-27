package puzzles

import (
	"bufio"
	"fmt"
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

func checkFresh(id int64, fresh [][]int64) bool {
	for _, p := range fresh {
		if id >= p[0] && id <= p[1] {
			return true
		}
	}
	return false
}
