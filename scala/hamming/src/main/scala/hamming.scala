object Hamming {
  def compute(x: String, y: String) = {
    require(x.length == y.length, s"not equal length: $x <-> $y")
    x.zip(y).count(p => p._1 != p._2)
  }
}
