class Phrase(s: String) {
  lazy val wordCount = s.split("(?i)[^'a-z0-9]+").map(_.toLowerCase).groupBy(identity).mapValues(_.size)
}
