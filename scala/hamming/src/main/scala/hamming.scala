object Hamming {
  def compute(x: String, y: String) = {
    if(x.length == y.length)
      x.zip(y).count(p => p._1 != p._2)
    else
      throw new IllegalArgumentException(s"not equal length: $x <-> $y")
  }
}
