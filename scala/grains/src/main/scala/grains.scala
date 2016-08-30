object Grains extends WithCache[Int, BigInt] {
  private def doSquare(x: Int) = BigInt(1) << (x - 1)

  def square(x: Int) = withCache(x, doSquare)

  def total = (1 to 64).map(square).sum
}

trait WithCache[K,V] {
  private val cache = collection.mutable.Map[K,V]()

  def withCache(key: K, f: K => V) = cache.getOrElseUpdate(key, f(key))
}
