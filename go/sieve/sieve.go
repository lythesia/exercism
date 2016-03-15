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
	// fit memory
	k := 0
	for _, v := range noPrime[2:] {
		if !v {
			k++
		}
	}
	primes := make([]int, k)
	for i, j := 2, 0; i <= n; i++ {
		if !noPrime[i] {
			primes[j] = i
			j++
		}
	}
	return primes
}
