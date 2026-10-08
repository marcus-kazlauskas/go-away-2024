package utils

func Abs[T int | int64](num T) T {
	if num > 0 {
		return num
	} else {
		return -num
	}
}

func Pow64(base int, n int) int64 {
	switch {
	case n <= 0:
		return 1
	case n%2 == 0:
		return Pow64(base, n/2) * Pow64(base, n/2)
	default:
		return int64(base) * Pow64(base, n-1)
	}
}
