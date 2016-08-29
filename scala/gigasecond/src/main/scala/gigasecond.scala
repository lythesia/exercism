import java.util.{GregorianCalendar, Calendar}

case class Gigasecond(cal: GregorianCalendar) {
  val date = {
    val ncal = cal.clone.asInstanceOf[GregorianCalendar]
    ncal.add(Calendar.SECOND, 1000000000)
    ncal
  }
}
