// +build !example

package hamming

import "fmt"

const testVersion = 3

func Distance(s1, s2 string) (int, error) {
	ans := 0

	if len(s1) != len(s2) {
		return ans, fmt.Errorf("length not equal: %d != %d", len(s1), len(s2))
	}

	for i := range s1 {
		if s1[i] != s2[i] {
			ans++
		}
	}
	return ans, nil
}
