package atbash

import (
	"strings"
	"unicode"
)

func Atbash(plain string) string {
	cipher := make([]rune, 0, len(plain))
	for _, ch := range plain {
		if unicode.IsLetter(ch) {
			cipher = append(cipher, 'a'+25-unicode.ToLower(ch)+'a')
		} else if unicode.IsDigit(ch) {
			cipher = append(cipher, ch)
		}
	}
	n := (len(cipher) + 4) / 5
	groups := make([]string, n)
	for i := 0; i < n; i++ {
		p := (i + 1) * 5
		if p >= len(cipher) {
			p = len(cipher)
		}
		groups[i] = string(cipher[i*5 : p])
	}
	return strings.Join(groups, " ")
}
