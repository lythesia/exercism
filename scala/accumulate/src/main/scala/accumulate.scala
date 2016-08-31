class Accumulate {
  def accumulate[A, B](f: A => B, list: List[A]): List[B] = list.foldRight(List.empty[B]) { (x, acc) => f(x) :: acc }
}
