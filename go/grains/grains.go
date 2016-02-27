package grains

import "fmt"

func Square(n int) (uint64, error) {
	var ans uint64 = 1
	if n <= 0 || n > 64 {
		return 0, fmt.Errorf("invalid square: %d", n)
	} else {
		return (ans << uint(n-1)), nil
	}
}

func Total() uint64 {
	return (1 << 64) - 1
}
