class Anagram(s: String) {
  def matches(seq: Seq[String]) = {
    val lower = s.toLowerCase
    val base = lower.sorted

    def isAnagram(e: String) = {
      val le = e.toLowerCase
      le != lower && le.sorted == base
    }

    seq.filter(isAnagram)
  }
}
