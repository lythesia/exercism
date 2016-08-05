object Hamming {
  def compute(x: String, y: String) = {
    require(x.length == y.length, s"not equal length: $x <-> $y")
    x.zip(y).count { case (c1, c2) => c1 != c2 }
  }
}
