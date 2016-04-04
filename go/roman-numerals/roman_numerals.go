package romannumerals

import (
	"fmt"
)

const testVersion = 2

var digit = []int{1000, 100, 10, 1}
var ten = []string{"M", "C", "X", "I"}
var fiv = []string{"D", "L", "V"}

func ToRomanNumeral(n int) (string, error) {
	ans := ""
	if n <= 0 || n >= 4000 {
		return ans, fmt.Errorf("invalid range")
	}
	for i := 0; i < 4; i++ {
		x := n / digit[i]
		if x > 0 {
			if x < 4 {
				for c := x; c != 0; c-- {
					ans += ten[i]
				}
			} else if x == 4 {
				ans += ten[i]
				ans += fiv[i-1]
			} else if x < 9 {
				ans += fiv[i-1]
				for c := x - 5; c != 0; c-- {
					ans += ten[i]
				}
			} else if x == 9 {
				ans += ten[i]
				ans += ten[i-1]
			}
			n -= x * digit[i]
		}
	}
	return ans, nil
}
