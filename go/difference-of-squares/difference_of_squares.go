package diffsquares

func SquareOfSums(n int) int {
  return (n*n + 2*n + 1)*(n*n) / 4
}

func SumOfSquares(n int) int {
  ans := 0
  for i := 1; i <= n; i++ {
    ans += i*i
  }
  return ans
}

func Difference(n int) int {
  return SquareOfSums(n) - SumOfSquares(n)
}
