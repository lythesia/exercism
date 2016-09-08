case class Triangle(a: Int, b: Int, c: Int) {
  import TriangleType._

  private def valid = List(a, b, c).forall(_ > 0) && (a + b > c) && (b + c > a) && (a + c > b)

  def triangleType: Type = {
    if(valid) {
      if(a == b && b == c) Equilateral
      else if(a == b || a == c || b == c) Isosceles
      else Scalene
    }
    else Illogical
  }
}

object TriangleType {
  sealed trait Type
  object Equilateral extends Type
  object Isosceles extends Type
  object Scalene extends Type
  object Illogical extends Type
}
