case class Year(yy: Int) {
  def isLeap = yy % 100 > 0 && yy % 4 == 0 || yy % 400 == 0
}
