package cryptosquare

import (
	"math"
	"strings"
	"unicode"
)

const testVersion = 2

func normalize(plain string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsDigit(r) || unicode.IsLetter(r) {
			return unicode.ToLower(r)
		}
		return -1
	}, plain)
}

func Encode(plain string) string {
	plain = normalize(plain)
	n := len(plain)
	w := int(math.Ceil(math.Sqrt(float64(n))))
	cipher := []string{}
	for i := 0; i < w; i++ {
		c := []byte{}
		for j := i; j < n; j += w {
			c = append(c, plain[j])
		}
		cipher = append(cipher, string(c))
	}
	return strings.Join(cipher, " ")
}
