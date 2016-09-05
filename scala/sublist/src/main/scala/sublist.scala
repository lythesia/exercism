import scala.annotation.tailrec

class Sublist {
  import Sublist.ListOp

  def sublist[T](x: List[T], y: List[T]): Sublist.Result = {
    if(x == y) Sublist.Equal
    else if(x has y) Sublist.Superlist
    else if(y has x) Sublist.Sublist
    else Sublist.Unequal
  }
}

object Sublist {
  sealed trait Result
  case object Equal extends Result
  case object Superlist extends Result
  case object Sublist extends Result
  case object Unequal extends Result

  implicit class ListOp[T](me: List[T]) {
    def has(other: List[T]): Boolean = {
      @tailrec
      def iter(rest: List[T]): Boolean = {
        if(rest.size < other.size) false
        else if(rest startsWith other) true
        else iter(rest.tail)
      }
      iter(me)
    }
  }
}
