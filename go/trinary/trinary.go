package trinary

import (
	"fmt"
	"strings"
)

func ParseTrinary(n string) (int64, error) {
	var trinary int64 = 0
	p := strings.IndexFunc(n, func(r rune) bool { return r != '0' })
	if p != -1 {
		for ; p != len(n); p++ {
			nTrinary := trinary*3 + int64(n[p]-'0')
			if trinary > nTrinary {
				return 0, fmt.Errorf("overflow: %s", n)
			}
			trinary = nTrinary
		}
	}
	return trinary, nil
}
