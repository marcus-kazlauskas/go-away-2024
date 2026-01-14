package utils

func Abs[T int | int64](num T) T {
	if num > 0 {
		return num
	} else {
		return -num
	}
}

func Pow64(base int, n int) int64 {
	if n <= 0 {
		return 1
	} else if n%2 == 0 {
		return Pow64(base, n/2) * Pow64(base, n/2)
	} else {
		return int64(base) * Pow64(base, n-1)
	}
}
