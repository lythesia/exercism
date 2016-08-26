import java.util.{Calendar, GregorianCalendar}

case class Meetup(month: Int, year: Int) {
  def first(weekday: Int) : GregorianCalendar  = Date(1, month, year) next weekday
  def second(weekday: Int) : GregorianCalendar = Date(8, month, year) next weekday
  def third(weekday: Int) : GregorianCalendar  = Date(15, month, year) next weekday
  def fourth(weekday: Int) : GregorianCalendar = Date(22, month, year) next weekday
  def teenth(weekday: Int) : GregorianCalendar = Date(13, month, year) next weekday
  def last(weekday: Int) : GregorianCalendar   = Date(1, month, year) last weekday
}

object Meetup {
  val Sun = Calendar.SUNDAY
  val Mon = Calendar.MONDAY
  val Tue = Calendar.TUESDAY
  val Wed = Calendar.WEDNESDAY
  val Thu = Calendar.THURSDAY
  val Fri = Calendar.FRIDAY
  val Sat = Calendar.SATURDAY
}

case class Date(day: Int, month: Int, year: Int) {
  val date = new GregorianCalendar(year, month - 1, day)

  private def is(weekday: Int) = date.get(Calendar.DAY_OF_WEEK) == weekday

  def next(weekday: Int) = {
    (day until (day + 7)).map(Date(_, month, year)).toStream.filter(_ is weekday).head
  }

  def last(weekday: Int) = {
    Date(date.getActualMaximum(Calendar.DAY_OF_MONTH) - 6, month, year) next weekday
  }
}

object Date {
  import scala.language.implicitConversions

  implicit def date2CalDate(d: Date): GregorianCalendar = d.date
}
