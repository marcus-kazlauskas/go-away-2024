package utils

func Abs64(num int64) int64 {
	if num > 0 {
		return num
	} else {
		return -num
	}
}

func Abs(num int) int {
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
