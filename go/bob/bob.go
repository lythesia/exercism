package bob // package name must match the package name in bob_test.go

import (
	"strings"
	"unicode"
)

const testVersion = 2 // same as targetTestVersion

func Hey(s string) string {
	s = strings.TrimSpace(s)
	if s == "" {
		return "Fine. Be that way!"
	} else if strings.IndexFunc(s, func(r rune) bool { return unicode.IsLetter(r) }) != -1 &&
		strings.IndexFunc(s, func(r rune) bool { return unicode.IsLower(r) }) == -1 {
		return "Whoa, chill out!"
	} else if strings.HasSuffix(s, "?") {
		return "Sure."
	} else {
		return "Whatever."
	}
}
