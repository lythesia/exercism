import scala.util.Random.{alphanumeric => alnum}

class Robot {

  private var _name = gen

  private def gen = (alnum.filter(_.isUpper).take(2) ++ alnum.filter(_.isDigit).take(3)).mkString

  def name = _name

  def reset(): Unit = _name = gen
}
