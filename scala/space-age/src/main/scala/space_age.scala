case class SpaceAge(seconds: Long) {
  import SpaceAge._

  // ...
  implicit class ExtendedDouble(n: Double) {
    def rounded(p: Int): Double = {
      val s = math pow (10, p); (math rint n * s) / s
    }
  }

  private def years(yearInSecs: Double) = (seconds / yearInSecs).rounded(2)

  def onEarth   = years(EARTH)
  def onMercury = years(MERCURY)
  def onVenus   = years(VENUS)
  def onMars    = years(MARS)
  def onJupiter = years(JUPITER)
  def onSaturn  = years(SATURN)
  def onUranus  = years(URANUS)
  def onNeptune = years(NEPTUNE)
}

object SpaceAge {
  val EARTH   = 31557600 * 1.0
  val MERCURY = 31557600 * 0.2408467
  val VENUS   = 31557600 * 0.61519726
  val MARS    = 31557600 * 1.8808158
  val JUPITER = 31557600 * 11.862615
  val SATURN  = 31557600 * 29.447498
  val URANUS  = 31557600 * 84.016846
  val NEPTUNE = 31557600 * 164.79132
}
