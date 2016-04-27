package octal

import "fmt"

func ParseOctal(s string) (int64, error) {
	decimal := 0
	for _, ch := range s {
		if ch >= '0' && ch <= '7' {
			decimal = decimal*8 + int(ch-'0')
		} else {
			return 0, fmt.Errorf("invalid octal: %s", s)
		}
	}
	return int64(decimal), nil
}
