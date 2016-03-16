package anagram

import (
	"reflect"
	"sort"
	"strings"
)

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Len() int {
	return len(s)
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func Detect(s string, can []string) []string {
	s = strings.ToLower(s)
	r := sortRunes(s)
	sort.Sort(r)

	ans := []string{}
	for _, v := range can {
		v = strings.ToLower(v)
		vv := sortRunes(v)
		sort.Sort(vv)
		if reflect.DeepEqual(vv, r) && v != s {
			ans = append(ans, v)
		}
	}
	return ans
}
