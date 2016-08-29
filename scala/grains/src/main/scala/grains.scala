object Grains {
  def square(x: Int) = BigInt(1) << (x - 1)
  val total = (BigInt(1) << 64) - 1
}
