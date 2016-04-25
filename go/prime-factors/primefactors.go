package prime

const testVersion = 2

func Factors(n int64) []int64 {
	ans := []int64{}
	for f := int64(2); f*f <= n; {
		switch {
		case n%f == 0:
			ans = append(ans, f)
			n /= f
		case f == 2:
			f++
		default:
			f += 2
		}
	}
	if n != 1 {
		ans = append(ans, n)
	}
	return ans
}
