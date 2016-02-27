package triangle

import "math"

type Kind int
const (
  Equ = iota
  Iso
  Sca
  NaT
)

func KindFromSides(a, b, c float64) Kind {
  if math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c) ||
     math.IsInf(a, 0) || math.IsInf(b, 0) || math.IsInf(c, 0) ||
     a + b <= c || b + c <= a || a + c <= b {
    return NaT
  }
  if a == b && b == c {
    return Equ
  } else if a == b || b == c || a == c {
    return Iso
  } else {
    return Sca
  }
}
