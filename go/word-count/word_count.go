package wordcount

import (
	"regexp"
	"strings"
)

const testVersion = 2

// Use this return type.
type Frequency map[string]int

// Just implement the function.
func WordCount(phrase string) Frequency {
	ans := Frequency{}
	for _, v := range regexp.MustCompile(`\w+`).FindAllString(strings.ToLower(phrase), -1) {
		ans[v]++
	}
	return ans
}
