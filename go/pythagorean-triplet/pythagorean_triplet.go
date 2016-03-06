package pythagorean

import "math"
import "sort"

type Triplet [3]int
type Lex []Triplet

// for sort
// type Interface interface {
//    Len() int    // total #
//    Swap(i, j int)  // swap [i] and [j]
//    Less(i, j int) bool　// less comparator
// }
func (t Lex) Len() int           { return len(t) }
func (t Lex) Swap(i, j int)      { t[i], t[j] = t[j], t[i] }
func (t Lex) Less(i, j int) bool { return t[i][0] < t[j][0] }

// a = m^2 - n^2
// b = 2mn
// c = m^2 + n^2
func Range(min, max int) []Triplet {
	var ans []Triplet
	lm := int(math.Sqrt(float64(max/2))) + 1
	for m := 2; m < lm; m++ {
		for n := 1 + m%2; n < m; n += 2 {
			if coprime(m, n) {
				a, b, c := m*m-n*n, 2*m*n, m*m+n*n
				if a > b {
					a, b = b, a
				}
				for k := 1; ; k++ {
					if k*c > max {
						break
					}
					if k*a < min {
						continue
					}
					ans = append(ans, Triplet{k * a, k * b, k * c})
				}
			}
		}
	}
	sort.Sort(Lex(ans))
	return ans
}

func coprime(x, y int) bool {
	if x > y {
		x, y = y, x
	}
	for y != 0 {
		x, y = y, x%y
	}
	return x == 1
}

func Sum(p int) []Triplet {
	var ans []Triplet
	triples := Range(1, p/2)
	for _, t := range triples {
		if t[0]+t[1]+t[2] == p {
			ans = append(ans, t)
		}
	}
	return ans
}
