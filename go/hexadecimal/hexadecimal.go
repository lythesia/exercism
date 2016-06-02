package hexadecimal

import (
	"fmt"
	"unicode"
)

func ParseHex(hex string) (int64, error) {
	if hex == "" {
		return 0, fmt.Errorf("syntax")
	}
	var n int64 = 0
	for _, r := range hex {
		r = unicode.ToLower(r)
		var v int64
		if unicode.IsDigit(r) {
			v = int64(r - '0')
		} else if r >= 'a' && r <= 'f' {
			v = int64(r-'a') + 10
		} else {
			return 0, fmt.Errorf("syntax")
		}
		new := n*16 + v
		if new < n {
			return 0, fmt.Errorf("range")
		}
		n = new
	}
	return n, nil
}

func HandleErrors(hexes []string) []string {
	results := make([]string, len(hexes))
	for i, hex := range hexes {
		_, err := ParseHex(hex)
		if err != nil {
			results[i] = err.Error()
		} else {
			results[i] = "none"
		}
	}
	return results
}
