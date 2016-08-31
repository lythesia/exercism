class Sublist {
  def sublist[T](x: List[T], y: List[T]): Sublist.Result = {
    if(x == y) Sublist.Equal
    else if(x containsSlice y) Sublist.Superlist
    else if(y containsSlice x) Sublist.Sublist
    else Sublist.Unequal
  }
}

object Sublist {
  sealed trait Result
  case object Equal extends Result
  case object Superlist extends Result
  case object Sublist extends Result
  case object Unequal extends Result
}
