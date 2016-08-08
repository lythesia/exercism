object Pangrams {
  def isPangram(s: String) = s.toLowerCase.filter(c => c>='a' && c<='z').toSet.size == 26
}
