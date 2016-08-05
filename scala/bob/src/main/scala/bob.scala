class Bob {
  private def blank(s: String) = s.isEmpty

  private def shout(s: String) = {
    val s1 = s.toUpperCase
    val s2 = s.toLowerCase
    s == s1 && s1 != s2
  }

  private def question(s: String) = s.endsWith("?")

  def hey(s: String) = {
    val ts = s.trim
    if(blank(ts)) "Fine. Be that way!"
    else if(shout(ts)) "Whoa, chill out!"
    else if(question(ts)) "Sure."
    else "Whatever."
  }
}
