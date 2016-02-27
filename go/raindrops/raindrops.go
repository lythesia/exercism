package raindrops

import "fmt"

var prim = [...]int { 3, 5, 7 }
var rain = [...]string { "Pling", "Plang", "Plong" }

func Convert(n int) string {
  var ans string
  for i := 0; i < 3; i++ {
    if n % prim[i] == 0 {
      ans += rain[i]
    }
  }
  if ans == "" {
    return fmt.Sprintf("%d", n)
  } else {
    return ans
  }
}
