class PhoneNumber(s: String) {
  lazy val number = {
    val n = s.filter(_.isDigit)
    if(n.length == 10)
      n
    else if(n.length == 11 && n(0) == '1')
      n.drop(1)
    else
      "0000000000"
  }

  lazy val areaCode = number.take(3)

  override def toString = s"(${areaCode}) ${number.substring(3,6)}-${number.takeRight(4)}"
}
