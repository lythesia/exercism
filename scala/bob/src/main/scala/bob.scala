class Bob {
  def blank(s: String) = s.isEmpty
  def shout(s: String) = {
    val s1 = s.toUpperCase
    val s2 = s.toLowerCase
    s == s1 && s1 != s2
  }
  def question(s: String) = s.last == '?'

  def hey(s: String) = s.trim match {
    case s if blank(s) => "Fine. Be that way!"
    case s if shout(s) => "Whoa, chill out!"
    case s if question(s) => "Sure."
    case _ => "Whatever."
  }
}
