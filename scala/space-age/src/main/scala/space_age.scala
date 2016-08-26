case class SpaceAge(seconds: Long) {
  import SpaceAge._

  // ...
  implicit class ExtendedDouble(n: Double) {
    def rounded(p: Int): Double = {
      val s = math pow (10, p); (math rint n * s) / s
    }
  }

  private def years(yearInSecs: Double) = (seconds / yearInSecs).rounded(2)

  def onEarth   = years(earth)
  def onMercury = years(mercury)
  def onVenus   = years(venus)
  def onMars    = years(mars)
  def onJupiter = years(jupiter)
  def onSaturn  = years(saturn)
  def onUranus  = years(uranus)
  def onNeptune = years(neptune)
}

object SpaceAge {
  val earth   = 31557600 * 1.0
  val mercury = 31557600 * 0.2408467
  val venus   = 31557600 * 0.61519726
  val mars    = 31557600 * 1.8808158
  val jupiter = 31557600 * 11.862615
  val saturn  = 31557600 * 29.447498
  val uranus  = 31557600 * 84.016846
  val neptune = 31557600 * 164.79132
}
