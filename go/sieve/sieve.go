package sieve

import "math"

func Sieve(n int) []int {
	noPrime := make([]bool, n+1) // default as false
	x := int(math.Sqrt(float64(n)))
	for i := 2; i <= x; i++ {
		if !noPrime[i] {
			for j := i * i; j <= n; j += i {
				noPrime[j] = true
			}
		}
	}
	primes := make([]int, 0, n)
	for i := 2; i <= n; i++ {
		if !noPrime[i] {
			primes = append(primes, i)
		}
	}
	return primes
}
