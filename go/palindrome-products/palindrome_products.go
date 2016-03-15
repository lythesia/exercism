package palindrome

import (
	"fmt"
	"strconv"
)

type Product struct {
	Product        int
	Factorizations [][2]int
}

func isPali(x int) bool {
	if x < 0 {
		return false
	}

	s := strconv.Itoa(x)
	n := len(s)

	for i := 0; i <= n/2; i++ {
		if s[i] != s[n-1-i] {
			return false
		}
	}
	return true
}

func Products(fmin, fmax int) (Product, Product, error) {
	if fmin > fmax {
		return Product{}, Product{}, fmt.Errorf("fmin > fmax")
	}

	prod := map[int]Product{}
	mmin, mmax := int((^uint(0))>>1), 0

	for i := fmin; i <= fmax; i++ {
		for j := i; j <= fmax; j++ {
			p := i * j
			if isPali(p) {
				if v, ok := prod[p]; ok {
					v.Factorizations = append(v.Factorizations, [2]int{i, j})
					prod[p] = v
				} else {
					prod[p] = Product{p, [][2]int{{i, j}}}
				}
				if p < mmin {
					mmin = p
				}
				if p > mmax {
					mmax = p
				}
			}
		}
	}

	if len(prod) == 0 {
		return Product{}, Product{}, fmt.Errorf("No palindromes")
	}
	return prod[mmin], prod[mmax], nil
}
