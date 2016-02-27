package slice

func Frist(n int, s string) string {
	return s[0:n]
}

func First(n int, s string) (string, bool) {
	if len(s) < n {
		return "err", false
	} else {
		return s[0:n], true
	}
}

func All(n int, s string) []string {
	if len(s) < n {
		return nil
	} else {
		end := len(s) - n
		ans := make([]string, end+1)
		for i := 0; i <= end; i++ {
			ans[i] = s[i : i+n]
		}
		return ans
	}
}
