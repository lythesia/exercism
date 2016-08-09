class DNA(dna: String) {
  val seq = "ATCG"
  require(dna.forall(seq contains _))
  lazy val nucleotideCounts = seq.map(_ -> 0).toMap ++ dna.groupBy(identity).mapValues(_.size)
}
