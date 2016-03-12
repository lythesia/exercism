package luhn

import (
	"strconv"
	"strings"
	"unicode"
)

func trim(n string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsDigit(r) {
			return r
		}
		return -1
	}, n)
}

func checksum(n string) int {
	sum := 0
	for i := len(n) - 1; i >= 0; i-- {
		s := int(n[i] - '0')
		if (len(n)-i)%2 == 0 {
			s *= 2
		}
		sum += s%10 + s/10
	}
	return sum
}

func Valid(n string) bool {
	n = trim(n)
	if len(n) == 0 {
		return false
	}
	return checksum(n)%10 == 0
}

func AddCheck(raw string) string {
	return raw + strconv.Itoa((10-checksum(trim(raw)+"0")%10)%10)
}
