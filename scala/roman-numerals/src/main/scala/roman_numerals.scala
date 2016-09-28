case class RomanNumeral(x: Int) {
  import RomanNumeral._

  private def num2roman(n: Int, i: Int): String = {
    if(n < 4) ONES(i) * n
    else if(n == 4) ONES(i) + FIVES(i)
    else if(n < 9) FIVES(i) + ONES(i) * (n-5)
    else ONES(i) + ONES(i + 1)
  }

  def value = {
    x.toString.reverse.map { _ - '0' }.zipWithIndex.map { case (digit, index) => num2roman(digit, index) }.reverse.mkString
  }
}

object RomanNumeral {
  val ONES = Array("I", "X", "C", "M")
  val FIVES = Array("V", "L", "D")
}
