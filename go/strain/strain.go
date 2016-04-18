package strain

type Ints []int
type Lists [][]int
type Strings []string

func (ints Ints) Keep(pred func(int) bool) Ints {
	var ans Ints
	for _, e := range ints {
		if pred(e) {
			ans = append(ans, e)
		}
	}
	return ans
}

func (ints Ints) Discard(pred func(int) bool) Ints {
	return ints.Keep(func(i int) bool { return !pred(i) })
}

func (lists Lists) Keep(pred func([]int) bool) Lists {
	var ans Lists
	for _, e := range lists {
		if pred(e) {
			ans = append(ans, e)
		}
	}
	return ans
}

func (strs Strings) Keep(pred func(string) bool) Strings {
	var ans Strings
	for _, e := range strs {
		if pred(e) {
			ans = append(ans, e)
		}
	}
	return ans
}
