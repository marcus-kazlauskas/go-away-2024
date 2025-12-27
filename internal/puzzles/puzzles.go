package puzzles

import "fmt"

func DataError() error {
	return fmt.Errorf("incorrect input data")
}
