case class Squares() {

  private def sq(n: Int) = n * n

  def squareOfSums(n: Int) = sq((1 to n) sum)

  def sumOfSquares(n: Int) = (1 to n) map sq sum

  def difference(n: Int) = squareOfSums(n) - sumOfSquares(n)
}
