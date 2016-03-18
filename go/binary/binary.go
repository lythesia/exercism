package binary

import (
	"fmt"
)

// what if overflow?
func ParseBinary(s string) (int, error) {
	ans := 0
	for _, v := range s {
		x := int(v - '0')
		if x >= 2 {
			return 0, fmt.Errorf("invalid binary: %v", s)
		}
		ans = (ans << 1) + x
	}
	return ans, nil
}
