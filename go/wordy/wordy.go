package wordy

import (
	"strconv"
	"strings"
)

func Answer(q string) (int, bool) {
	var op string
	var ans, c int
	if !strings.HasPrefix(q, "What is") {
		goto invalid
	}

	q = q[8 : len(q)-1]
	for _, e := range strings.Split(q, " ") {
		if i, err := strconv.Atoi(e); err == nil {
			c++
			if op == "" {
				ans = i
			} else {
				switch op {
				case "plus":
					ans += i
				case "minus":
					ans -= i
				case "times":
				case "multipliedby":
					ans *= i
				case "dividedby":
					ans /= i
				}
				op = ""
			}
		} else {
			op += e
		}
	}
	if c < 2 {
		goto invalid
	}
	return ans, true
invalid:
	return 0, false
}
