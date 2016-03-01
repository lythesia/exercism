package summultiples

func MultipleSummer(divisors ...int) func(int) int {
	return func(limit int) int {
		ans := 0
		for i := 1; i < limit; i++ {
			for _, d := range divisors {
				if i%d == 0 {
					ans += i
					break
				}
			}
		}
		return ans
	}
}
