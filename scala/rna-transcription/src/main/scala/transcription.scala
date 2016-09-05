case class Dna() {
  private val conv = Map('G' -> 'C', 'C' -> 'G', 'T' -> 'A', 'A' -> 'U')
  def toRna(dna: String) = dna map conv
}
