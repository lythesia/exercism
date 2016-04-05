package say

import (
	"strings"
)

var largen = []uint64{1000000000000000000, 1000000000000000, 1000000000000, 1000000000, 1000000, 1000}
var largew = []string{"quintillion", "quadrillion", "trillion", "billion", "million", "thousand"}
var ones = []string{"", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
var tens = []string{"", "", "twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety"}
var ten = []string{"ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen"}

func conv(n uint64) []string {
	ans := make([]string, 0)
	if n >= 100 {
		x := n / 100
		ans = append(ans, ones[x], "hundred")
		n -= x * 100
	}
	if n >= 20 {
		x := n / 10
		if n%10 > 0 {
			ans = append(ans, tens[x]+"-"+ones[n%10])
		} else {
			ans = append(ans, tens[x])
		}
	} else if n >= 10 {
		ans = append(ans, ten[n%10])
	} else if n > 0 {
		ans = append(ans, ones[n])
	}
	return ans
}

func Say(n uint64) string {
	ans := make([]string, 0)
	for i, e := range largen {
		x := n / e
		if x > 0 {
			ans = append(ans, conv(x)...)
			ans = append(ans, largew[i])
		}
		n -= x * e
	}
	if n > 0 {
		ans = append(ans, conv(n)...)
	}
	if len(ans) == 0 {
		return "zero"
	}
	return strings.Join(ans, " ")
}
