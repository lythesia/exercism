package prime

func Nth(n int) (int, bool) {
	if n <= 0 {
		return 0, false
	}
	primes := []int{0, 2, 3, 5, 7, 11, 13}
Next:
	for next := primes[len(primes)-1] + 2; len(primes) <= n; next += 2 {
		for f := 1; primes[f]*primes[f] <= next; f++ {
			if next%primes[f] == 0 {
				continue Next
			}
		}
		primes = append(primes, next)
	}
	return primes[n], true
}
