package sieve

import "math"

func Sieve(n int) []int {
	noprime := make([]bool, n+1) // default as false
	x := int(math.Sqrt(float64(n)))
	for i := 2; i <= x; i++ {
		if !noprime[i] {
			for j := i * i; j <= n; j += i {
				noprime[j] = true
			}
		}
	}
	primes := []int{}
	for i := 2; i <= n; i++ {
		if !noprime[i] {
			primes = append(primes, i)
		}
	}
	return primes
}
