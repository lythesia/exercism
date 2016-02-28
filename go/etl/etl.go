package etl

import "strings"

type oldt map[int][]string
type newt map[string]int

func Transform(old oldt) newt {
	var ans = newt{}
	for k, v := range old {
		for _, e := range v {
			ans[strings.ToLower(e)] = k
		}
	}
	return ans
}
